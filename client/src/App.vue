<template>
    <van-row style="margin-top: 5px">
        <van-col class="main" span="24" v-if="state.installed_mask">
            <div class="main">
                <van-nav-bar title="夸克跨链桥"></van-nav-bar>
                <van-form @submit="onDeposit">
                    <van-cell title="" clickable>
                        <template #label>
                            <div>主网:{{state.networkId}}</div>
                            <div>账号:{{state.account}}</div>
                            <div>余额:{{state.network_balance}}</div>
                        </template>
                    </van-cell>

                    <van-cell-group title="兑换目标">
                        <van-cell v-for="(item,index) in state.targets" :key="index" :title="item.symbol" clickable
                                  @click="set_target(item)" center>
                            <template #label>
                                <div>手续费: {{item.bridgeFee}}%</div>
                                <div>目标链: {{item.chain}}</div>
                                <div>{{item.remote}}</div>
                            </template>
                            <van-button size="mini" v-if="state.target === item" type="danger" icon="success">
                            </van-button>
                        </van-cell>
                    </van-cell-group>

                    <van-field v-model="state.bridge_number" type="number" label="数额"
                               :placeholder="state.token_balance">
                        <template #button>
                            <van-button size="small" type="primary" @click="bridgeMax" :loading="state.loading.balance">
                                全部
                            </van-button>
                        </template>
                    </van-field>
                    <div style="margin: 16px;">
                        <van-button block type="primary" native-type="submit" :loading="state.loading.deposit">
                            跨链
                        </van-button>
                    </div>
                </van-form>
            </div>
        </van-col>
        <van-col :span="24" v-else>
            <van-button block type="danger">请安装小狐狸钱包</van-button>
        </van-col>
    </van-row>


</template>

