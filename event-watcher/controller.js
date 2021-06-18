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

const logSave = (pairId,recipient,value,fromChain,toChain,depositHash,depositTime) => {
    const data = {pairId,recipient,value,fromChain,toChain,depositHash,depositTime}
    connection.query('INSERT INTO log SET ?', data, function (error, results, fields) {
        if (error) throw error;
    });
}

const logUpdate = ()=>{

}

// 跨链桥ABI
const abiBridge = [
    "event Deposit(uint toChainId, address fromToken, address toToken, address recipient, uint256 value)",
    "event DepositNative(uint toChainId, bool isMain, address recipient, uint256 value)",
    "event WithdrawDone(uint toChainId, address fromToken, address toToken, address recipient, uint256 value)",
    "event WithdrawNativeDone(uint fromChainId, address recipient,bool isMain, uint256 value)",
    "function withdraw(uint toChainId,address toToken,address recipient,uint256 value)",
    "function withdrawNative(uint fromChainId, address payable recipient, bool isMain, uint256 value)"
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
        contract.on("WithdrawDone", async (toChainId, fromToken, toToken, recipient, value, event) => {
            toChainId = toChainId.toString()
            const pair = await getPair(item.chainId, toChainId, fromToken, toToken)
            if (pair) {
                // let number = ethers.utils.formatUnits(value, pair['decimal'])
                console.log("[到账][成功] 链 " +toChainId, "代币 " + pair.name, "地址 " + recipient, "数额 " + value)
            }
        })
        contract.on("DepositNative", async (toChainId, isMain, recipient, value, event) => {
            toChainId = toChainId.toString()
            isMain = isMain ? 1 : 0
            value = value.toString() * 1
            // console.log(toChainId, isMain, recipient, value.toString())
            const lock = await getChainLock(item.chainId)
            const numberNow = event.blockNumber
            if (numberNow > lock) {
                const pair = await getPairNative(item.chainId, toChainId, isMain)
                const fee = Math.ceil(value * pair['bridgeFee'] / 100)
                const toContract = bridgeContracts[toChainId]
                if (toContract) {
                    const final = value - fee
                    try {
                        await toContract.withdrawNative(item.chainId, recipient, !isMain, final.toString())
                        const showFinal = ethers.utils.formatUnits(final.toString(), pair['decimal'])
                        logSave(pair.id,recipient,value,item.chainId,toChainId,event.transactionHash,new Date().getTime())
                        console.log("[主网币][跨链][成功] 链 " + item.chainId, "到链 " + toChainId, "地址 " + recipient, "数额 " + showFinal)
                    } catch (e) {
                        console.log(e)
                    }
                }

                await setChainLock(item.chainId, numberNow)
            }
        })

        contract.on("Deposit", async (toChainId, fromToken, toToken, recipient, value, event) => {
            toChainId = toChainId.toString()
            // console.log(toChainId, fromToken, toToken, recipient, value)
            const lock = await getChainLock(item.chainId)
            const numberNow = event.blockNumber
            if (numberNow > lock) {
                const pair = await getPair(item.chainId, toChainId, fromToken, toToken)
                value = ethers.utils.formatUnits(value, pair['decimal'] * 1)
                let fee = Math.ceil(value * pair['bridgeFee'] / 100)
                value -= fee
                const toContract = bridgeContracts[toChainId]
                if (toContract) {
                    const final = ethers.utils.parseUnits(value.toString(), pair['decimal'])
                    toContract.withdraw(item.chainId, fromToken, recipient, final).then(_ => {
                        logSave(pair.id,recipient,value,item.chainId,toChainId,event.transactionHash,new Date().getTime())
                        console.log("[跨链][成功] 链 " + item.chainId, "到链 " + toChainId, "代币 " + toToken, "地址 " + recipient, "数额 " + value)
                    }).catch(error => {
                        console.log("[跨链][失败] 链 " + item.chainId, "到链 " + toChainId, "代币 " + toToken, "地址 " + recipient, "数额 " + value)
                        // console.log(error)
                    })
                }
                await setChainLock(item.chainId, numberNow.toString())
            }
        })
    })
}

main().then(_ => {

})