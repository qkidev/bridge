const {ethers} = require('ethers')
require('dotenv').config()
const mysql = require('mysql')

const connection = mysql.createConnection({
    host: process.env.DB_HOST,
    user: process.env.DB_USER,
    password: process.env.DB_PASSWORD,
    database: process.env.DB_DATABASE,
})

connection.connect()

// 取系统配置跨链审核的值
const getIsCheck = () => {
    return new Promise((resolve, reject) => {
        connection.query('SELECT * FROM `setting` WHERE `name` = ?', [`withdraw-check`], (err, res, fields) => {
            if (err) {
                return reject(err)
            } else {
                return resolve(res[0].value)
            }
        })
    })
}

// 取链当前锁值
const getChainLock = (chain) => {
    return new Promise((resolve, reject) => {
        connection.query('SELECT * FROM `setting` WHERE `name` = ?', [`${chain}_lock_number`], (err, res, fields) => {
            if (err) {
                return reject(err)
            } else {
                return resolve(res[0].value)
            }
        })
    })
}

const setChainLock = (chainId, number) => {
    return new Promise((resolve, reject) => {
        connection.query('UPDATE setting SET `value` = ? WHERE `name` = ?', [number, `${chainId}_lock_number`], (error, results, fields) => {
            if (error) {
                return reject(error)
            } else {
                return resolve(results)
            }
        })
    })
}

const getChains = () => {
    return new Promise((resolve, reject) => {
        connection.query('SELECT * FROM chain WHERE `status`= 1', (error, results, fields) => {
            if (error) {
                return reject(error)
            } else {
                return resolve(results)
            }
        })
    })
}

const getPair = (fromChain, toChain, fromToken, toToken) => {
    return new Promise((resolve, reject) => {
        connection.query('SELECT * FROM `pair` WHERE `fromChain` = ? AND `toChain` = ? AND `fromToken` = ? AND `toToken` = ?', [fromChain, toChain, fromToken, toToken], (err, res, fields) => {
            if (err) {
                return reject(err)
            } else {
                return resolve(res[0])
            }
        })
    })
}

const getPairNative = (fromChainId, toChainId, isMain) => {
    return new Promise((resolve, reject) => {
        connection.query('SELECT * FROM `pair` WHERE `fromChain` = ? AND `toChain` = ? AND `isMain` = ? AND `isNative` = 1', [fromChainId, toChainId, isMain], (err, res, fields) => {
            if (err) {
                return reject(err)
            } else {
                return resolve(res[0])
            }
        })
    })
}

const logSave = (pairId, recipient, value, fromChain, toChain, depositHash) => {
    let depositTime = (new Date()).getTime()
    depositTime= (depositTime/1000).toFixed()
    const data = {pairId, recipient, value, fromChain, toChain, depositHash, depositTime}
    connection.query('INSERT INTO log SET ?', data, function (error, results, fields) {
        if (error) throw error;
    });
}

const withdrawDone = (hash, pairId, recipient, value) => {
    let time = (new Date()).getTime()
    time= (time/1000).toFixed()
    connection.query('UPDATE log SET withdrawHash = ?, withdrawTime = ? WHERE pairId = ? AND recipient= ? AND value = ?', [hash, time, pairId, recipient, value], function (error, results, fields) {
        if (error) throw error;
    });

}

// 跨链桥ABI
const abiBridge = [
    "event Deposit(uint toChainId, address fromToken, address toToken, address recipient, uint256 value)",
    "event DepositNative(uint toChainId, bool isMain, address recipient, uint256 value)",
    "event WithdrawDone(uint toChainId, address fromToken, address toToken, address recipient, uint256 value)",
    "event WithdrawNativeDone(uint fromChainId, address recipient,bool isMain, uint256 value)",
    "function withdraw(uint toChainId,address toToken,address recipient,uint256 value)",
    "function withdrawNative(uint toChainId, address payable recipient, bool isMain, uint256 value)"
]

// 全部链的跨链桥合约
let bridgeContracts = {}

