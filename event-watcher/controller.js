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
                return reject(err)
            } else {
                return resolve(res[0])
            }
        })
    })
}

const logSave = (pairId, recipient, value, fromChain, toChain, depositHash) => {
    const depositTime = Math.round(new Date() / 1000)
    const data = {pairId, recipient, value, fromChain, toChain, depositHash, depositTime}
    connection.query('INSERT INTO log SET ?', data, function (error, results, fields) {
        if (error) throw error;
    });
}

const withdrawDone = (depositHash, withdrawHash) => {
    const time = Math.round(new Date() / 1000)
    connection.query("UPDATE log SET withdrawHash = ?, withdrawTime = ? WHERE depositHash = ?", [withdrawHash, time, depositHash], function (error, results, fields) {
        if (error) throw error;
    });

}

// 跨链桥ABI
const abiBridge = [
    {
        "inputs": [],
        "stateMutability": "nonpayable",
        "type": "constructor"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "toChainId",
                "type": "uint256"
            },
            {
                "indexed": false,
                "internalType": "address",
                "name": "fromToken",
                "type": "address"
            },
            {
                "indexed": false,
                "internalType": "address",
                "name": "toToken",
                "type": "address"
            },
            {
                "indexed": false,
                "internalType": "address",
                "name": "recipient",
                "type": "address"
            },
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "value",
                "type": "uint256"
            }
        ],
        "name": "Deposit",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "toChainId",
                "type": "uint256"
            },
            {
                "indexed": false,
                "internalType": "bool",
                "name": "isMain",
                "type": "bool"
            },
            {
                "indexed": false,
                "internalType": "address",
                "name": "recipient",
                "type": "address"
            },
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "value",
                "type": "uint256"
            }
        ],
        "name": "DepositNative",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "toChainId",
                "type": "uint256"
            },
            {
                "indexed": false,
                "internalType": "address",
                "name": "fromToken",
                "type": "address"
            },
            {
                "indexed": false,
                "internalType": "address",
                "name": "toToken",
                "type": "address"
            },
            {
                "indexed": false,
                "internalType": "address",
                "name": "recipient",
                "type": "address"
            },
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "value",
                "type": "uint256"
            },
            {
                "indexed": false,
                "internalType": "bytes",
                "name": "depositHash",
                "type": "bytes"
            }
        ],
        "name": "WithdrawDone",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "fromChainId",
                "type": "uint256"
            },
            {
                "indexed": false,
                "internalType": "address",
                "name": "recipient",
                "type": "address"
            },
            {
                "indexed": false,
                "internalType": "bool",
                "name": "isMain",
                "type": "bool"
            },
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "value",
                "type": "uint256"
            },
            {
                "indexed": false,
                "internalType": "bytes",
                "name": "depositHash",
                "type": "bytes"
            }
        ],
        "name": "WithdrawNativeDone",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "address",
                "name": "_address",
                "type": "address"
            }
        ],
        "name": "adminChanged",
        "type": "event"
    },
    {
        "inputs": [],
        "name": "admin",
        "outputs": [
            {
                "internalType": "address",
                "name": "",
                "type": "address"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "chainId",
                "type": "uint256"
            },
            {
                "internalType": "address",
                "name": "toToken",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "value",
                "type": "uint256"
            }
        ],
        "name": "deposit",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "toChainId",
                "type": "uint256"
            },
            {
                "internalType": "bool",
                "name": "isMain",
                "type": "bool"
            },
            {
                "internalType": "uint256",
                "name": "value",
                "type": "uint256"
            }
        ],
        "name": "depositNative",
        "outputs": [],
        "stateMutability": "payable",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "manager",
        "outputs": [
            {
                "internalType": "address",
                "name": "",
                "type": "address"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "toChainId",
                "type": "uint256"
            },
            {
                "internalType": "bool",
                "name": "isRun",
                "type": "bool"
            },
            {
                "internalType": "bool",
                "name": "isMain",
                "type": "bool"
            },
            {
                "internalType": "address",
                "name": "fromToken",
                "type": "address"
            }
        ],
        "name": "nativeInsert",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address payable",
                "name": "recipient",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "value",
                "type": "uint256"
            }
        ],
        "name": "nativeTransfer",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            },
            {
                "internalType": "bool",
                "name": "",
                "type": "bool"
            }
        ],
        "name": "natives",
        "outputs": [
            {
                "internalType": "bool",
                "name": "isRun",
                "type": "bool"
            },
            {
                "internalType": "bool",
                "name": "isMain",
                "type": "bool"
            },
            {
                "internalType": "address",
                "name": "local",
                "type": "address"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "owner",
        "outputs": [
            {
                "internalType": "address",
                "name": "",
                "type": "address"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address payable",
                "name": "newAdmin",
                "type": "address"
            }
        ],
        "name": "setAdmin",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "_manager",
                "type": "address"
            }
        ],
        "name": "setManager",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "toChainId",
                "type": "uint256"
            },
            {
                "internalType": "bool",
                "name": "isMain",
                "type": "bool"
            },
            {
                "internalType": "bool",
                "name": "state",
                "type": "bool"
            }
        ],
        "name": "setNativeIsRun",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address payable",
                "name": "newOwner",
                "type": "address"
            }
        ],
        "name": "setOwner",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "toChainId",
                "type": "uint256"
            },
            {
                "internalType": "address",
                "name": "toToken",
                "type": "address"
            },
            {
                "internalType": "bool",
                "name": "state",
                "type": "bool"
            }
        ],
        "name": "setTokenIsRun",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "toChainId",
                "type": "uint256"
            },
            {
                "internalType": "address",
                "name": "toToken",
                "type": "address"
            },
            {
                "internalType": "address",
                "name": "fromToken",
                "type": "address"
            },
            {
                "internalType": "bool",
                "name": "isRun",
                "type": "bool"
            },
            {
                "internalType": "bool",
                "name": "isMain",
                "type": "bool"
            }
        ],
        "name": "tokenInsert",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "fromToken",
                "type": "address"
            },
            {
                "internalType": "address",
                "name": "recipient",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "value",
                "type": "uint256"
            }
        ],
        "name": "tokenTransfer",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            },
            {
                "internalType": "address",
                "name": "",
                "type": "address"
            }
        ],
        "name": "tokens",
        "outputs": [
            {
                "internalType": "bool",
                "name": "isRun",
                "type": "bool"
            },
            {
                "internalType": "bool",
                "name": "isMain",
                "type": "bool"
            },
            {
                "internalType": "address",
                "name": "local",
                "type": "address"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "toChainId",
                "type": "uint256"
            },
            {
                "internalType": "address",
                "name": "toToken",
                "type": "address"
            },
            {
                "internalType": "address",
                "name": "recipient",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "value",
                "type": "uint256"
            },
            {
                "internalType": "bytes",
                "name": "depositHash",
                "type": "bytes"
            }
        ],
        "name": "withdraw",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "toChainId",
                "type": "uint256"
            },
            {
                "internalType": "address payable",
                "name": "recipient",
                "type": "address"
            },
            {
                "internalType": "bool",
                "name": "isMain",
                "type": "bool"
            },
            {
                "internalType": "uint256",
                "name": "value",
                "type": "uint256"
            },
            {
                "internalType": "bytes",
                "name": "depositHash",
                "type": "bytes"
            }
        ],
        "name": "withdrawNative",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "stateMutability": "payable",
        "type": "receive"
    }
]

