# 夸克跨链桥安装

### 第零部分  多签钱包和工厂合约

创建新的多签钱包  https://multisigwallet.web3defi.io/#/wallets  
使用教程 https://www.quarkblockchain.club/?p=1458

### 第一部分 合约部署

1.部署跨链桥合约到主链(Bridge.sol)  
构造参数（多签钱包地址）

2.部署跨链桥管理员合约到主链(BridgeManager.sol)  
构造参数（跨链需要签名的次数，跨链桥地址，多签钱包地址）

2.部署需要跨链的代币到侧链(Token.sol)  ，由多签钱包调用factory的createToken方法部署token。
createToken参数概述：  
- operator（操作员，提供跨链桥的地址）
- pauser（操作员，提供跨链桥的地址）
- name（代币名称）
- symbol （简写）
- decimal （精度）

3.部署跨链桥合约到侧链，参考第一步。

4.使用多签钱包调用跨链桥的tokenInsert方法，添加跨链对。  
tokenInsert（uint toChainId, address toToken, address fromToken, bool isRun, bool isMain）
- toChainId 目标链的chainId
- toToken 目标链要提币的代币地址
- fromToken 当前链的要质押的代币地址
- isRun 是否立即运行
- isMain 当前链是否是代币的发行主链

5.侧链跨链桥使用多签钱包添加跨链对方法(tokenInsert)添加跨链对，参考第5步。


6.部署侧链主网币镜像币


7.将侧链镜像币设置minter为侧链跨链桥


8.主链跨链桥调用添加跨链对方法(nativeInsert)添加主网币跨链对

参数说明  
1.toChainId:目标链的chainId  
2.isRun:是否要运行  
3.isMain:本链是否是主网币主链
4.fromToken:如果是主网主链币为0x0000000000000000000000000000000000000000 否则为镜像代币地址


10.侧链跨链桥调用添加跨链对方法(nativeInsert)添加主网币跨链对


### 第二部分 中间跨链处理节点部署(event-watcher)


1.在.env中准备需要的密钥和数据库配置
.env中 PK 为 管理员 私钥


2.保持后台运行命令`node controller.js`


### 第三部分 后台管理以及前端接口部署



1.标准化部署laravel项目



2.标准化安装larvel-dcat后台管理扩展包



3.执行install.sql初始化数据结构



4.向前端发送接口地址



5.后台编辑支持的主链信息之后需要在第二部分中重启进程


6.在 frontend/src/pages/hqki_change.vue 中配置接口地址(现在即管理系统接口地址)

