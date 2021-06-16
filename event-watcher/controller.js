const {
    ethers
} = require('ethers')

const axios = require('axios').default

const fs = require("fs")

require('dotenv').config()

const mysql = require('mysql')

const connection = mysql.createConnection({
    host: process.env.DB_HOST,
    user: process.env.DB_USER,
    password: process.env.DB_PASSWORD,
    database: process.env.DB_DATABASE,
})

connection.connect()

const getBlockLock = (chain) => {
    return new Promise((resolve, reject) => {
        fs.readFile(__dirname + "/lock/" + chain + ".lock", (err, data) => {
            if (err) {
                fs.writeFile(__dirname + "/lock/" + chain + ".lock", "0", _ => {
                    resolve("0")
                })
            } else {
                resolve(data.toString())
            }
        })
    })
}

const getNativeFee = (fromChain, toChain, value) => {
}

const setBlockLock = (chain, number) => {
    fs.writeFile(__dirname + "/lock/" + chain + ".lock", number.toString(), _ => {})
}

// 跨链桥合约部署参数 管理员地址 是为了保持部署地址nonce为初始值 则每个链部署的跨链桥合约地址一样

const tokens = {
    "qk": {
        "0xE06235AAAA893336f1B38e10499Bef7Acb483EF9": {
            fee: 2,
            decimal: 0
        },
        "0x059d5e8be85b0E517F930e00Ddd52AC1Aaf61115": {
            fee: 2,
            decimal: 0
        },
        "0x5Cad8a1340E624Be90adfd787F6F3248DdF17321": {
            fee: 2,
            decimal: 6
        }
    },
    "ropsten": {
        "0xA52191450806181dEe22AAc5fc074C83c057AC43": {
            fee: 2,
            decimal: 0
        },
        "0xb6Cb36df5b88A3e5b547b04368Ebef30f04409a9": {
            fee: 2,
            decimal: 0
        },
        "0x0B8880087d8A76e03802Cd000e36126665e24363": {
            fee: 2,
            decimal: 6
        }
    }
}

// 跨链桥地址
const addressBridges = {
    // "7545": "0x53AF942f88b73f835fB3eE1D48A846Da92029184",
    // "8545": "0x3736Ad67E30cCD77C6740Ea4a8e549c04cE439F2",
    qk: "0xcF70a42585473F160e7F5191dfe97fA15d5D8F3B",
    ropsten: "0xF891Ca04EA0276516A7823Ff787f923934dd56Aa",
    // HECO: "",
    // ETH: "",
    // BSC: "",
    // OKEX: "",
}

// 管理员密钥
const pk = {
    qk: process.env.PK_QK,
    ropsten: process.env.PK_ROPSTEN,
    // "7545": process.env.PK_7545,
    // "8545": process.env.PK_8545,
}

// 跨链桥ABI
const abiBridge = [
    "event Deposit(string chain, address remote, address recipient, uint256 value)",
    "event DepositNative(string chain, address recipient, uint256 value)",
    "event WithdrawDone(address local, address remote, address recipient, uint256 value)",
    "event WithdrawNativeDone(string fromChain, address recipient, uint256 value)",
    "function withdraw(string chain, address remote, address recipient, uint256 value)",
    "function withdrawNative(string memory fromChain, address payable recipient, uint256 value)"
]

// 支持链主网
const urls = [
    // {
    //     name: "7545",
    //     url: "http://127.0.0.1:7545"
    // },
    // {
    //     name: "8545",
    //     url: "http://127.0.0.1:8545"
    // },
    // {
    //     name: "ETH",
    //     url: "https://mainnet.infura.io/v3/#"
    // },
    {
        name: "qk",
        url: "https://hz.node.quarkblockchain.cn"
    },
    {
        name: "ropsten",
        url: "https://ropsten.infura.io/v3/fddb1d98064647dd8bce19afb8b48059"
    },
    // {
    //     name: "HECO",
    //     url: "https://http-mainnet-node.huobichain.com"
    // },
    // {
    //     name: "BSC",
    //     url: "https://bsc-dataseed.binance.org"
    // },
    // {
    //     name: "OKEX",
    //     url: "https://exchainrpc.okex.org"
    //     // https://okexchain-docs.readthedocs.io/en/latest/developers/quick-start-for-mainnet.html#metamask
    // },
]

// 全部链的跨链桥合约
let bridgeContracts = {}