<!--https://chainid.network/-->
<script setup>
    import {ethers} from 'ethers'
    import {Toast} from 'vant'

    import {reactive, onMounted, ref} from "vue";
    import {f} from "../dist/assets/vendor.7349c458";

    let provider = ref({})
    let signer = ref({})

    const state = reactive({
        bridge: "",
        loading: {balance: false},
        show: {deposit: {target: false}},
        deposit: {target: {chain: '', symbol: ''}, balance: 0},
        targets: [],
        bridges: {
            qk: "0x74bC68dB910f79392fFBb867747dC1b616e76c83"
        },
        to_networks: ['quarkblockchain', 'ropsten'],
        network: {},
        network_balance: "",
        cct_balance: "",
        cct_recharge_number: "",
        showRecharge: false,
        cct_address: "0x46677d627fA43fCFE7D5c382387bE8dF96c6ffd5",
        bridge_address: "0x74bC68dB910f79392fFBb867747dC1b616e76c83",
        bridge_balance: "",
        maxFrom: 0,
        fromNetwork: '',
        toNetwork: '',
        account: "",
        accounts: [],
        bridge_number: "",
        showFromNetworkPicker: false,
        showToNetworkPicker: false,
        installed_mask: true,
        exchange_loading: false,
        recharge_loading: false,
        token_balance: "",
        networkId: ""
    });

    onMounted(async () => {
        if (typeof window.ethereum !== 'undefined') {
            window.ethereum.on('chainChanged', _ => location.reload())
            state.accounts = await ethereum.request({method: 'eth_requestAccounts'})
            state.account = state.accounts[0]
            provider = new ethers.providers.Web3Provider(web3.currentProvider)
            signer = provider.getSigner()
            const network = await provider.getNetwork()
            state.networkId = await ethereum.request({method: 'net_version'})
            if (network.chainId === 20181205) network.name = "quarkblockchain"
            get_bridge(state.networkId)
            get_targets(state.networkId)
            state.network = network
            const balance = await provider.getBalance(state.account)
            state.network_balance = ethers.utils.formatEther(balance)
        } else {
            state.installed_mask = false
        }
    })

    const set_target = async (target) => {
        if (state.target !== target) {
            state.target = target
            await getBalance(target)
        }
    }

    const getBalance = async (target) => {
        state.token_balance = ""
        state.loading.balance = true
        const abi = [
            "function balanceOf(address) view returns (uint256)"
        ]
        const contract = new ethers.Contract(target.local, abi, signer)
        const balance = await contract.balanceOf(state.account)
        state.token_balance = ethers.utils.formatUnits(balance, target.decimal)
        state.loading.balance = false
    }

    const get_bridge = (networkId) => {
        const bridges = {
            "3777": "0x53AF942f88b73f835fB3eE1D48A846Da92029184",
            "4777": "0x3736Ad67E30cCD77C6740Ea4a8e549c04cE439F2"
        }
        state.bridge = bridges[networkId]
    }

    const get_targets = (chain) => {
        const targets = {
            "3777": [
                {
                    chain: "8545",
                    remote: "0xFf99daF379794921b8100b3D262A9668Fb3c068E",
                    local: "0x50AdE7689Ddd197E29A68b516d30f77D83006C45",
                    symbol: "cmm",
                    decimal: 8,
                    min: 10,
                    tokenFee: 5,// 代币转账费
                    bridgeFee: 2 // 跨链手续费
                },
                {
                    chain: "8545",
                    remote: "0x8B52c9e3d66034b413FF90087FED530e092c7920",
                    local: "0xCA371E99AE6FbB2883B4A5d8c6604f0E747796eC",
                    symbol: "cnn",
                    decimal: 0,
                    min: 10,
                    tokenFee: 5,// 代币转账费
                    bridgeFee: 2 // 跨链手续费
                },
            ],
            "4777": [
                {
                    chain: "7545",
                    remote: "0x50AdE7689Ddd197E29A68b516d30f77D83006C45",
                    local: "0xFf99daF379794921b8100b3D262A9668Fb3c068E",
                    symbol: "cmm",
                    min: 10,
                    decimal: 8,
                    tokenFee: 5,// 代币转账费
                    bridgeFee: 2 // 跨链手续费
                },
                {
                    chain: "7545",
                    remote: "0xCA371E99AE6FbB2883B4A5d8c6604f0E747796eC",
                    local: "0x8B52c9e3d66034b413FF90087FED530e092c7920",
                    symbol: "cmm",
                    min: 10,
                    decimal: 0,
                    tokenFee: 5,// 代币转账费
                    bridgeFee: 2 // 跨链手续费
                }
            ]
        }
        state.targets = targets[chain]
        set_target(state.targets[0])
    }

    const onDeposit = async () => {

        if (state.bridge_number < state.target.min) {
            Toast("最低跨链数额为" + state.target.min)
            return false
        }

        state.loading.deposit = true
        const token = new ethers.Contract(state.target.local, [
            "function approve(address spender, uint256 amount) external returns (bool)"
        ], signer)
        try {
            const value = ethers.utils.parseUnits(state.bridge_number.toString(), state.target.decimal)
            const approve = await token.approve(state.bridge, value)
            await approve.wait()
        } catch (e) {
            Toast("授权失败")
            state.loading.deposit = false
            return false
        }
        const abi = [
            "function deposit(string memory chain,address remote,uint256 value)"
        ]
        const bridge = new ethers.Contract(state.bridge, abi, signer)

        let tokenFee = 0
        if (state.target.tokenFee) {
            tokenFee = Math.ceil(state.bridge_number * (state.target.tokenFee / 100))
        }
        let final = state.bridge_number - tokenFee
        try {
            final = ethers.utils.parseUnits(final.toString(), state.target.decimal)
            const tx = await bridge.deposit(state.target.chain, state.target.remote, final)
            await tx.wait()
        } catch (e) {
            const message = e.data.message
            const arr = message.split(':')
            const messages = {
                "revert token is can not use bridge": `${state.target.symbol} 暂时无法跨链`
            }
            Toast(messages[arr[1].trim()] || '跨链失败')
            state.loading.deposit = false
            return false
        }

        state.loading.deposit = false
        Toast("跨链完成")
    }

    const setDepositTarget = async (target) => {
        state.deposit.target = target
        state.show.deposit.target = false
    }

    const getChainTargets = async (chain) => {
        state.targets = []
        const abi = [
            "function tokenLength() view returns (uint256)",
            "function tokens(uint256) view)",
        ]
        const bridge = new ethers.Contract(state.bridges[chain], abi)
        const length = await bridge.tokenLength()
        for (let i = 0; i < length; i++) {
            const token = await bridge.tokens(i)
            state.targets.unshift(token)
        }
        state.deposit.target = state.targets[0]
    }

    // 兑换全部
    const bridgeMax = () => {
        state.bridge_number = state.token_balance
    }

    // 充值全部
    const rechargeMax = () => {
        state.cct_recharge_number = state.cct_balance
    }

    const onShowRecharge = async () => {
        const abi = [
            "function balanceOf(address) view returns (uint256)"
        ]
        const contract = new ethers.Contract(state.cct_address, abi, signer)
        const balance = await contract.balanceOf(state.account)
        state.cct_balance = balance.toString()
        state.showRecharge = true
    }

    const onRecharge = async () => {
        const abi = [
            "function transfer(address, uint256) public returns(bool success)"
        ]

        if (state.cct_recharge_number < 1) {
            Toast("请输入要充值的数额")
            return false
        }

        state.recharge_loading = true

        const contract = new ethers.Contract(state.cct_address, abi, signer)
        const tx = await contract.transfer(state.bridge_address, state.cct_recharge_number)
        await tx.wait()
        await getBridgeBalance()
        state.recharge_loading = false
        state.cct_recharge_number = ""
        state.showRecharge = false
    }

    const getBridgeBalance = async () => {
        const abi = [
            "function balances(address) view returns (uint256)"
        ]
        const contract = new ethers.Contract(state.bridge_address, abi, signer)
        const balance = await contract.balances(state.account)
        state.bridge_balance = balance.toString()
    }

    const onExchange = async () => {
        const abi = [
            "function exchange(string,uint256)"
        ]

        if (!state.toNetwork) {
            Toast("请选择转入链")
            return false
        }

        if (state.bridge_number < 1) {
            Toast("输入兑换数额")
            return false
        }
        if (state.bridge_number > state.bridge_balance) {
            Toast("跨链桥余额不足")
            return false
        }
        state.exchange_loading = true

        const contract = new ethers.Contract(state.bridge_address, abi, signer)
        const networks = {
            ropsten: "ROP",
            quarkblockchain: "QK"
        }
        const tx = await contract.exchange(networks[state.toNetwork], state.bridge_number)
        await tx.wait()
        state.exchange_loading = false
        state.bridge_number = ""
        Toast("兑换成功")
    };

    const changeChain = (chainId, chainName, name, symbol, decimals, rpcUrls, blockExplorerUrls) => {
        // https://chainid.network/
        window.ethereum.request({
            method: 'wallet_addEthereumChain',
            params: [{
                chainId, chainName,
                nativeCurrency: {
                    name,
                    symbol,
                    decimals
                },
                //rpcUrls: ['https://api.avax.network/ext/bc/C/rpc']
                rpcUrls,
                // blockExplorerUrls: ['https://cchain.explorer.avax.network/']
                blockExplorerUrls
            }]
        })
    }

    const onConfirmToNetwork = (value) => {
        state.toNetwork = value;
        state.showToNetworkPicker = false;
    };

    const onConfirmFromNetwork = (value) => {
        state.fromNetwork = value;
        state.showFromNetworkPicker = false;
    };

    // This starter template is using Vue 3 experimental <script setup> SFCs
    // Check out https://github.com/vuejs/rfcs/blob/script-setup-2/active-rfcs/0000-script-setup.md
</script>

<style>
    body {
        background: #f7f8fa;
    }

    .main {
        max-width: 450px;
        margin: 0 auto;
    }
</style>
