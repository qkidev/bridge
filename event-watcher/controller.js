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

// 取系统配置跨链审核的值
const getIsCheck = () => {
    return new Promise((resolve, reject) => {
        connection.query('SELECT * FROM `setting` WHERE `name` = ?', [`withdraw-check`], (err, res, fields) => {
            if (err) {
                console.log("GetIsCheckError")
                process.exit(0)
            } else {
                return resolve(res[0].value)
            }
        })
    })
}

const getChains = () => {
    return new Promise((resolve, reject) => {
        connection.query('SELECT * FROM chain WHERE `status`= 1', (error, results, fields) => {
            if (error) {
                console.log("GetChainsError")
                process.exit(0)
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
                console.log("GetPairError")
                process.exit(0)
            } else {
                return resolve(res[0])
            }
        })
    })
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
                console.log("GetPairNativeError")
                process.exit(0)
            } else {
                return resolve(res[0])
            }
        })
    })
}

const logSave = (pairId, recipient, value, fromChain, toChain, depositHash, fee, amount, overMax) => {
    const depositTime = Math.round(new Date() / 1000)
    const data = {pairId, recipient, value, fromChain, toChain, depositHash, depositTime, fee, amount, overMax}
    connection.query('INSERT INTO log SET ?', data, function (error, results, fields) {
        if (error) {
            console.log("LogSaveError")
            process.exit(0)
        }
    });
}

const withdrawDone = (depositHash, withdrawHash) => {
    const time = Math.round(new Date() / 1000)
    connection.query("UPDATE log SET withdrawHash = ?, withdrawTime = ? WHERE depositHash = ?", [withdrawHash, time, depositHash], function (error, results, fields) {
        if (error) {
            // console.log(error)
            process.exit(0)
        }
    })
}

const getGwei = (chainId) => {
    return new Promise((resolve, reject) => {
        connection.query('SELECT * FROM `chain` WHERE `chainId` = ?', [chainId], (err, res, fields) => {
            if (err) {
                process.exit(0);
            } else {
                return resolve(res[0]['manager_gwei'])
            }
        })
    })
}

const submitWithdraw = (depositHash) => {
    return new Promise((resolve, reject) => {
        connection.query("UPDATE log SET `withdrawSubmit` = 1 WHERE `depositHash` = ?", [depositHash], function (error, results, fields) {
            if (error) return reject(error)
            return resolve(results)
        });
    })
}


// 跨链桥ABI
const abiBridge = abi.bridge()

// 管理员合约ABI
const abiManager = abi.bridgeManager()

// 全部链的跨链桥合约
let bridgeContracts = {}

let managerContracts = {}


