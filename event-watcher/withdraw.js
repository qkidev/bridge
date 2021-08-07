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
                return reject(error)
            } else {
                return resolve(results)
            }
        })
    })
}

const getUnWithdrawLog = () => {
    return new Promise((resolve, reject) => {
        connection.query('SELECT * FROM log_new WHERE `withdrawHash`is null AND `withdrawSubmit` = 0 ORDER BY `id` DESC', (error, results, fields) => {
            if (error) {
                return reject(error)
            } else {
                return resolve(results)
            }
        })
    })
}

const submitWithdraw = (id) => {
    return new Promise((resolve, reject) => {
        connection.query("UPDATE log_new SET `withdrawSubmit` = 1 WHERE `id` = ?", [id], function (error, results, fields) {
            if (error) return reject(error)
            return resolve(results)
        });
    })
}


const getPairById = (id) => {
    return new Promise((resolve, reject) => {
        connection.query('SELECT * FROM `pair` WHERE `id` = ?', [id], (err, res, fields) => {
            if (err) {
                return reject(err)
            } else {
                return resolve(res[0])
            }
        })
    })
}

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


const main = async () => {

    const isCheck = await getIsCheck()

    const chains = await getChains()

    const gweis = {}
    const managers = {}

    for (const chain of chains) {
        gweis[chain.chainId] = chain.gwei + ""
        let isRight = false
        let provider = {}
        while (!isRight) {
            provider = new ethers.providers.JsonRpcProvider(chain.url)
            try {
                const num = await provider.getBlockNumber()
                console.log(num)
                isRight = true
            } catch (e) {
                // console.log(e)
            }
        }
        const wallet = new ethers.Wallet(process.env.PK, provider)
        managers[chain.chainId] = new ethers.Contract(chain['bridge_manager'], abi.bridgeManager(), wallet)

    }

    const logs = await getUnWithdrawLog()
    for (const log of logs) {
        console.log(log['id'])
        const pair = await getPairById(log['pairId'])
        if (!isCheck || pair['limit'] === 0 || log['value'] <= pair['limit']) {
            const manager = managers[log['toChain']]
            if (manager) {
                let isSuccess = false
                let tryNum = 0
                while (!isSuccess) {
                    try {
                        tryNum += 1
                        const amount = ethers.utils.parseUnits(log['amount'], pair['decimal'])
                        const isNative = pair['isNative'] === 1
                        await manager['submitTransaction'](log['fromChain'], log['depositHash'], pair['fromToken'], log['recipient'], amount, isNative, !pair['isMain'], {
                            gasPrice: ethers.utils.parseUnits(gweis[log['toChain']], 'gwei'),
                        })
                        await submitWithdraw(log['id'])
                        console.log("SubmitTransaction")
                        isSuccess = true
                    } catch (e) {
                        console.log(e)
                        console.log(tryNum)
                        if (tryNum > 100) isSuccess = true
                    }
                }
            }
        }
    }
}

main().then()