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


const logSave = (pairId, recipient, value, fromChain, toChain, depositHash) => {
    const depositTime = Math.round(new Date() / 1000)
    const data = {pairId, recipient, value, fromChain, toChain, depositHash, depositTime}
    connection.query('INSERT INTO log SET ?', data, function (error, results, fields) {
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


const main = async () => {
    const chains = await getChains()
    for (const chain of chains) {
        let isRight = false
        let provider, num
        provider = new ethers.providers.JsonRpcProvider(chain.url)
        num = await provider.getBlockNumber()
        const bridge = new ethers.Contract(chain.bridge, abi.bridge(), provider)
        const logs = await bridge.queryFilter(bridge.filters.Deposit(), chain.syncNumber, num - 2)

        await setSyncNumber(chain.chainId, num - 2)
        isRight = true
    }
}

main().then()