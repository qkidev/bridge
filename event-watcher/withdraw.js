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
        connection.query('SELECT * FROM log WHERE `withdrawSubmit` = 0 AND `overMax` = 0 AND is_fail = 0 AND `withdrawHash`is null ORDER BY `id` DESC', (error, results, fields) => {
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
        connection.query("UPDATE log SET `withdrawSubmit` = 1 WHERE `id` = ?", [id], function (error, results, fields) {
            if (error) return reject(error)
            return resolve(results)
        });
    })
}

const withdrawFail = (id, remark) => {
    return new Promise((resolve, reject) => {
        connection.query("UPDATE log SET `is_fail` = 1, remark = ? WHERE `id` = ?", [remark, id], function (error, results, fields) {
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
    await withdraw()
    setInterval(async () => {
        await withdraw()
    }, 60000)
}

const withdraw = async () => {
    const isCheck = await getIsCheck()

    const chains = await getChains()

    const gweis = {}
    const managers = {}

    for (const chain of chains) {
        gweis[chain.chainId] = chain['manager_gwei'] + ""
        const provider = new ethers.providers.JsonRpcProvider(chain.url)
        const number = await provider.getBlockNumber()
        console.log("[" + chain['name'] + "]" + "当前高度:" + number)
        const wallet = new ethers.Wallet(process.env.PK, provider)
        // let address = await wallet.getAddress();
        // console.log("管理员地址:" + address)
        managers[chain.chainId] = new ethers.Contract(chain['bridge_manager'], abi.bridgeManager(), wallet)
    }

    const logs = await getUnWithdrawLog()
    for (const log of logs) {
        // 60秒以内的跳过
        const now = Math.round(new Date() / 1000)
        if ((now - log['depositTime']) < 60) continue

        const pair = await getPairById(log['pairId'])

        let overLimit = log['value'] > pair['limit']
        if (pair['limit'] === 0) overLimit = false
        if (!isCheck) overLimit = false

        if (!overLimit && !pair['isStop']) {
            const manager = managers[log['toChain']]
            if (manager) {
                let isSuccess = false
                let tryNum = 0
                console.log("开始处理" + " [logId: " + log['id'] + "]")
                while (!isSuccess) {
                    try {
                        tryNum += 1
                        const amount = ethers.utils.parseUnits(log['amount'], pair['decimal'])
                        const isNative = pair['isNative'] === 1
                        let gasPrice = ethers.utils.parseUnits(gweis[log['toChain']], 'gwei')
                        const tx = await manager['submitTransaction'](log['fromChain'], log['depositHash'], pair['fromToken'], log['recipient'], amount, isNative, !pair['isMain'], {
                            gasPrice: gasPrice,
                            gasLimit: '40000'
                        })
                        await tx.wait()
                        await submitWithdraw(log['id'])
                        console.log("SubmitTransaction")
                        isSuccess = true
                    } catch (e) {
                        console.log("[logId: " + log['id'] + "]" + "第" + tryNum + "次重试")
                        if (tryNum > 5) {
                            isSuccess = true
                            let _err = e.error?.error?.message || e.error?.data?.message || 'error'
                            let remark = "[" + e.code + "] error: " + e.reason +"; "+ _err;
                            console.log(remark)
                            await withdrawFail(log['id'], remark)
                        }
                    }
                }
            }
        }
    }
}

main().then()