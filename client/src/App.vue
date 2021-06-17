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
            if (state.target.isMain) {
                state.tokenBalance = state.networkBalance
            } else {
                await getBalance(target)
            }
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
            // "20181205": "0xcF70a42585473F160e7F5191dfe97fA15d5D8F3B",
            // "3": "0xF891Ca04EA0276516A7823Ff787f923934dd56Aa",
            "3777": "0xfc40D4D82F13D55eA4774B081cb91fF18ad45AdB",
            "4777": "0xBed6437470c099AcC5d4674558d46C5a8bf1eeb4"
        }
        state.bridge = bridges[networkId]
    }

    const getTargets = (chain) => {
        const targets = {
            "3777": [
                {
                    chain: "4777",
                    remote: "0xC69Cdd72d03BFb18f287fd5BEe6B82D57F5e6d39",
                    local: "0x81A7099c33ec3767Af27BcDe23C084ca16A4Dbeb",
                    symbol: "CMM",
                    decimal: "0",
                    min: "10",
                    tokenFee: 5,
                    bridgeFee: 2,
                    isNative: false
                },
                {
                    chain: "4777",
                    remote: "0xcda14f12e483bbD64aF26A4f283350cb9e2a96A4",
                    local: "",
                    symbol: "H3777",
                    decimal: "18",
                    min: "1",
                    tokenFee: 0,
                    bridgeFee: 0,
                    isNative: true,
                    isMain: true
                },
            ],
            "4777": [
                {
                    chain: "3777",
                    remote: "0x81A7099c33ec3767Af27BcDe23C084ca16A4Dbeb",
                    local: "0xC69Cdd72d03BFb18f287fd5BEe6B82D57F5e6d39",
                    symbol: "CMM",
                    decimal: "0",
                    min: "10",
                    tokenFee: 5,
                    bridgeFee: 2
                },
                {
                    chain: "3777",
                    remote: "",
                    local: "0xcda14f12e483bbD64aF26A4f283350cb9e2a96A4",
                    symbol: "H3777",
                    decimal: "0",
                    min: "1",
                    tokenFee: 0,
                    bridgeFee: 0,
                    isNative: true,
                    isMain: false,
                },
            ],
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
        let tokenFee = 0
        if (state.target.tokenFee) {
            tokenFee = Math.ceil(state.bridgeNumber * (state.target.tokenFee / 100))
        }
        let final = state.bridgeNumber - tokenFee
        final = ethers.utils.parseUnits(final.toString(), state.target.decimal)

        console.log(state.target.isNative)
        if (state.target.isNative) {
            const abi = [
                "function depositNative(uint toChainId, uint256 value) payable"
            ]
            const bridge = new ethers.Contract(state.bridge, abi, signer)
            if (state.target.isMain) {
                try {
                    const tx = await bridge.depositNative(state.target.chain, final, {
                        gasLimit: 60000,
                        gasPrice: ethers.utils.parseUnits('1', 'gwei'),
                        value: final,
                    })
                    await tx.wait()
                } catch (e) {
                    console.log(e)
                }
            } else {
                try {
                    console.log(state.target.chain, final)
                    const tx = await bridge.depositNative(state.target.chain, final)
                    await tx.wait()
                } catch (e) {
                    console.log(e)
                }
            }

        } else {
            const token = new ethers.Contract(state.target.local, [
                "function approve(address spender, uint256 amount) external returns (bool)",
                "function allowance(address,address) view returns (uint256)"
            ], signer)
            const value = ethers.utils.parseUnits(state.bridgeNumber.toString(), state.target.decimal)
            const allowance = await token.allowance(state.account, state.bridge)
            if (allowance.toString() * 1 < value.toString() * 1) {
                try {
                    const approve = await token.approve(state.bridge, ethers.utils.parseUnits("10000000000000", state.target.decimal))
                    await approve.wait()
                } catch (e) {
                    state.loading.deposit = false
                    return false
                }
            }

            const abi = [
                "function deposit(uint chainId,address toToken,uint256 value)",
            ]
            const bridge = new ethers.Contract(state.bridge, abi, signer)
            try {
                // console.log(state.target.chain, state.target.remote, final.toString())
                const tx = await bridge.deposit(state.target.chain, state.target.remote, final)
                await tx.wait()
            } catch (e) {
                state.loading.deposit = false
                const message = e.data.message
                const arr = message.split(':')
                const messages = {
                    "revert token is can not use bridge": `${state.target.symbol} 暂时无法跨链`
                }
                Toast(messages[arr[1].trim()] || '跨链失败')

                return false
            }
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
