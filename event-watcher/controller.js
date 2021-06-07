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
const bridge_address = {
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
const bridge_abi = [
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
    // }
    // {
    //     name: "Local",
    //     url: "http://127.0.0.1:7545"
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
    wallets[item.name] = new ethers.Wallet(pk[item.name], provider);
})

// 设置跨链桥合约
urls.forEach(item => {
    bridgeContracts[item.name] = new ethers.Contract(bridge_address[item.name], bridge_abi, wallets[item.name])
})

// 收件人是应付款的地址
// 金额（以wei为单位）指定应发送的以太币数量
// 随机数可以是任何唯一数字，以防止重放攻击
// 合约地址 用于防止跨合约的重播攻击

async function main() {
    console.log(process.env.pk_heco)
    urls.forEach(item => {

        // 转账筛选器
        const filter = {
            // https://docs.ethers.io/v5/api/providers/provider/#Provider--events
            // https://docs.ethers.io/v5/api/providers/types/#providers-EventFilter
            // address: "0x3Cf054A2c78A536373c8Fc998AA643B4072A4181",
            topics: [
                ethers.utils.id('Transfer(address,address,uint256)'),
                null,
                ethers.utils.hexZeroPad(bridge_address[item.name], 32) // 跨链桥合约地址
            ]
        }

        const contract = bridgeContracts[item.name]
        contract.on("Deposit", (chain, token, address, value) => {
            const _value = value.toString()
            // 调用跨链桥转账
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