<template>
  <div class="bridge_container">
    <div class="flex_h_between_center">
      <img src="../assets/bridge/logo.png" alt="" class="logo_img">
      <div class="addr_wrap flex_h_center_center" v-if="address != ''">
        <span class="fStyle22_ffffff">{{subAddress(address)}}</span>
        <img src="../assets/bridge/wallet.png" alt="" class="wallet_icon">
      </div>
    </div>
    <div class="flex_h_between_center mt_70" v-if="token != null">
      <span class="fStyle28_4570B3">资产</span>
      <span class="fStyle24_ffffff upper-case">{{token.name}}:{{tokenBalance}}</span>
    </div>
    <div class="section flex_h_between_center mt_20" @click="showAssetModel = true">
      <div class="flex1 flex_h_center_start" v-if="token != null">
        <img :src="requireIcon(token.icon,token.fromChain)" alt="" class="coin_logo" />
        <span class="fStyle30_ffffff upper-case">{{token.name}}</span>
      </div>
      <div class="flex1 flex_h_center_start" v-else>
        <span class="fStyle30_ffffff">请选择资产</span>
      </div>
      <div>
      <i class="iconfont icon-arrow-right-bold" style="color:#fff"></i></div>
    </div>
    <div class="alignLeft mt_70"><span class="fStyle28_4570B3">跨链桥</span></div>
    <div class="section flex_h_center_center mt_20">
      <img src="../assets/bridge/dot.png" alt="" class="dot_icon"/>
      <div class="flex1 flex_v_start_center" v-if="currNetwork != null">
        <div class="flex_h_between_center border" style="width: 100%">
          <span class="fStyle26_5E81BB">从</span>
          <div class="flex1 flex_h_center_start"><img :src="currNetwork.fromChainData.icon" alt="" class="coin_logo_min" />
          <span class="fStyle30_ffffff upper-case">{{currNetwork.fromChainData.name}}</span></div>
        </div>
        <div class="flex_h_between_center" style="width: 100%" @click="showNetworkModel = true">
          <span class="fStyle26_5E81BB">从</span>
          <div class="flex1 flex_h_center_start"><img :src="currNetwork.toChainData.icon" alt="" class="coin_logo_min" />
          <span class="fStyle30_ffffff upper-case">{{currNetwork.toChainData.name}}</span></div>
          <i class="iconfont icon-arrow-down-filling arrow_down" ></i>
        </div>
      </div>
      <img src="../assets/bridge/exchange.png" alt="" class="exchange_icon" @click="changeChain(currNetwork.toChainData.chainId)"/>
    </div>
    <div class="alignLeft mt_70"><span class="fStyle28_4570B3">数量</span></div>
    <div class="bridge_input_wrap mt_20">
      <input type="text" class="fStyle30_ffffff" v-model="amount" />
      <div class="all" @click="amount = tokenBalance"><span class="fStyle30_587AFF">全部</span></div>
    </div>
    <div class="flex_h_start_start bridge_bg mt_30">
      <div class="flex_h_center_start amount_wrap" v-if="amount != '' && amount != 0 && currNetwork != null && token != null">
        <span class="fStyle24_6391DE">你会收到≈</span>
        <img :src="token.icon" alt="" class="coin_logo_miner" />
        <span class="fStyle24_6391DE upper-case">{{receiveAmount}} {{token.name}}</span>
        <div class="link_wrap"><span class="fStyle20_ACB6FF upper-case">{{currNetwork.toChainData.name}}</span></div>
      </div>
      <div class="alignLeft instruction">
          <li>跨链资产都会转入另外主网您当前地址({{address}})内，如果没有看到，请添加合约地址。</li>
          <li>如果要切换跨链兑换的主网，请切换钱包的主网。</li>
          <li>跨链时间在5分钟-10分钟。</li>
          <li>更多资产类似请自行申请跨链。</li>
          <li>本跨链桥为分布式跨链，即将开放跨链节点申请，敬请关注。</li>
      </div>
    </div>
    <div class="bridge_submit_btn flex_h_center_center" @click="verifySubmit">
      <span class="fStyle32_ffffff">确定兑换</span>
    </div>
    <!-- 资产类型弹框 -->
    <van-popup v-model="showAssetModel" round>
      <div class="popup_container">
        <div class="flex_h_between_center popup_head">
          <span class="fStyle32_08162C_w5">选择资产</span>
          <i class="iconfont icon-delete-filling" @click="showAssetModel= false"></i>
        </div>
        <div class="list_wrap">
          <div class="flex_h_center_start popup_item" v-for="(item, index) in tokenList" :key="index" @click="setToken(item)">
              <img :src="requireIcon(item.icon,item.fromChain)" alt="" class="coin_logo_big" />
            <div class="flex_v_center_start flex1">
              <span class="fStyle28_333333_w5 upper-case">{{item.name}}</span>
              <span class="fStyle22_999999">{{item.title}}</span>
              <span class="fStyle22_999999">{{shortAddress(item.fromToken)}}</span>
            </div>
            <i class="iconfont icon-select-bold checked" v-if="item.fromToken == token.fromToken"></i>
          </div>
        </div>
      </div>
    </van-popup>
    <!-- 网络类型弹框 -->
    <van-popup v-model="showNetworkModel" round>
      <div class="popup_container">
        <div class="flex_h_between_center popup_head">
          <span class="fStyle32_08162C_w5">选择主链</span>
          <i class="iconfont icon-delete-filling" @click="showNetworkModel= false"></i>
        </div>
        <div class="list_wrap">
          <div class="flex_h_center_start popup_item" v-for="(item, index) in netWorkList" :key="index" @click="setChain(item)">
            <img :src="item.toChainData.icon" alt="" class="coin_logo_big" />
            <div class="flex_v_center_start flex1">
              <span class="fStyle28_333333_w5 upper-case">{{item.toChainData.name}}</span>
            </div>
            <i class="iconfont icon-select-bold checked" v-if="item.toChainData.chainId == currNetwork.toChainData.chainId"></i>
          </div>
        </div>
      </div>
    </van-popup>
    <van-popup v-model="showOrderModel" round>
      <div class="popup_container" v-if="token != null && currNetwork != null">
        <div class="flex_h_between_center popup_head">
          <span class="fStyle32_08162C_w5">确认兑换</span>
          <i class="iconfont icon-delete-filling" @click="showOrderModel= false"></i>
        </div>
        <div class="list_wrap flex_v_center_center mt_30">
          <span class="fStyle48_333333_w6">{{amount}}</span>
          <div class="flex_h_center_center"><span class="fStyle26_333333 upper-case">{{token.name}}</span><div class="link_wrap blue_bg"><span class="fStyle20_FFFFFF upper-case">{{currNetwork.fromChainData.name}}</span></div></div>
          <img src="../assets/bridge/arrow_down.png" alt="" class="arrow_down_icon" />
          <span class="fStyle48_333333_w6">{{receiveAmount}}</span>
          <div class="flex_h_center_center"><span class="fStyle26_333333 upper-case">{{token.name}}</span><div class="link_wrap grey_bg"><span class="fStyle20_FFFFFF upper-case">{{currNetwork.toChainData.name}}</span></div></div>
          <div class="popup_submit mt_70 flex_h_center_center" @click="submit"><span class="fStyle32_ffffff">确定兑换</span></div>
        </div>
      </div>
    </van-popup>
     <van-popup v-model="loadingModel" :close-on-click-overlay="false">
       <div class="loading-wrap flex_v_center_center">
         <img src="../assets/loading.gif" alt="" class="logo_icon">
         <span class="fStyle26_333333">区块确认中，请等待...</span>
       </div>
     </van-popup>
  </div>
