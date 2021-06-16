<template>
    <van-row style="margin-top: 5px">
        <van-col class="main" span="24" v-if="state.maskInstall">
            <div class="main">
                <van-nav-bar title="夸克跨链桥"></van-nav-bar>
                <van-form @submit="onDeposit">
                    <van-cell title="" clickable>
                        <template #label>
                            <div>主网:{{state.networkId}}</div>
                            <div>账号:{{state.account}}</div>
                            <div>余额:{{state.networkBalance}}</div>
                        </template>
                    </van-cell>

                    <van-cell-group title="跨链目标">
                        <van-cell v-for="(item,index) in state.targets" :key="index" :title="item.symbol" clickable
                                  @click="setTarget(item)" center>
                            <template #label>
                                <div>手续费: {{item.bridgeFee}}%</div>
                                <div>目标链: {{item.chain}}</div>
                                <div>{{item.remote}}</div>
                            </template>
                            <van-button size="mini" v-if="state.target === item" type="danger" icon="success">
                            </van-button>
                        </van-cell>
                    </van-cell-group>

                    <van-field v-model="state.bridgeNumber" type="number" label="数额"
                               :placeholder="state.tokenBalance">
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

    let provider = ref({})
    let signer = ref({})

    const state = reactive({
        bridge: "",
        loading: {balance: false, deposit: false},
        targets: [],
        network: {},
        networkBalance: "",
        account: "",
        accounts: [],
        bridgeNumber: "",
        maskInstall: true,
        tokenBalance: "",
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
            getBridge(state.networkId)
            getTargets(state.networkId)
            state.network = network
            const balance = await provider.getBalance(state.account)
            state.networkBalance = ethers.utils.formatEther(balance)
        } else {
            state.maskInstall = false
        }
    })

    const setTarget = async (target) => {
        if (state.target !== target) {
            state.bridgeNumber = ""
            state.target = target
            await getBalance(target)
        }
    }

    const getBalance = async (target) => {
        state.tokenBalance = ""
        state.loading.balance = true
        const abi = [
            "function balanceOf(address) view returns (uint256)"
        ]
        const contract = new ethers.Contract(target.local, abi, signer)
        const balance = await contract.balanceOf(state.account)
        state.tokenBalance = ethers.utils.formatUnits(balance, target.decimal)
        state.loading.balance = false
    }

    const getBridge = (networkId) => {
        const bridges = {
            "20181205": "0xcF70a42585473F160e7F5191dfe97fA15d5D8F3B",
            "3": "0xF891Ca04EA0276516A7823Ff787f923934dd56Aa",
            // "3777": "0x53AF942f88b73f835fB3eE1D48A846Da92029184",
            // "4777": "0x3736Ad67E30cCD77C6740Ea4a8e549c04cE439F2"
        }
        state.bridge = bridges[networkId]
    }

    const getTargets = (chain) => {
        const targets = {
            "20181205": [
                {
                    chain: "ropsten",
                    remote: "0xA52191450806181dEe22AAc5fc074C83c057AC43",
                    local: "0xE06235AAAA893336f1B38e10499Bef7Acb483EF9",
                    symbol: "cmm",
                    decimal: 0,
                    min: 10,
                    tokenFee: 5,
                    bridgeFee: 2
                },
                {
                    chain: "ropsten",
                    remote: "0xb6Cb36df5b88A3e5b547b04368Ebef30f04409a9",
                    local: "0x059d5e8be85b0E517F930e00Ddd52AC1Aaf61115",
                    symbol: "cnn",
                    decimal: 0,
                    min: 10,
                    tokenFee: 5,
                    bridgeFee: 2
                },
                {
                    chain: "ropsten",
                    remote: "0x0B8880087d8A76e03802Cd000e36126665e24363",
                    local: "0x5Cad8a1340E624Be90adfd787F6F3248DdF17321",
                    symbol: "PI",
                    decimal: 6,
                    min: 10,
                    tokenFee: 0,
                    bridgeFee: 2
                },
            ],
            "3": [
                {
                    chain: "qk",
                    remote: "0xE06235AAAA893336f1B38e10499Bef7Acb483EF9",
                    local: "0xA52191450806181dEe22AAc5fc074C83c057AC43",
                    symbol: "cmm",
                    min: 10,
                    decimal: 0,
                    tokenFee: 5,
                    bridgeFee: 2
                },
                {
                    chain: "qk",
                    remote: "0x059d5e8be85b0E517F930e00Ddd52AC1Aaf61115",
                    local: "0xb6Cb36df5b88A3e5b547b04368Ebef30f04409a9",
                    symbol: "cnn",
                    min: 10,
                    decimal: 0,
                    tokenFee: 5,
                    bridgeFee: 2
                },
                {
                    chain: "qk",
                    remote: "0x5Cad8a1340E624Be90adfd787F6F3248DdF17321",
                    local: "0x0B8880087d8A76e03802Cd000e36126665e24363",
                    symbol: "PI",
                    decimal: 6,
                    min: 10,
                    tokenFee: 0,
                    bridgeFee: 2
                },
            ]
        }
        state.targets = targets[chain]
        setTarget(state.targets[0])
    }

    const onDeposit = async () => {

        if (state.bridgeNumber < state.target.min) {
            Toast("最低跨链数额为" + state.target.min)
            return false
        }

        state.loading.deposit = true
        const token = new ethers.Contract(state.target.local, [
            "function approve(address spender, uint256 amount) external returns (bool)"
        ], signer)
        try {
            console.log('开始授权')
            const value = ethers.utils.parseUnits(state.bridgeNumber.toString(), state.target.decimal)
            const approve = await token.approve(state.bridge, value,{
                gasLimit:180000,
                gasPrice:ethers.utils.parseUnits("1","gwei")
            })
            console.log('等待打包')
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
            tokenFee = Math.ceil(state.bridgeNumber * (state.target.tokenFee / 100))
        }
        let final = state.bridgeNumber - tokenFee
        try {
            final = ethers.utils.parseUnits(final.toString(), state.target.decimal)
            console.log(state.target.chain, state.target.remote)
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

    // 跨链全部
    const bridgeMax = () => {
        state.bridgeNumber = state.tokenBalance
    }

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
