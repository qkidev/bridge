const {ethers} = require('ethers')
require('dotenv').config()
const mysql = require('mysql')

const connection = mysql.createConnection({
    host: process.env.DB_HOST,
    user: process.env.DB_USER,
    password: process.env.DB_PASSWORD,
    database: process.env.DB_DATABASE,
})

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

const missingWithdraw = () => {
    setInterval(async () => {
        connection.query('SELECT * FROM `log` WHERE `withdrawHash` IS NULL ORDER BY `id` desc', null, (err, res) => {
            // res[0] = log
            if (!err) {
                connection.query('SELECT * FROM `chain` WHERE `chainId` = ?', [res[0].toChain], (err2, res2) => {
                    // res2[0] = chain
                    if (!err2) {
                        connection.query('SELECT * FROM `pair` WHERE `id` = ?', [res[0].pairId], async (err3, res3) => {
                            // res3[0] = pair
                            if (!err3) {
                                try {
                                    const provider = new ethers.providers.JsonRpcProvider(res2[0].url)
                                    const wallet = new ethers.Wallet(process.env.PK, provider)
                                    const manager = new ethers.Contract(res2[0]['bridge_manager'], abiManager, wallet)
                                    const fee = Math.ceil(res[0].value * res3[0]['bridgeFee'] / 100)
                                    // console.log(fee)
                                    const final = ethers.utils.parseUnits((res[0].value - fee).toString(), res3[0]['decimal'])
                                    // console.log(final)
                                    console.log(res[0].fromChain, res3[0]['fromToken'], res[0].recipient, final, res[0].depositHash)
                                    const tx = await manager.withdraw(res[0].fromChain, res3[0]['fromToken'], res[0].recipient, final, res[0].depositHash)
                                    await tx.wait()
                                    console.log(tx,'ok')
                                }catch (e) {
                                    console.log('e')
                                }
                            }
                        })
                    }
                })
            }
        })

    }, 1000)
}

missingWithdraw()