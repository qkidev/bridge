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
    const depositTime = Math.round(new Date() / 1000)
    const data = {pairId, recipient, value, fromChain, toChain, depositHash, depositTime}
    connection.query('INSERT INTO log SET ?', data, function (error, results, fields) {
        if (error) throw error;
    });
}

const withdrawDone = (depositHash,withdrawHash) => {
    const time =Math.round(new Date() / 1000)
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
                "name": "hash",
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
    chains.forEach(item => {
        const provider = new ethers.providers.JsonRpcProvider(item.url)
        const wallet = new ethers.Wallet(process.env.PK, provider)
        bridgeContracts[item.chainId] = new ethers.Contract(item.bridge, abiBridge, provider)
        managerContracts[item.chainId] = new ethers.Contract(item.bridge_manager, abiManager, wallet)
    })

    chains.forEach(item => {
        const contract = bridgeContracts[item.chainId]

        // 主网币跨入成功
        contract.on("WithdrawNativeDone", async (toChainId, recipient, isMain, value, depositHash,event) => {
            withdrawDone(depositHash,event.transactionHash)
        })

        // 代币跨入成功
        contract.on("WithdrawDone", async (toChainId, fromToken, toToken, recipient, value,depositHash, event) => {
            // console.log(depositHash,event.transactionHash)
            // toChainId = toChainId.toString()
            // console.log (toChainId, fromToken, toToken, recipient, value,event.transactionHash)
            withdrawDone(depositHash,event.transactionHash)
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
                if (!isCheck || value <= pair['limit']) {
                    const manager = managerContracts[toChainId]
                    if (manager) {
                        const fee = Math.ceil(value * pair['bridgeFee'] / 100)
                        const final = ethers.utils.parseUnits((value - fee).toString(), pair['decimal'])
                        try {
                            await manager['withdrawNative'](item.chainId, recipient, !isMain, final, event.transactionHash)
                        } catch (e) {
                            console.log(e)
                        }
                    }
                }
            }
        })

        // 执行代币跨出操作
        contract.on("Deposit", async (toChainId, fromToken, toToken, recipient, value, event) => {
            value = value.toString() * 1
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
                if (!isCheck || pair['limit'] === 0 || value <= pair['limit']) {
                    const manager = managerContracts[toChainId]
                    if (manager) {
                        const fee = Math.ceil(value * pair['bridgeFee'] / 100)
                        const final = ethers.utils.parseUnits((value - fee).toString(), pair['decimal'])
                        // console.log(item.chainId, fromToken, recipient, final.toString(), event.transactionHash)
                        manager['withdraw'](item.chainId, fromToken, recipient, final, event.transactionHash)
                    }
                }
            }
        })
    })
}

main().then()