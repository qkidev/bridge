名称说明：

原生网：一个资产的来源主网，一个主网可能在一个资产下面是原生网，在另外一个主网下是跨链网。

跨链网：一个资产跨链出去的主网，一个主网可能在一个资产下面是原生网，在另外一个主网下是跨链网。

资产：一种token。

跨链桥合约：由跨链钱合约管理增加的跨链，各种跨链网的资产owner都是跨链桥，然后跨链桥可以进行mint和burn。

数据：

资产列表：
合约地址=> {是否是原生网}

跨链列表：
当前合约=>目标合约=>true

注册资产：把一个资产注册到跨链桥后，才可以进行跨链，原生网和跨链网都需要注册才能使用。

参数：本主网是否是原生网、本主网token合约、目标主网合约地址

逻辑：写入资产到资产列表，跨链列表

存入资产：把一个资产存入跨链桥，进入另外一个主网。如果当前主网是资产的原生网，资产进行质押，如果是跨链网就burn。

参数：token合约、数量、目标主网、目标主网token合约地址

事件：token合约、数量、目标主网、目标主网token合约地址

逻辑：根据 跨链列表 判断跨链网络是否支持，支持就操作资产。

取出资产：把跨链资产在本主网进行取出，由跨链桥管理员执行。如果当前主网是资产的原生网，把资产transfer给地址，如果是跨链网就mint。

参数：token合约、数量、资产接收地址、存入资产的tx hash
    
事件：token合约、数量、资产接收地址、存入资产的tx hash

发布token（可选）：自动发布token，并成为管理员，可以做到注册资产里面，如果本主网token合约地址是空，就自动发布一个token

跨链桥管理员：可以是一个单独的钱包地址，也可以是一个合约地址，这个合约由支持多签、支持去中心化。

特别说明：只支持原生网和跨链网双向互转，跨链网和跨链网不支持。

token合约：标准erc20 token，带管理员功能，可以mint、burn。


运行流程:
接收跨链请求,目标资产是否允许跨链,授权转账给桥地址,发送跨链事件给中间节点,
中间节点接收跨链事件(入库),节点操作目标链合约转账目标资产给目标地址(入库),接收交易完成通知(入库).
节点定时查询记录,存在跨出未到账的情况尝试再次转账到地址.

系统流程:
主链部署代币,主链部署桥,侧链部署代币,侧链部署桥,侧链将代币的铸币权限给侧链桥,桥将代币的创建者添加到代币创建者列表中.
用户在桥调用桥的跨链且提供(目标链,目标代币,数额),允许授权资产给桥,用户在目标链查看接收到的代币

中间流程:
给操作账户在链中充值足购的主链币做手续费.
监听桥的跨链事件,调用目标链中桥的提现合约且提供(链,发送链代币地址,接收账号,数额).

用户测试:
选择跨链的代币,请求跨链传入参数,允许转账授权

管理员测试:
部署桥,添加支持的代币,删除支持的代币,设置代币状态,添加白名单,删除白名单,添加代币创建者,删除代币创建者,转移代币到指定账号

节点测试:
查看事件是否完成,数据是否记录