</template>

<script>
import {networkApi} from '../utils/request/api';
import { ethers } from "ethers"
import {Toast} from 'vant'
import {asyncUtils, Decimal} from '../utils/utils'
import {NORMAL_ABI, BRIDGE_ABI} from '../utils/const'
import Big from 'big.js'
export default {
  data () {
    return {
      provider: null,
      signer: null,
      chainId: 0,
      address: '',
      showAssetModel: false,
      showNetworkModel: false,
      showOrderModel: false,
      tokenList: [],
      token: null,
      netWorkList: [],
      currNetwork: null, // 跨链桥兑换到的主网地址
      tokenBalance: 0.00,
      tokenDecimal: 0,
      tokenContract: null,
      amount: '',
      min_gasprice: '10',
      bridgeContract: null,
      bridgeAddress: '',
      loadingModel: false,
      // loading: false,
      gweis:{}
    }
  },
  mixins: [asyncUtils, Decimal],
  watch: {
    token () {
      this.getTokenBalance()
    },
    chainId () {
      this.getTokenList()
    }
  },
  computed: {
    receiveAmount () {
      if(this.amount == '' || this.amount == 0) {
        return 0.00
      }
      let fee = Decimal.add(this.currNetwork.bridgeFee, this.currNetwork.tokenFee)
      fee = Decimal.sub(100, fee)
      fee = Decimal.div(fee, 100)
      return Decimal.mul(this.amount, fee)
    }
  },
  async mounted () {
     if (typeof ethereum == "undefined") {
      Toast('请安装metamask插件、或者使用owncoin打开')
    } else {
      window.ethereum.enable();
      await this.init();
      // 监听链的改变
      window.ethereum.on('chainChanged', async () => {
        await this.init()
      });
      window.ethereum.on('accountsChanged', async (accounts) => {
        this.address = accounts[0]
        this.getTokenBalance()
      });
    }
  },
  methods: {
      requireIcon(url, chainId) {
          if (!url) {
              url = require(`../assets/icon/${chainId}.png`)
          }
          return url
      },
    shortAddress(address){
        if (address === "0x0000000000000000000000000000000000000000")
          return "主网币"
      return address.slice(0,14)+"****"+address.slice(-14)
    },
    async init() {
      let customHttpProvider = new ethers.providers.Web3Provider(
        window.ethereum
      );
      this.provider = customHttpProvider;
      this.signer = customHttpProvider.getSigner();
      let [err, address] = await this.to(this.signer.getAddress())
      if(err == null) {
        this.address = address
      }
      let network = await customHttpProvider.getNetwork();
      this.chainId = network.chainId;
    },

    async getTokenBalance() {
      if(this.token.fromToken == ethers.constants.AddressZero) {
        this.tokenDecimal = this.currNetwork.decimal
        let [err, balance] = await this.to(this.provider.getBalance(this.address))
        this.doResponse(err, balance, 'tokenBalance', this.tokenDecimal)
      } else {
        let tokenAddress = this.token.fromToken
        let contract = new ethers.Contract(tokenAddress, NORMAL_ABI, this.signer)
        this.tokenContract = contract
        let [, decimal] = await this.to(contract.decimals())
        this.tokenDecimal = decimal
        let [err, balance] = await this.to(contract.balanceOf(this.address))
        this.doResponse(err, balance, 'tokenBalance', this.tokenDecimal)
      }
    },
    async getBalance() {
      if(this.token.fromToken == ethers.constants.AddressZero){
        let [err, balance] = await this.to(this.provider.getBalance(this.address))
        this.doResponse(err, balance, 'tokenBalance', this.tokenDecimal)
      } else {
        let [err, balance] = await this.to(this.tokenContract.balanceOf(this.address))
        this.doResponse(err, balance, 'tokenBalance', this.tokenDecimal)
      }
    },
    async getTokenList () {
      let json = await networkApi({chainId: this.chainId})
      if(json.code === 0) {

        // 保存gweis
        this.gweis = json.gweis

        let tempList = []
        for (let i in (json.data || [])){
          let first = (json.data || [])[i].length > 0 ? (json.data || [])[i][0] : []
          let obj = {
            name: i,
            icon: first ? first.icon : '',
            fromToken: first ? first.fromToken : '',
            toToken: first?first.toToken:'',
            fromChain: first?first.fromChain:'',
            title: first ? first.title : '',
            networks: (json.data || [])[i]
          }
          tempList.push(obj)
        }
        this.tokenList = tempList
        this.setToken((this.tokenList).length > 0 ? (this.tokenList || [])[0] : null)
      }
    },
    setToken(tokenObj) {
      if(tokenObj == null){
        return
      }
      this.netWorkList = tokenObj.networks || []
      this.currNetwork = (tokenObj.networks || []).length > 0 ? (tokenObj.networks || [])[0] : null
      this.bridgeAddress = this.currNetwork.fromChainData.bridge
      this.token = tokenObj
      this.showAssetModel = false
    },
    setChain(chainObj) {
      this.currNetwork = chainObj
      this.bridgeAddress = chainObj.fromChainData.bridge
      this.showNetworkModel = false
    },
    async set_target(target) {
      if (this.target !== target) {
          this.target = target
      }
    },
    verifySubmit() {
      if(this.token == null) {
        Toast('请选择需要兑换的资产类型');
        return
      }
      if(this.currNetwork == null) {
        Toast('请选择需要跨链兑换的主网');
        return
      }
      this.showOrderModel = true
    },
    // 查询授权
    async allowanceContract(amount, callback) {
      let [err2, allowce] = await this.to(
        this.tokenContract.allowance(this.address, this.bridgeAddress)
      );
      if (err2 == null) {
        let hex = ethers.utils.hexValue(allowce);
        let Value = Big(this.hex2int(hex));
        if (Decimal.sub(Value, amount) < 0) {
          // if(this.coinAddress.toUpperCase() === '0xdf0e293cc3c7ba051763ff6b026da0853d446e38'.toUpperCase() && Decimal.sub(Value, amount) != 0){
          //   await this.approveContract(false, 0);
          // }
          await this.approveContract(callback);
        } else {
          callback();
        }
      } else {
        await this.approveContract(callback);
      }
    },
    // 授权
    async approveContract(callback, amount) {
      let approveAmount = amount != undefined ? amount :
        "1000000000000000000000000000000000000000000000000000000000000000000000000000";
      const gasLimit = await this.getEstimateGas(() =>
        this.tokenContract.estimateGas.approve(
          this.bridgeAddress,
          approveAmount,
          {
            gasPrice: ethers.utils.parseUnits(this.gweis[this.chainId]+"", "gwei"),
          }
        )
      );
      if(gasLimit == 0) {
        this.loadingModel = false
        return 0;
      }
      let [err, res] = await this.to(
        this.tokenContract.approve(this.bridgeAddress, approveAmount, {
          gasLimit,
          gasPrice: ethers.utils.parseUnits(this.gweis[this.chainId]+"", "gwei"),
        })
      );
      if (this.doResponse(err, res)) {
        Toast("权限申请中...");
        this.queryTransation(this.provider, res.hash, callback);
      } else {
        this.loadingModel = false
      }
    },
    async submit() {
      let amount = this.calcAmount();
      console.log(amount)
      let minAmount = this.calcWei(this.currNetwork.minValue);
      let maxAmount = this.calcWei(this.currNetwork.maxValue);
      if(Big(amount).lt(Big(minAmount))) {
        Toast("最低跨链数额为" + this.currNetwork.minValue)
        return false
      }
      if(!(Big(maxAmount).eq(Big(0))) && Big(amount).gt(Big(10)) ) {
        Toast('最高跨链数额为'+this.currNetwork.maxValue)
        return false
      }
      // let bridgeAddress = this.currNetwork.fromChainData.bridge
      let constract = new ethers.Contract(this.bridgeAddress, BRIDGE_ABI, this.signer)
      this.bridgeContract = constract
      this.loadingModel = true
      if(this.currNetwork.isNative == 1) {
        if(this.currNetwork.isMain == 1) {
          this.exchangeMain()
        } else {
          this.allowanceContract(amount, this.exchangeMain);
        }
      } else {
        this.allowanceContract(amount, this.exchange);
      }
    },
    calcAmount() {
      let amount;
      if(Number(this.currNetwork.tokenFee) == 0) {
        amount = this.calcWei(this.amount)
      } else {
        let tempAmount =Decimal.mul(Decimal.div(Decimal.sub('100', this.currNetwork.tokenFee), '100'), this.amount).toFixed(this.tokenDecimal)
        amount =this.calcWei(tempAmount)
      }
      return (amount).split('.')[0];
    },
    calcWei(amount) {
      return Decimal.mul(amount, ethers.BigNumber.from(10).pow(Number(this.tokenDecimal))).toFixed()
    },
    async exchangeMain() {
      let amount = this.calcAmount();
      let value = this.currNetwork.isMain ? amount : 0
      const gasLimit = await this.getEstimateGas(() =>
        this.bridgeContract.estimateGas.depositNative(
          this.currNetwork.toChain, this.currNetwork.isMain, amount,
          {
            value: value,
            gasPrice: ethers.utils.parseUnits(this.gweis[this.chainId]+"", "gwei"),
          }
        )
      );
      if(gasLimit == 0) {
        this.loadingModel = false
        return 0;
      }
      let [err, res] = await this.to(this.bridgeContract.depositNative(this.currNetwork.toChain, this.currNetwork.isMain, amount, {
            gasLimit,
            value: value,
            gasPrice: ethers.utils.parseUnits(this.gweis[this.chainId]+"", "gwei"),
          }))
          if(this.doResponse(err, res)) {
            this.queryTransation(this.provider, res.hash, async () => {
              this.loadingModel = false
              this.showOrderModel = false
              this.getBalance();
              this.amount = ''
            });
          } else {
            this.loadingModel = false
          }
    },
    async exchange() {
      let amount = this.calcAmount()
      console.log(amount)
      const gasLimit = await this.getEstimateGas(() =>
        this.bridgeContract.estimateGas.deposit(
          this.currNetwork.toChain, this.currNetwork.toToken, amount,
          {
            gasPrice: ethers.utils.parseUnits(this.gweis[this.chainId]+"", "gwei"),
          }
        )
      );
      if(gasLimit == 0) {
        this.loadingModel = false
        return 0;
      }
      let [err, res] = await this.to(this.bridgeContract.deposit(this.currNetwork.toChain, this.currNetwork.toToken,  amount, {
        gasLimit,
        gasPrice: ethers.utils.parseUnits(this.gweis[this.chainId]+"", "gwei"),
      }))
      if(this.doResponse(err, res)) {
        this.queryTransation(this.provider, res.hash, async () => {
          this.loadingModel = false
          this.showOrderModel = false
          this.getBalance();
          this.amount = ''
        });
      } else {
        this.loadingModel = false
      }
    },
    // 切换网络
    changeChain(chainId) {
      const origin = (window.location.origin || '').split('#')[0]
      const url = origin + '/#/home'
      this.openNativeReadPaper(url, chainId);
    },
    openNativeReadPaper(url, chainId) {
      let obj = {url}
      if(chainId) {
        Object.assign(obj, {chainId})
      }
      const dataStr = JSON.stringify({
        type: 'changeChain',
        funcName: 'h5ParseJson',
        data: JSON.stringify(obj)
      });
      this.toNative(dataStr);
    },
    // 和原生通信
    toNative(dataStr) {
      var u = navigator.userAgent;
      var isAndroid = u.indexOf('Android') > -1 || u.indexOf('Adr') > -1;
      let res = '';
      // onCall为和原生端定义的方法名，这个属性可变
      if (isAndroid) {
        res = window.postMessage.onCall(dataStr);
      } else {
        res = window.webkit.messageHandlers.onCall.postMessage(dataStr);
      }
      return res;
    },
  }
}
</script>

