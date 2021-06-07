const {
    ethers
} = require('ethers')

require('dotenv').config()

const mysql = require('mysql')

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
    // 7545: "",
    // 8545: "",
    qk: "",
    rop: "",
    // HECO: "",
    // ETH: "",
    // BSC: "",
    // OKEX: "",
}

// 管理员密钥
const pk = {
    qk: "",
    rop: "",
    // 7545: "",
    // 8545: ""
}

// 跨链桥ABI
const abiBridge = [
    "event Deposit(string chain, address remote, address recipient, uint256 value)",
    "event WithdrawDone(address local, address recipient, uint256 value)",
    "function withdraw(string memory chain, address remote, address recipient, uint256 value)",
]

// 支持链主网
const urls = [
    // {
    //     name: 7545,
    //     url: "http://127.0.0.1:7545"
    // },
    // {
    //     name: 8545,
    //     url: "http://127.0.0.1:8545"
    // },
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
        contract.on("WithdrawDone", (local, recipient, value) => {
            connection.connect()
            connection.query("UPDATE bridge_logs SET isDone = ? WHERE ",[1])
            connection.end()
        })

        contract.on("Deposit", (chain, token, address, value) => {
            const _value = value.toString()
            const toContract = bridgeContracts[chain]
            if (toContract) {
                toContract.withdraw(item.name, token, address, _value).then(_ => {
                    console.log("[提币] 链 " + item.name, "到链 " + chain, "代币 " + token, "地址 " + address, "数额 " + _value)
                }).catch(error => {
                    console.log("[提币失败] 链 " + item.name, "到链 " + chain, "代币 " + token, "地址 " + address, "数额 " + _value)
                    console.log(error.message)
                })
            }
        })
    })
}

main().then(_ => {

})