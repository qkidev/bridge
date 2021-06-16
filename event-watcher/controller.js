const {
    ethers
} = require('ethers')

const axios = require('axios').default

const fs = require("fs")

require('dotenv').config()

const mysql = require('mysql')

const baseUrl = "http:127.0.0.1/api/"

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


const setBlockLock = (chain, number) => {
    fs.writeFile(__dirname + "/lock/" + chain + ".lock", number.toString(), _ => {
    })
}


// 管理员密钥
const pk = {
    qk: process.env.PK_QK,
    ropsten: process.env.PK_ROPSTEN,
    heco: process.env.PK_HECO,
    eth: process.env.PK_ETH,
    okex: process.env.PK_OKEX,
    bsc: process.env.PK_BSC,
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

const getUrl = async () => {
    const response = await axios("http://127.0.0.1/api/chains")
    return response.data
}


// 支持链主网
const urls = getUrl()

// 全部链的跨链桥合约
let bridgeContracts = {}

// 设置提供者和钱包和跨链桥合约
urls.forEach(item => {
    const provider = new ethers.providers.JsonRpcProvider(item.url)
    const wallet = new ethers.Wallet(pk[item.name], provider)
    bridgeContracts[item.name] = new ethers.Contract(item.bridge, abiBridge, wallet)
})


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
                const fee = getNativeFee(item.name, toChain, value)
                const toContract = bridgeContracts[toChain]
                if (toContract) {
                    toContract.withdrawNative(item.name, recipient, value - fee).then(_ => {
                        logInsert({
                            depositHash: event.blockHash,
                            chainFrom: item.name,
                            chainTo: toChain,
                            block: blockNow,
                            toToken: "",
                            address: recipient,
                            value: value - fee,
                            status: 'deposit native success'
                        })
                        console.log("[跨链][成功] 链 " + item.name, "到链 " + toChain, "地址 " + recipient, "数额 " + value - fee)
                    })
                }

                setBlockLock(item.name, blockNow)
            }
        })

        contract.on("Deposit", async (toChain, toToken, address, value, event) => {
            // console.log(item.name,chain,token,address,value,event.blockNumber)
            const blockLock = await getBlockLock(item.name)
            const blockNow = event.blockNumber
            if (blockNow > blockLock) {
                const response = await axios.get(`http://127.0.0.1/api/token?fromChain=${item.name}&toChain=${toChain}&toToken=${toToken}`)
                const token = response.data

                let _value = ethers.utils.formatUnits(value, token.decimal)
                let fee = 0
                if (token.bridgeFee) {
                    fee = Math.ceil(value * (token.bridgeFee / 100))
                }
                _value = _value - fee
                const toContract = bridgeContracts[toChain]
                if (toContract) {
                    const final = ethers.utils.parseUnits(_value.toString(), token.decimal)
                    toContract.withdraw(item.name, toToken, address, final).then(_ => {
                        logInsert({
                            depositHash: event.blockHash,
                            chainFrom: item.name,
                            chainTo: toChain,
                            block: blockNow,
                            toToken: token,
                            address,
                            value: _value,
                            status: 'deposit success'
                        })
                        console.log("[跨链][成功] 链 " + item.name, "到链 " + toChain, "代币 " + toToken, "地址 " + address, "数额 " + _value)
                    }).catch(error => {
                        logInsert({
                            chainFrom: item.name,
                            chainTo: toChain,
                            block: blockNow,
                            depositHash: event.blockHash,
                            toToken: token,
                            address,
                            value: _value,
                            status: 'deposit error'
                        })
                        console.log("[跨链][失败] 链 " + item.name, "到链 " + toChain, "代币 " + toToken, "地址 " + address, "数额 " + _value)
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