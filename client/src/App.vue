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
                        <van-cell v-for="(item,index) in state.targets" :key="index" :title="item.title" clickable
                                  @click="setTarget(item)" center>
                            <template #label>
                                <div>手续费: {{item.bridgeFee}}%</div>
                                <div>目标链: {{item.toChain}}</div>
                                <div>{{item.toToken}}</div>
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

    const baseUri = "http://148.70.33.227:9580/" // http://127.0.0.1:8000/

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
            if (state.target.isMain && state.target.isNative) {
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
        const contract = new ethers.Contract(target.fromToken, abi, signer)
        const balance = await contract.balanceOf(state.account)
        state.tokenBalance = ethers.utils.formatUnits(balance, target.decimal)
        state.loading.balance = false
    }

    const getBridge = (networkId) => {
        fetch(baseUri + 'api/bridges')
            .then(function (response) {
                return response.json();
            })
            .then(function (bridges) {
                state.bridge = bridges[networkId]
            });
    }

    const getTargets = (chain) => {
        fetch(baseUri +'api/pairs')
            .then(function (response) {
                return response.json();
            })
            .then(function (targets) {
                state.targets = targets[chain]
                setTarget(state.targets[0])
            });
    }

    const onDeposit = async () => {
        // console.log(state.target)

        if (state.bridgeNumber < state.target.minValue) {
            Toast("最低跨链数额为" + state.target.minValue)
            return false
        }
        state.loading.deposit = true
        let tokenFee = 0
        if (state.target.tokenFee) {
            tokenFee = Math.ceil(state.bridgeNumber * (state.target.tokenFee / 100))
        }
        let final = state.bridgeNumber - tokenFee
        final = ethers.utils.parseUnits(final.toString(), state.target.decimal)
        // console.log(final.toString())
        // console.log(state.target.isNative)
        if (state.target.isNative) {
            const abi = [
                "function depositNative(uint toChainId, bool isMain, uint256 value) payable"
            ]
            const bridge = new ethers.Contract(state.bridge, abi, signer)
            // console.log(state.target.toChain, !!state.target.isMain, final.toString())
            if (state.target.isMain) {
                try {
                    const tx = await bridge.depositNative(state.target.toChain, state.target.isMain, final, {
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
                    const token = new ethers.Contract(state.target.fromToken, [
                        "function approve(address spender, uint256 amount) external returns (bool)",
                        "function allowance(address,address) view returns (uint256)"
                    ], signer)
                    const allowance = await token.allowance(state.account, state.bridge)
                    if (allowance.toString() * 1 < final) {
                        try {
                            const approve = await token.approve(state.bridge, ethers.utils.parseUnits("10000000000000", state.target.decimal))
                            await approve.wait()
                        } catch (e) {
                            state.loading.deposit = false
                            return false
                        }
                    }
                    const tx = await bridge.depositNative(state.target.toChain, !!state.target.isMain, final)
                    await tx.wait()
                } catch (e) {
                    console.log(e)
                }
            }

        } else {
            // console.log(signer)
            //console.log(state.target.fromToken)
            const token = new ethers.Contract(state.target.fromToken, [
                "function approve(address spender, uint256 amount) external returns (bool)",
                "function allowance(address,address) view returns (uint256)"
            ], signer)
            //console.log(state.bridge)
            const allowance = await token.allowance(state.account, state.bridge)
            //console.log(allowance.toString())
            if (allowance.toString() * 1 < final.toString() * 1) {
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
                // console.log(state.bridge,state.target.toChain, state.target.toToken, final.toString())
                const tx = await bridge.deposit(state.target.toChain, state.target.toToken, final)
                await tx.wait()
            } catch (e) {
                console.log(e)
                Toast( '跨链失败')
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
        font-family: "ubuntu", "Microsoft YaHei UI Light", sans-serif !important;
        background: #f7f8fa;
    }

    .main {
        max-width: 450px;
        margin: 0 auto;
    }
</style>
