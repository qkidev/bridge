const {
    ethers
} = require('ethers')

// 跨链桥合约部署参数 管理员地址 是为了保持部署地址nonce为初始值 则每个链部署的跨链桥合约地址一样

// 跨链桥地址
const bridge_address = {
    7545: "",
    8545: "",
    // QK: "0x58B05C1430d8b19A6541b0b4c4c2e8746F97F9af",
    // HECO: "0x58B05C1430d8b19A6541b0b4c4c2e8746F97F9af",
    // ETH: "0x58B05C1430d8b19A6541b0b4c4c2e8746F97F9af",
    // BSC: "0x58B05C1430d8b19A6541b0b4c4c2e8746F97F9af",
    // OKEX: "0x58B05C1430d8b19A6541b0b4c4c2e8746F97F9af",
}

// 管理员密钥
const pk = {
    7545: "",
    8545: ""
}



// 跨链桥ABI
const bridge_abi = [
    "event Sended(string,address,address,uint256)",
    "function recharge(address,address,uint256)",
    "function send(address,address,uint256)",
]

// 支持的跨链代币
const tokens = {
    7545: {
        // 外链代币 : 本链代币
        '': '' // CCT(8545)-CCT(7545)
    },
    8545: {
        // 外链代币 : 本链代币
        '': '' // CCT(7545)-CCT(8545)
    }
    // Local: {
    //     0x66006a5360d82B05bBf59bC394aC0093eAecf87C: 0x5DC98770a4BBA0A4cE9443E2198D5B6ebB7ADDFA, // QK-FEG
    // },
    // QK: {
    //     0x5DC98770a4BBA0A4cE9443E2198D5B6ebB7ADDFA: 0x66006a5360d82B05bBf59bC394aC0093eAecf87C, // Local-FEG
    // },
    // HECO: {
    //     0x5ba42785254fC7AC9282F2515892AB20BcD63aEA: 0x5ba42785254fC7AC9282F2515892AB20BcD63aEA
    // }
}

// 支持链主网
const urls = [{
        name: 7545,
        url: "http://127.0.0.1:7545"
    },
    {
        name: 8545,
        url: "http://127.0.0.1:8545"
    }
    // {
    //     name: "Local",
    //     url: "http://127.0.0.1:7545"
    // },
    // {
    //     name: "ETH",
    //     url: "https://mainnet.infura.io/v3/fddb1d98064647dd8bce19afb8b48059"
    // },
    // {
    //     name: "QK",
    //     url: "http://sg.node.quarkblockchain.org"
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
let bridge_contracts = {}

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
    bridge_contracts[item.name] = new ethers.Contract(bridge_address[item.name], bridge_abi, wallets[item.name])
})

// 收件人是应付款的地址
// 金额（以wei为单位）指定应发送的以太币数量
// 随机数可以是任何唯一数字，以防止重放攻击
// 合约地址 用于防止跨合约的重播攻击

async function main() {

    urls.forEach(item => {

        // 转账筛选器
        const filter = {
            // https://docs.ethers.io/v5/api/providers/provider/#Provider--events
            // https://docs.ethers.io/v5/api/providers/types/#providers-EventFilter
            // address: "",
            topics: [
                ethers.utils.id('Transfer(address,address,uint256)'),
                null,
                ethers.utils.hexZeroPad(bridge_address[item.name], 32) // 跨链桥合约地址
            ]
        }

        const contract = bridge_contracts[item.name]
        contract.on("Sended", (network, token, address, value) => {
            // console.log(network, token, address, value)
            const _value = value.toString()
            // 调用跨链桥转账
            const toContract = bridge_contracts[network]
            if (toContract) {
                toContract.send(tokens[network][token], address, _value).then(_ => {
                    console.log("[提币] 链 " + item.name, "到链 " + network, "代币 " + token, "地址 " + address, "数额 " + _value)
                }).catch(error => {
                    console.log("[提币失败] 链 " + item.name, "到链 " + network, "代币 " + token, "地址 " + address, "数额 " + _value)
                    console.log(error.message)
                })
            }
        })

        const provider = providers[item.name]
        provider.on(filter, log => {
            if (!log.removed) {
                const token = log.address
                // 获取转账数额
                const num = ethers.utils.defaultAbiCoder.decode(
                    ['uint256'],
                    log.data
                ).toString()

                // 取发送地址
                const from = ethers.utils.defaultAbiCoder.decode(
                    ['address'],
                    log.topics[1]
                )[0]
                if (num > 0) {
                    contract.recharge(token, from, num)
                    console.log("[充值] 链 " + item.name, "代币 " + token, "地址 " + from, "数额 " + num)
                }
            }
        })
    })
}

main()