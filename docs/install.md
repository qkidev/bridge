# 夸克跨链桥安装


### 第一部分 合约部署

1.部署跨链桥合约到主链(Bridge.sol)

2.部署跨链桥管理员合约到主链(BridgeManager.sol)

2.部署需要跨链的代币到侧链(Token.sol)

3.部署跨链桥合约到侧链

4.将侧链的代币设置minter为侧链跨链桥

5.主链跨链桥调用添加跨链对方法(tokenInsert)添加跨链对

参数说明  
1.toChainId:目标链的chainId  
2.toToken:目标链代币地址  
3.fromToken:本链代币地址  
4.isRun:是否要运行  
5.isMain:本链是否是代币主链

6.侧链跨链桥电泳添加跨链对方法(tokenInsert)添加跨链对


7.部署侧链主网币镜像币


8.将侧链镜像币设置minter为侧链跨链桥


9.主链跨链桥调用添加跨链对方法(nativeInsert)添加主网币跨链对

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