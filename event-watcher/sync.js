const {ethers} = require('ethers')
require('dotenv').config()
const mysql = require('mysql')

const abi = require("./lib/abi")

const connection = mysql.createConnection({
    host: process.env.DB_HOST,
    user: process.env.DB_USER,
    password: process.env.DB_PASSWORD,
    database: process.env.DB_DATABASE,
})

connection.connect()

// console.log(process.env)

// console.log(process.argv)

const getChains = () => {
    return new Promise((resolve, reject) => {
        connection.query('SELECT * FROM chain WHERE `status`= 1', (error, results, fields) => {
            if (error) {
                process.exit(0);
            } else {
                return resolve(results)
            }
        })
    })
}


const setSyncNumber = (chainId, number) => {
    connection.query("UPDATE chain SET syncNumber = ? WHERE chainId = ?", [number, chainId], function (error, results, fields) {
        if (error) process.exit(0);
    });
}


const logSave = (pairId, recipient, value, fromChain, toChain, depositHash, fee, amount) => {
    const depositTime = Math.round(new Date() / 1000)
    const data = {pairId, recipient, value, fromChain, toChain, depositHash, depositTime, fee, amount}
    connection.query('INSERT INTO log SET ?', data, function (error, results, fields) {
        if (error) process.exit(0);
    });
}


const getLog = (hash) => {
    return new Promise((resolve, reject) => {
        connection.query('SELECT * FROM `log` WHERE `depositHash` = ?', [hash], (error, res, fields) => {
            if (error) {
                return reject(error)
            } else {
                return resolve(res[0])
            }
        })
    })
}

const getPairNative = (fromChainId, toChainId, isMain) => {
    return new Promise((resolve, reject) => {
        connection.query('SELECT * FROM `pair` WHERE `fromChain` = ? AND `toChain` = ? AND `isMain` = ? AND `isNative` = 1', [fromChainId, toChainId, isMain], (error, res, fields) => {
            if (error) {
                return reject(error)
            } else {
                return resolve(res[0])
            }
        })
    })
}

const getPair = (fromChain, toChain, fromToken, toToken) => {
    return new Promise((resolve, reject) => {
        connection.query('SELECT * FROM `pair` WHERE `fromChain` = ? AND `toChain` = ? AND `fromToken` = ? AND `toToken` = ?', [fromChain, toChain, fromToken, toToken], (error, res, fields) => {
            if (error) {
                return reject(error)
            } else {
                return resolve(res[0])
            }
        })
    })
}

const withdrawDone = (depositHash, withdrawHash) => {
    const time = Math.round(new Date() / 1000)
    connection.query("UPDATE log SET withdrawHash = ?, withdrawTime = ? WHERE depositHash = ?", [withdrawHash, time, depositHash], function (error, results, fields) {
        if (error) {
            console.log(error)
        }
    })
}


const main = async () => {
    await sync()
    setInterval(async () => {
        await sync()
    }, 60000)
}

const sync = async () => {
    const chains = await getChains()
    for (const chain of chains) {
        let provider, num
        provider = new ethers.providers.JsonRpcProvider(chain.url)
        num = await provider.getBlockNumber()
        let toNum = chain['syncLimit'] > 0 ? chain['syncNumber'] + chain['syncLimit'] : num
        if (toNum >= num) toNum = num - 2
        if ((toNum - chain['syncNumber']) < 5) continue
        const bridge = new ethers.Contract(chain.bridge, abi.bridge(), provider)

        const depositLogs = await bridge.queryFilter(bridge.filters.Deposit(), chain['syncNumber'], toNum)
        for (const log of depositLogs) {
            const exist = await getLog(log.transactionHash)
            if (exist) continue
            const toChainId = log.args[0].toString()
            const fromToken = log.args[1]
            const toToken = log.args[2]
            const recipient = log.args[3]
            let value = log.args[4].toString()
            const pair = await getPair(chain.chainId, toChainId, fromToken, toToken)
            if (pair) {
                value = (value / 10 ** pair['decimal']).toFixed(pair['decimal'])
                let fee = (value * pair['bridgeFee'] / 100).toFixed(pair['decimal'])
                const haveFeeMin = pair['feeMin'] * 1 > 0
                const haveFeeMax = pair['feeMax'] * 1 > 0

                if (haveFeeMin) {
                    if (fee < pair['feeMin']) fee = pair['feeMin'] * 1
                }
                if (haveFeeMax) {
                    if (fee > pair['feeMax']) fee = pair['feeMax'] * 1
                }

                let amount = ethers.utils.parseUnits((value - fee).toFixed(pair['decimal']), pair['decimal'])
                amount = (amount / 10 ** pair['decimal']).toFixed(pair['decimal'])
                logSave(pair.id, recipient, value, chain.chainId, toChainId, log.transactionHash, fee, amount)
            }
        }

        const depositNativeLogs = await bridge.queryFilter(bridge.filters.DepositNative(), chain['syncNumber'], toNum)
        for (const log of depositNativeLogs) {
            const exist = await getLog(log.transactionHash)
            if (exist) continue
            const toChainId = log.args[0].toString()
            const isMain = log.args[1]
            let recipient = log.args[2]
            let value = log.args[3].toString()
            const pair = await getPairNative(chain.chainId, toChainId, isMain)
            if (pair) {
                value = (value / 10 ** pair['decimal']).toFixed(pair['decimal'])
                let fee = (value * pair['bridgeFee'] / 100).toFixed(pair['decimal'])
                const haveFeeMin = pair['feeMin'] * 1 > 0
                const haveFeeMax = pair['feeMax'] * 1 > 0

                if (haveFeeMin) {
                    if (fee < pair['feeMin']) fee = pair['feeMin'] * 1
                }
                if (haveFeeMax) {
                    if (fee > pair['feeMax']) fee = pair['feeMax'] * 1
                }

                let amount = ethers.utils.parseUnits((value - fee).toFixed(pair['decimal']), pair['decimal'])
                amount = (amount / 10 ** pair['decimal']).toFixed(pair['decimal'])
                logSave(pair.id, recipient, value, chain.chainId, toChainId, log.transactionHash, fee, amount)
            }
        }

        const withdrawDownLogs = await bridge.queryFilter(bridge.filters.WithdrawDone(), chain['syncNumber'], toNum)
        for (const log of withdrawDownLogs) {
            const depositHash = log.args[5]
            const withdrawHash = log.transactionHash
            withdrawDone(depositHash, withdrawHash)
        }

        const withdrawNativeDownLogs = await bridge.queryFilter(bridge.filters.WithdrawNativeDone(), chain['syncNumber'], toNum)
        for (const log of withdrawNativeDownLogs) {
            const depositHash = log.args[4]
            const withdrawHash = log.transactionHash
            withdrawDone(depositHash, withdrawHash)
        }


        console.log("新增 链id:"+chain.chainId+" 数量:"+ toNum)
        await setSyncNumber(chain.chainId, toNum)
    }
}

main().then()
