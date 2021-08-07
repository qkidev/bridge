const {ethers} = require('ethers')
require('dotenv').config()
const mysql = require('mysql')

const abi = require("./lib/abi")

const abiBridge = abi.bridge()

const abiManager = abi.bridgeManager()

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
    connection.query('INSERT INTO log_new SET ?', data, function (error, results, fields) {
        if (error) process.exit(0);
    });
}


const getLog = (hash) => {
    return new Promise((resolve, reject) => {
        connection.query('SELECT * FROM `log` WHERE `depositHash` = ?', [hash], (err, res, fields) => {
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
                process.exit(0);
            } else {
                return resolve(res[0])
            }
        })
    })
}

const getPair = (fromChain, toChain, fromToken, toToken) => {
    return new Promise((resolve, reject) => {
        connection.query('SELECT * FROM `pair` WHERE `fromChain` = ? AND `toChain` = ? AND `fromToken` = ? AND `toToken` = ?', [fromChain, toChain, fromToken, toToken], (err, res, fields) => {
            if (err) {
                process.exit(0);
            } else {
                return resolve(res[0])
            }
        })
    })
}

const main = async () => {
    const chains = await getChains()
    for (const chain of chains) {
        let isRight = false
        let provider, num
        provider = new ethers.providers.JsonRpcProvider(chain.url)
        num = await provider.getBlockNumber()
        let toNum
        if ((num - chain['syncNumber']) > 5000) {
            // toNum = chain['syncNumber'] + 5000
            toNum = num - 2
        } else {
            toNum = num - 2
        }
        const bridge = new ethers.Contract(chain.bridge, abi.bridge(), provider)
        const logs = await bridge.queryFilter(bridge.filters.Deposit(), chain['syncNumber'], toNum)
        for (const log of logs) {
            const toChainId = log.args[0].toString()
            const fromToken = log.args[1]
            const toToken = log.args[2]
            const recipient = log.args[3]
            let value = log.args[4].toString()
            const pair = await getPair(chain.chainId, toChainId, fromToken, toToken)
            if (pair) {
                value = (value / 10 ** pair['decimal']).toFixed(pair['decimal'])
                const fee = (value * pair['bridgeFee'] / 100).toFixed(pair['decimal'])
                let amount = ethers.utils.parseUnits((value - fee).toFixed(pair['decimal']), pair['decimal'])
                amount = (amount / 10 ** pair['decimal']).toFixed(pair['decimal'])
                logSave(pair.id, recipient, value, chain.chainId, toChainId, log.transactionHash, fee, amount)
            }

        }
        await setSyncNumber(chain.chainId, toNum)
        isRight = true
    }
}

main().then()