async function main() {

    // 支持链主网
    const chains = await getChains()


    // 设置提供者和钱包和跨链桥合约
    for (const item of chains) {
        let isRight = false
        let provider = {}
        while (!isRight) {
            provider = new ethers.providers.JsonRpcProvider(item.url)
            try {
                const num = await provider.getBlockNumber()
                console.log(num)
                isRight = true
            } catch (e) {
            }
        }

        const wallet = new ethers.Wallet(process.env.PK, provider)
        bridgeContracts[item.chainId] = new ethers.Contract(item.bridge, abiBridge, provider)
        managerContracts[item.chainId] = new ethers.Contract(item.bridge_manager, abiManager, wallet)

    }


    chains.forEach(item => {
        const contract = bridgeContracts[item.chainId]

        // 主网币跨入成功
        contract.on("WithdrawNativeDone", async (toChainId, recipient, isMain, value, depositHash, event) => {
            console.log("WithdrawNativeDone")
            withdrawDone(depositHash, event.transactionHash)
        })

        // 代币跨入成功
        contract.on("WithdrawDone", async (toChainId, fromToken, toToken, recipient, value, depositHash, event) => {
            console.log("WithdrawDone")
            // console.log(depositHash,event.transactionHash)
            // toChainId = toChainId.toString()
            // console.log (toChainId, fromToken, toToken, recipient, value,event.transactionHash)
            withdrawDone(depositHash, event.transactionHash)
        })

        // 主网币跨出
        contract.on("DepositNative", async (toChainId, isMain, recipient, value, event) => {
            console.log("DepositNative")
            toChainId = toChainId.toString()
            isMain = isMain ? 1 : 0
            const pair = await getPairNative(item.chainId, toChainId, isMain)
            value = ethers.utils.formatEther(value).toString() * 1
            let overMax = value > pair['maxValue'] * 1 ? 1 : 0
            if (pair['maxValue'] * 1 === 0) overMax = 0

            let fee = (value * pair['bridgeFee'] / 100).toFixed(pair['decimal'])

            const haveFeeMin = pair['feeMin'] * 1 > 0
            const haveFeeMax = pair['feeMax'] * 1 > 0

            if (haveFeeMin) {
                if (fee < pair['feeMin']) fee = pair['feeMin'] * 1
            }
            if (haveFeeMax) {
                if (fee > pair['feeMax']) fee = pair['feeMax'] * 1
            }

            const amount = ethers.utils.parseUnits((value - fee).toFixed(pair['decimal']), pair['decimal'])
            const amountStr = ethers.utils.formatEther(amount).toString()

            const log = await getLog(event.transactionHash)

            if (typeof (log) === "undefined") {
                logSave(pair.id, recipient, value, item.chainId, toChainId, event.transactionHash, fee.toString(), amountStr, overMax)
            }
            const isCheck = await getIsCheck()

            let overLimit = value > pair['limit']
            if (pair['limit'] === 0) overLimit = false
            if (!isCheck) overLimit = false

            // 检查配置的审核状态
            if (!overLimit && !overMax) {
                const manager = managerContracts[toChainId]
                if (manager) {
                    let isSuccess = false
                    let tryNum = 0
                    while (!isSuccess) {
                        try {
                            tryNum += 1
                            const gwei = await getGwei(toChainId)
                            const tx = await manager['submitTransaction'](item.chainId, event.transactionHash, "0x0000000000000000000000000000000000000000", recipient, amount, true, !isMain, {
                                gasPrice: ethers.utils.parseUnits(gwei + "", 'gwei'),
                            })
                            await tx.wait()
                            await submitWithdraw(event.transactionHash)
                            console.log("SubmitTransaction")
                            isSuccess = true
                        } catch (e) {
                            console.log(tryNum)
                            if (tryNum > 10) isSuccess = true
                        }
                    }
                }
            }
        })

        // 代币跨出
        contract.on("Deposit", async (toChainId, fromToken, toToken, recipient, value, event) => {
            console.log("Deposit")
            try {
                toChainId = toChainId.toString()
                const pair = await getPair(item.chainId, toChainId, fromToken, toToken)
                value = (value.toString() / 10 ** pair['decimal']).toFixed(pair['decimal'])


                if (pair['tokenFee'] * 1 > 0) {
                    value = value * (1 - (pair['tokenFee'] / 100)).toFixed(pair['decimal'])
                }


                let overMax = value > pair['maxValue'] ? 1 : 0
                if (pair['maxValue'] * 1 === 0) overMax = 0

                const haveFeeMin = pair['feeMin'] * 1 > 0
                const haveFeeMax = pair['feeMax'] * 1 > 0

                let fee = (value * pair['bridgeFee'] / 100).toFixed(pair['decimal'])

                if (haveFeeMin) {
                    if (fee < pair['feeMin']) fee = pair['feeMin'] * 1
                }
                if (haveFeeMax) {
                    if (fee > pair['feeMax']) fee = pair['feeMax'] * 1
                }

                const amountStr = (value - fee).toFixed(pair['decimal'])
                const amount = ethers.utils.parseUnits(amountStr, pair['decimal'])

                const log = await getLog(event.transactionHash)

                if (typeof (log) === "undefined") {
                    logSave(pair.id, recipient, value, item.chainId, toChainId, event.transactionHash, fee.toString(), amountStr, overMax)
                }
                const isCheck = await getIsCheck()

                let overLimit = value > pair['limit']
                if (pair['limit'] === 0) overLimit = false
                if (!isCheck) overLimit = false

                // 检查配置的审核状态
                if (!overLimit && !overMax && !pair['isStop']) {
                    const manager = managerContracts[toChainId]
                    if (manager) {
                        let isSuccess = false
                        let tryNum = 0
                        while (!isSuccess) {
                            try {
                                tryNum += 1
                                // console.log(item.chainId, event.transactionHash, fromToken, recipient, amount.toString(),false,!pair['isMain'])
                                const gwei = await getGwei(toChainId)
                                const tx = await manager['submitTransaction'](item.chainId, event.transactionHash, fromToken, recipient, amount, false, !pair['isMain'], {
                                    gasPrice: ethers.utils.parseUnits(gwei + "", 'gwei'),
                                    // gasLimit: 200000
                                })
                                await tx.wait()
                                await submitWithdraw(event.transactionHash)
                                console.log("SubmitTransaction")
                                // console.log(event.transactionHash, tx.hash)
                                isSuccess = true
                            } catch (e) {
                                console.log(e.message)
                                console.log(tryNum)
                                if (tryNum > 10) isSuccess = true
                            }
                        }
                    }
                }
            } catch (e) {
                console.log(e.message)
            }
        })
    })
}

main().then()