const {
    ethers
} = require('ethers')

const fs = require("fs")
const path = require('path')

require('dotenv').config()

const mysql = require('mysql')

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

const setBlockLock = (chain, number) => {
    fs.writeFile(__dirname + "/lock/" + chain + ".lock", number.toString(), _ => {
    })
}

const connection = mysql.createConnection({
    host: process.env.DB_HOST,
    username: process.env.DB_USER,
    password: process.env.DB_PASSWORD,
    database: process.env.DB_DATABASE,
})

// connection.connect()

// 跨链桥合约部署参数 管理员地址 是为了保持部署地址nonce为初始值 则每个链部署的跨链桥合约地址一样

// 跨链桥地址
const addressBridges = {
    "7545": "0x6028c7291Ce8779364253417bBA6CEe24DC10115",
    "8545": "0x84c1B4327B6c7904fC6B96F6E1D33bf2669a9125",
    // qk: "",
    // rop: "",
    // HECO: "",
    // ETH: "",
    // BSC: "",
    // OKEX: "",
}

// 管理员密钥
const pk = {
    // qk: "",
    // rop: "",
    "7545": "da5eb18de6d2773c92c73cdb6a18481ac219780bd5680acd47710e0346b48c6f",
    "8545": "9e2660017e5673db80ce6dc3d4bf89c2928cb4697570dbf01559b065e0eb6c72"
}

// 跨链桥ABI
const abiBridge = [
    "event Deposit(string chain, address remote, address recipient, uint256 value)",
    "event WithdrawDone(address local, address recipient, uint256 value)",
    "function withdraw(string chain, address remote, address recipient, uint256 value)",
]

// 支持链主网
const urls = [
    {
        name: "7545",
        url: "http://127.0.0.1:7545"
    },
    {
        name: "8545",
        url: "http://127.0.0.1:8545"
    },
    // {
    //     name: "ETH",
    //     url: "https://mainnet.infura.io/v3/#"
    // },
    // {
    //     name: "QK",
    //     url: "http://sg.node.quarkblockchain.org"
    // },
    // {
    //     name: "ROP",
    //     url: "https://ropsten.infura.io/v3/#"
    // },
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

async function main() {
    urls.forEach(item => {
        const contract = bridgeContracts[item.name]
        // contract.on("WithdrawDone", (local, recipient, value, event) => {
        //     connection.connect()
        //     connection.query("UPDATE bridge_logs SET isDone = ? WHERE ",[1])
        //     connection.end()
        // })

        contract.on("Deposit", async (chain, token, address, value, event) => {
            const blockLock = await getBlockLock(item.name)
            const blockNow = event.blockNumber
            if (blockNow > blockLock) {
                const _value = value.toString()
                const toContract = bridgeContracts[chain]
                if (toContract) {
                    toContract.withdraw(item.name, token, address, _value).then(_ => {
                        console.log("[跨链][成功] 链 " + item.name, "到链 " + chain, "代币 " + token, "地址 " + address, "数额 " + _value)
                    }).catch(error => {
                        console.log("[跨链][失败] 链 " + item.name, "到链 " + chain, "代币 " + token, "地址 " + address, "数额 " + _value)
                    })
                }
                setBlockLock(item.name, blockNow)
            }
        })
    })
}

main().then(_ => {

})