// 管理员合约ABI
const abiManager = [
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "_signLimit",
                "type": "uint256"
            },
            {
                "internalType": "address",
                "name": "_bridgeAddress",
                "type": "address"
            }
        ],
        "stateMutability": "nonpayable",
        "type": "constructor"
    },
    {
        "inputs": [],
        "name": "bridgeAddress",
        "outputs": [
            {
                "internalType": "address",
                "name": "",
                "type": "address"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "bytes",
                "name": "",
                "type": "bytes"
            }
        ],
        "name": "isComplete",
        "outputs": [
            {
                "internalType": "bool",
                "name": "",
                "type": "bool"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "",
                "type": "address"
            }
        ],
        "name": "isManager",
        "outputs": [
            {
                "internalType": "bool",
                "name": "",
                "type": "bool"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "_address",
                "type": "address"
            }
        ],
        "name": "managerAdd",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "_address",
                "type": "address"
            }
        ],
        "name": "managerDel",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            },
            {
                "internalType": "bytes",
                "name": "",
                "type": "bytes"
            },
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "name": "multiSigns",
        "outputs": [
            {
                "internalType": "address",
                "name": "",
                "type": "address"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "owner",
        "outputs": [
            {
                "internalType": "address",
                "name": "",
                "type": "address"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "_bridgeAddress",
                "type": "address"
            }
        ],
        "name": "setBridgeAddress",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address payable",
                "name": "_owner",
                "type": "address"
            }
        ],
        "name": "setOwner",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "num",
                "type": "uint256"
            }
        ],
        "name": "setSignLimit",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "signLimit",
        "outputs": [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "toChainId",
                "type": "uint256"
            },
            {
                "internalType": "address",
                "name": "toToken",
                "type": "address"
            },
            {
                "internalType": "address",
                "name": "recipient",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "value",
                "type": "uint256"
            },
            {
                "internalType": "bytes",
                "name": "depositHash",
                "type": "bytes"
            }
        ],
        "name": "withdraw",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "toChainId",
                "type": "uint256"
            },
            {
                "internalType": "address payable",
                "name": "recipient",
                "type": "address"
            },
            {
                "internalType": "bool",
                "name": "isMain",
                "type": "bool"
            },
            {
                "internalType": "uint256",
                "name": "value",
                "type": "uint256"
            },
            {
                "internalType": "bytes",
                "name": "hash",
                "type": "bytes"
            }
        ],
        "name": "withdrawNative",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    }
]

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
            value = (value.toString() / 10 ** pair['decimal']).toFixed(pair['decimal'])
            const max = pair.maxValue || 100000000000
            if (value >= pair.minValue && value <= max) {
                const log = await getLog(event.transactionHash)
                if (typeof (log) === "undefined") {
                    logSave(pair.id, recipient, value, item.chainId, toChainId, event.transactionHash)
                    // console.log(`[主网币][跨链][成功] ${value} 个 ${pair.name} 从 ${pair.fromChain} 到 ${pair.toChain}`)
                }
                const isCheck = await getIsCheck()
                // 检查配置的审核状态
                if (!isCheck || pair['limit'] === 0 || value <= pair['limit']) {
                    const manager = managerContracts[toChainId]
                    if (manager) {
                        let isSuccess = false
                        let tryNum = 0
                        while (!isSuccess) {
                            try {
                                tryNum += 1
                                const fee = (value * pair['bridgeFee'] / 100).toFixed(pair['decimal'])
                                const amount = ethers.utils.parseUnits((value - fee).toFixed(pair['decimal']), pair['decimal'])
                                await manager['submitTransaction'](item.chainId, event.transactionHash, "0x0000000000000000000000000000000000000000", recipient, amount)
                                isSuccess = true
                            } catch (e) {
                                // console.log(e)
                            }
                        }
                    }
                }
            }
        })

        // 执行代币跨出操作
        contract.on("Deposit", async (toChainId, fromToken, toToken, recipient, value, event) => {
            console.log("Deposit")
            try {
                toChainId = toChainId.toString()
                const pair = await getPair(item.chainId, toChainId, fromToken, toToken)
                value = (value.toString() / 10 ** pair['decimal']).toFixed(pair['decimal'])
                const max = pair.maxValue || 100000000000
                if (value >= pair.minValue && value <= max) {
                    const log = await getLog(event.transactionHash)
                    if (typeof (log) === "undefined") {
                        logSave(pair.id, recipient, value, item.chainId, toChainId, event.transactionHash)
                    }
                    const isCheck = await getIsCheck()
                    // 检查配置的审核状态
                    if (!isCheck || pair['limit'] === 0 || value <= pair['limit']) {
                        const manager = managerContracts[toChainId]
                        if (manager) {
                            let isSuccess = false
                            let tryNum = 0
                            while (!isSuccess) {
                                try {
                                    tryNum += 1
                                    const fee = (value * pair['bridgeFee'] / 100).toFixed(pair['decimal'])
                                    const amount = ethers.utils.parseUnits((value - fee).toFixed(pair['decimal']), pair['decimal'])
                                    // console.log(item.chainId, fromToken, recipient, amount.toString(), event.transactionHash)
                                    const tx = await manager['submitTransaction'](item.chainId, event.transactionHash, fromToken, recipient, amount)
                                    console.log(event.transactionHash, tx.hash)
                                    isSuccess = true
                                } catch (e) {
                                    console.log(tryNum)
                                }
                            }
                        }
                    }
                }
            } catch (e) {
                console.log(e)
            }
        })
    })
}

main().then()