<style lang="scss">
@import '//at.alicdn.com/t/font_2611671_qucasfx1j2.css';
.bridge_container {
  min-height: 100vh;
  background-color: #08162C;
  padding: 36px;
  box-sizing: border-box;
  .addr_wrap{
    padding: 12px;
    background-color: #2B2D41;
    border-radius: 28px;
  }
  .section{
    background-color: #0C1D3B;
    border-radius: 10px;
    padding: 30px;
  }
  .dot_icon{
    width: 12px;
    height: 123px;
    margin-right: 25px;
  }
  .border{
    padding-bottom: 35px;
    margin-bottom: 35px;
    border-bottom: 1px dashed #164495;
  }
  .arrow_down{
    color: #8CAFFF;
    font-size: 20px;
  }
  .bridge_input_wrap{
    border: 1px solid #6177D0;
    border-radius: 10px;
    height: 110px;
    position: relative;
    input{
      width: 100%;
      height: 100%;
      margin: 0;
      padding: 0;
      display: block;
      box-sizing: border-box;
      padding: 0 120px 0 26px;
      margin-right: 80px;
      background-color: transparent;
      outline: none;
      border: none;
    }
    .all{
      position: absolute;
      right: 26px;
      top: 50%;
      transform: translateY(-50%);
    }
  }
  .amount_wrap{
    height: 40px;
  }
  .link_wrap{
      background-color: #003183;
      border-radius: 6px;
      padding-left: 14px;
      padding-right: 14px;
      margin-left: 20px;
      &.blue_bg{
        background-color: #27ABE0;
      }&.grey_bg{
        background-color: #131A27;
      }
    }
    .popup_submit{
      background-color: #436AFF;
      height: 88px;
      border-radius: 10px;
      width: 100%;
      margin-bottom: 28px;
    }
  .bridge_submit_btn{
    background-color: #436AFF;
    height: 100px;
    border-radius: 10px;
  }
  .popup_container{
    width: calc(100vw - 150px);
    .popup_head{
      padding: 30px 30px 35px 30px;
      border-bottom: 1px solid #F4F4F4;
    }
    .list_wrap{
      padding-left: 30px;
      padding-right: 30px;
      padding-bottom: 30px;
      max-height: 500px;
    }
    .popup_item{
      padding-top: 40px;
      padding-bottom: 40px;
      border-bottom: 1px solid #F4F4F4;
      .coin_logo_big{
        width: 54px;
        height: 54px;
        margin-right: 10px;
        border-radius: 50%;
      }
      .checked{
        color: #436AFF;
      }
    }
  }
  .loading-wrap{
    background: #ffffff;
    border-radius: 10px;
    padding-bottom: 20px;
    img{
      width: 40%;
      border-radius: 10px;
    }
  }

  $fColor_4570B3: #4570B3;
  $fColor_ffffff: #ffffff;
  $fColor_5E81BB: #5E81BB;
  $fColor_587AFF: #587AFF;
  $fColor_253B80: #253B80;
  $fColor_6391DE: #6391DE;
  $fColor_ACB6FF: #ACB6FF;
  $fColor_08162C: #08162C;
  $fColor_333333: #333333;
  $fColor_999999: #999999;
  $fSize_20: 20px;
  $fSize_22: 22px;
  $fSize_24: 24px;
  $fSize_26: 26px;
  $fSize_28: 28px;
  $fSize_30: 30px;
  $fSize_32: 32px;
  $fSize_48: 48px;
  .fStyle20_ACB6FF{
    font-size: $fSize_20;
    color: $fColor_ACB6FF;
  }
  .fStyle20_FFFFFF{
    font-size: $fSize_20;
    color: $fColor_ffffff;
  }
  .fStyle22_999999{
    font-size: $fSize_22;
    color: $fColor_999999;
  }
  .fStyle22_ffffff{
    font-size: $fSize_22;
    color: $fColor_ffffff;
  }
  .fStyle24_6391DE{
    font-size: $fSize_24;
    color: $fColor_6391DE;
  }
  .fStyle24_ffffff{
    font-size: $fSize_24;
    color: $fColor_ffffff;
  }
  .fStyle26_5E81BB{
    font-size: $fSize_26;
    color: $fColor_5E81BB;
  }
  .fStyle26_333333{
    font-size: $fSize_26;
    color: $fColor_333333;
  }
  .fStyle28_4570B3{
    font-size: $fSize_28;
    color: $fColor_4570B3;
  }
  .fStyle28_333333_w5{
    font-size: $fSize_28;
    color: $fColor_333333;
    font-weight: 500;
  }
  .fStyle30_ffffff{
    font-size: $fSize_30;
    color: $fColor_ffffff;
  }
  .fStyle30_587AFF{
    font-size: $fSize_30;
    color: $fColor_587AFF;
  }
  .fStyle30_253B80{
    font-size: $fSize_30;
    color: $fColor_253B80;
  }
  .fStyle32_ffffff{
    font-size: $fSize_32;
    color: $fColor_ffffff;
  }
  .fStyle32_08162C_w5{
    font-size: $fSize_32;
    font-weight: 500;
    color: $fColor_08162C;
  }
  .fStyle48_333333_w6{
    font-size: $fSize_48;
    font-weight: 600;
    color: $fColor_333333;
  }

  .logo_img{
    width: 291px;
    height: 58px;
  }
  .wallet_icon{
    width: 27px;
    height: 23px;
    margin-left: 30px;
  }
  .coin_logo{
    width: 44px;
    height: 44px;
    margin-right: 18px;
    border-radius: 50%;
  }
  .coin_logo_min{
    width: 32px;
    height: 32px;
    margin-right: 12px;
    margin-left: 12px;
    border-radius: 50%;
  }
  .coin_logo_miner{
    width: 24px;
    height: 24px;
    margin-right: 6px;
    margin-left: 6px;
    border-radius: 50%;
  }
  .exchange_icon{
    width: 30px;
    height: 36px;
    margin-left: 45px;
  }
  .bridge_bg{
    height: 418px;
    background-image: url('../assets/bridge/bg.png');
    background-size: 408px 418px;
    background-repeat: no-repeat;
    background-position: right;
  }
  .arrow_down_icon{
    width: 30px;
    height: 36px;
    margin-top: 25px;
    margin-bottom: 25px;
  }

  @font-face {
    font-family: 'iconfont';  /* Project id 2611671 */
    src: url('//at.alicdn.com/t/font_2611671_qucasfx1j2.woff2?t=1623738550612') format('woff2'),
        url('//at.alicdn.com/t/font_2611671_qucasfx1j2.woff?t=1623738550612') format('woff'),
        url('//at.alicdn.com/t/font_2611671_qucasfx1j2.ttf?t=1623738550612') format('truetype');
  }
  img{
    display: block;
  }
}
</style>