async function main() {
    // 支持链主网
    const chains = await getChains()
    // 设置提供者和钱包和跨链桥合约
    chains.forEach(item => {
        const provider = new ethers.providers.JsonRpcProvider(item.url)
        const wallet = new ethers.Wallet(process.env['PK_' + item.name], provider)
        bridgeContracts[item.chainId] = new ethers.Contract(item.bridge, abiBridge, wallet)
    })

    chains.forEach(item => {
        const contract = bridgeContracts[item.chainId]

        // 主网币跨入成功
        contract.on("WithdrawNativeDone", async (toChainId, recipient, isMain, value, event) => {
            toChainId = toChainId.toString()
            const pair = await getPairNative(item.chainId, toChainId, isMain)
            if (pair) {
                withdrawDone(event.transactionHash, pair.id, recipient, value)
                console.log(`[主网币][到账][成功] ${value} 个  ${pair.name} 从  ${item.chainId} 到 ${toChainId}`)
            } else {
                console.log(`跨链对不存在: fromChainId=${item.chainId}&toChainId=${toChainId}&isMain=${isMain}`)
            }
        })

        // 代币跨入成功
        contract.on("WithdrawDone", async (toChainId, fromToken, toToken, recipient, value, event) => {
            toChainId = toChainId.toString()
            const pair = await getPair(item.chainId, toChainId, fromToken, toToken)
            if (pair) {
                console.log("[到账][成功] 链 " + toChainId, "代币 " + pair.name, "地址 " + recipient, "数额 " + value)
                withdrawDone(event.transactionHash, pair.id, recipient, value)
            }
        })

        // 主网币跨出
        contract.on("DepositNative", async (toChainId, isMain, recipient, value, event) => {
            toChainId = toChainId.toString()
            isMain = isMain ? 1 : 0
            value = value.toString() * 1
            const lock = await getChainLock(item.chainId)
            const numberNow = event.blockNumber

            // 检查链高度
            if (numberNow > lock) {
                const pair = await getPairNative(item.chainId, toChainId, isMain)
                logSave(pair.id, recipient, value, item.chainId, toChainId, event.transactionHash)
                console.log(`[主网币][跨链][成功] ${value} 个 ${pair.name} 从 ${pair.fromChain} 到 ${pair.toChain}`)
                await setChainLock(item.chainId, numberNow)
                const isCheck = await getIsCheck()
                // 检查配置的审核状态
                if (!isCheck || (isCheck && value <= pair['limit'])) {
                    const toContract = bridgeContracts[toChainId]
                    if (toContract) {
                        const fee = Math.ceil(value * pair['bridgeFee'] / 100)
                        const final = ethers.utils.parseUnits((value - fee).toString(),pair['decimal'])
                        try {
                            await toContract['withdrawNative'](item.chainId, recipient, !isMain, final)
                        } catch (e) {
                            console.log(e)
                        }
                    }
                }
            }
        })

        // 执行代币跨出操作
        contract.on("Deposit", async (toChainId, fromToken, toToken, recipient, value, event) => {
            toChainId = toChainId.toString()
            const lock = await getChainLock(item.chainId)
            const numberNow = event.blockNumber

            // 检查区块高度
            if (numberNow > lock) {
                const pair = await getPair(item.chainId, toChainId, fromToken, toToken)
                logSave(pair.id, recipient, value, item.chainId, toChainId, event.transactionHash)
                await setChainLock(item.chainId, numberNow.toString())
                const isCheck = await getIsCheck()
                // 检查配置的审核状态
                if (!isCheck || (isCheck && value <= pair['limit'])) {
                    const toContract = bridgeContracts[toChainId]
                    if (toContract) {
                        const fee = Math.ceil(value * pair['bridgeFee'] / 100)
                        const final =  ethers.utils.parseUnits((value - fee).toString(),pair['decimal'])
                        toContract['withdraw'](item.chainId, fromToken, recipient, final)
                    }
                }
            }
        })
    })
}

main().then()