// 全部服务提供者
let providers = {}

// 全部支持链钱包
let wallets = {}

// 设置提供者和钱包
urls.forEach(item => {
    const provider = new ethers.providers.JsonRpcProvider(item.url)
    providers[item.name] = provider
    wallets[item.name] = new ethers.Wallet(pk[item.name], provider)
})

// 设置跨链桥合约
urls.forEach(item => {
    bridgeContracts[item.name] = new ethers.Contract(addressBridges[item.name], abiBridge, wallets[item.name])
})

const getFeeValue = (chain, token, value) => {
    const feePoint = tokens[chain][token].fee
    let feeValue = 0
    if (feePoint) {
        feeValue = Math.ceil(value * (feePoint / 100))
    }
    return feeValue
}

const logInsert = (data) => {
    connection.query('INSERT INTO bridge_logs SET ?', data, (error, results, fields) => {
        if (error) throw error;
    })
}

const logUpdate = (status, hash, chainTo, recipient, value) => {
    connection.query('UPDATE bridge_logs SET status = ?, withdrawHash = ? WHERE withdrawHash is null  and chainTo = ? AND address = ? AND `value` = ?', [status, hash, chainTo, recipient, value], (error, results, fields) => {
        if (error) throw error;
    })
}

async function main() {
    urls.forEach(item => {
        const contract = bridgeContracts[item.name]
        contract.on("WithdrawDone", (local, remote, recipient, value, event) => {
            // console.log(event.blockHash)
            // console.log(item.name, local, recipient, value, event.blockNumber)
            const decimal = tokens[item.name][local].decimal
            let number = ethers.utils.formatUnits(value, decimal)
            number = (number * 1).toString()
            logUpdate('withdraw done', event.blockHash, item.name, recipient, number)
            console.log("[到账][成功] 链 " + item.name, "代币 " + local, "地址 " + recipient, "数额 " + number)
        })

        contract.on("DepositNative", async (toChain, recipient, value, event) => {
            const blockLock = await getBlockLock(item.name)
            const blockNow = event.blockNumber
            if (blockNow > blockLock) {
                const feeValue = getNativeFee(item.name, toChain, value)
                const toContract = bridgeContracts[toChain]
                if (toContract) {
                    toContract.withdrawNative(item.name, recipient, value - feeValue).then(_ => {
                        logInsert({
                            depositHash: event.blockHash,
                            chainFrom: item.name,
                            chainTo: toChain,
                            block: blockNow,
                            toToken: "",
                            address,
                            value: _value,
                            status: 'deposit native success'
                        })
                        console.log("[跨链][成功] 链 " + item.name, "到链 " + token, "地址 " + address, "数额 " + _value)
                    })
                }

                setBlockLock(item.name, blockNow)
            }
        })

        contract.on("Deposit", async (chain, token, address, value, event) => {
            // console.log(item.name,chain,token,address,value,event.blockNumber)
            const decimal = tokens[item.name][token].decimal

            const blockLock = await getBlockLock(item.name)
            const blockNow = event.blockNumber
            if (blockNow > blockLock) {
                let _value = ethers.utils.formatUnits(value, decimal)
                const feeValue = getFeeValue(item.name, token, _value)
                _value -= feeValue
                const toContract = bridgeContracts[chain]
                if (toContract) {
                    const final = ethers.utils.parseUnits(_value.toString(), decimal)
                    toContract.withdraw(item.name, token, address, final).then(_ => {
                        logInsert({
                            depositHash: event.blockHash,
                            chainFrom: item.name,
                            chainTo: chain,
                            block: blockNow,
                            toToken: token,
                            address,
                            value: _value,
                            status: 'deposit success'
                        })
                        console.log("[跨链][成功] 链 " + item.name, "到链 " + chain, "代币 " + token, "地址 " + address, "数额 " + _value)
                    }).catch(error => {
                        logInsert({
                            chainFrom: item.name,
                            chainTo: chain,
                            block: blockNow,
                            depositHash: event.blockHash,
                            toToken: token,
                            address,
                            value: _value,
                            status: 'deposit error'
                        })
                        console.log("[跨链][失败] 链 " + item.name, "到链 " + chain, "代币 " + token, "地址 " + address, "数额 " + _value)
                        // console.log(error)
                    })
                }
                setBlockLock(item.name, blockNow)
            }
        })
    })
}

main().then(_ => {

})