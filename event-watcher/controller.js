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

const logSave = () => {

}

// 跨链桥ABI
const abiBridge = [
    "event Deposit(uint toChainId, address fromToken, address toToken, address recipient, uint256 value)",
    "event DepositNative(uint toChainId, bool isMain, address recipient, uint256 value)",
    "event WithdrawDone(uint toChainId, address fromToken, address toToken, address recipient, uint256 value)",
    "event WithdrawNativeDone(uint fromChainId, address recipient, uint256 value)",
    "function withdraw(uint toChainId,address toToken,address recipient,uint256 value)",
    "function withdrawNative(uint fromChainId, address payable recipient, uint256 value)"
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
                let number = ethers.utils.formatUnits(value, pair['decimal'])
                console.log("[到账][成功] 链 " + item.name, "代币 " + pair.name, "地址 " + recipient, "数额 " + number)
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
                        console.log(item.chainId, recipient, final)
                        await toContract.withdrawNative(item.chainId, recipient, final.toString())
                        const showFinal = ethers.utils.formatUnits(final.toString(), pair['decimal'])
                        console.log("[主网币][跨链][成功] 链 " + item.chainId, "到链 " + toChainId, "地址 " + recipient, "数额 " + showFinal)
                    } catch (e) {
                        console.log(e)
                    }
                    // logInsert({
                    //     depositHash: event.blockHash,
                    //     chainFrom: item.name,
                    //     chainTo: toChain,
                    //     block: numberNow,
                    //     toToken: "",
                    //     address: recipient,
                    //     value: value - fee,
                    //     status: 'deposit native success'
                    // })
                    // console.log("[跨链][成功] 链 " + item.name, "到链 " + toChainId, "地址 " + recipient, "数额 " + final)
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
                    // console.log(item.chainId, toToken, recipient, final)
                    toContract.withdraw(item.chainId, fromToken, recipient, final).then(_ => {
                        // logInsert({
                        //     depositHash: event.blockHash,
                        //     chainFrom: item.name,
                        //     chainTo: toChain,
                        //     block: numberNow,
                        //     toToken: token,
                        //     address,
                        //     value: _value,
                        //     status: 'deposit success'
                        // })
                        console.log("[跨链][成功] 链 " + item.name, "到链 " + toChainId, "代币 " + toToken, "地址 " + recipient, "数额 " + value)
                    }).catch(error => {
                        // console.log(error)
                        // logInsert({
                        //     chainFrom: item.name,
                        //     chainTo: toChain,
                        //     block: numberNow,
                        //     depositHash: event.blockHash,
                        //     toToken: token,
                        //     address,
                        //     value: _value,
                        //     status: 'deposit error'
                        // })
                        console.log("[跨链][失败] 链 " + item.name, "到链 " + toChainId, "代币 " + toToken, "地址 " + recipient, "数额 " + value)
                    })
                }
                await setChainLock(item.chainId, numberNow.toString())
            }
        })
    })
}

main().then(_ => {

})