const getBalancePairs = () => {
    return new Promise((resolve, reject) => {
        fetch("/admin/getPairs").then(response => response.json())
            .then(myJson => resolve(myJson)).catch(error => reject(error))
    })
}

const getChain = (chainId) => {
    return new Promise((resolve, reject) => {
        fetch("/admin/getChain?chainId=" + chainId).then(response => response.json())
            .then(myJson => resolve(myJson)).catch(error => reject(error))
    })
}

const getPrivateKey = async (chainId) => {
    return new Promise((resolve, reject) => {
        fetch("/admin/getPrivateKey?chainId=" + chainId).then(response => response.json())
            .then(myJson => resolve(myJson)).catch(error => reject(error))
    })
}

const postBalance = (chainId, chainName, token, name, balance) => {
    return new Promise((resolve, reject) => {
        fetch("/admin/syncBalance?chainId=" + chainId + "&chainName=" + chainName + "&token=" + token + "&name=" + name + "&balance=" + balance).then(response => response.json())
            .then(myJson => resolve(myJson)).catch(error => reject(error))
    })
}


const syncBalance = async () => {
    const abi = [
        "function balanceOf(address) view returns (uint256)"
    ]
    const pairs = await getBalancePairs()
    for (const pair of pairs) {
        const chain = await getChain(pair['fromChain'])
        const privateKey = await getPrivateKey(pair['fromChain'])
        if (privateKey.length) {
            const provider = new ethers.providers.JsonRpcProvider(chain.url)
            if (pair.isNative && pair.isMain) {
                let balance = await provider.getBalance(chain.bridge)
                balance = ethers.utils.formatUnits(balance, pair.decimal)
                await postBalance(pair['fromChain'],chain['name'], pair['fromToken'], pair['name'], balance)
            } else {
                const contract = new ethers.Contract(pair['fromToken'], abi, provider)
                let balance = await contract.balanceOf(chain.bridge)
                balance = ethers.utils.formatUnits(balance, pair.decimal)
                await postBalance(pair['fromChain'],chain['name'], pair['fromToken'], pair['name'], balance)
            }
        }
    }
}


syncBalance().then()
