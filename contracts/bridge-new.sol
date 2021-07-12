// SPDX-License-Identifier: SimPL-2.0
pragma solidity ^0.7.0;

/**
 * @dev Interface of the ERC20 standard as defined in the EIP.
 */
interface IERC20 {
    /**
     * @dev Returns the amount of tokens in existence.
     */
    function totalSupply() external view returns (uint256);

    /**
     * @dev Returns the amount of tokens owned by `account`.
     */
    function balanceOf(address account) external view returns (uint256);

    /**
     * @dev Moves `amount` tokens from the caller's account to `recipient`.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {Transfer} event.
     */
    function transfer(address recipient, uint256 amount)
    external
    returns (bool);

    /**
     * @dev Returns the remaining number of tokens that `spender` will be
     * allowed to spend on behalf of `owner` through {transferFrom}. This is
     * zero by default.
     *
     * This value changes when {approve} or {transferFrom} are called.
     */
    function allowance(address owner, address spender)
    external
    view
    returns (uint256);

    /**
     * @dev Sets `amount` as the allowance of `spender` over the caller's tokens.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * IMPORTANT: Beware that changing an allowance with this method brings the risk
     * that someone may use both the old and the new allowance by unfortunate
     * transaction ordering. One possible solution to mitigate this race
     * condition is to first reduce the spender's allowance to 0 and set the
     * desired value afterwards:
     * https://github.com/ethereum/EIPs/issues/20#issuecomment-263524729
     *
     * Emits an {Approval} event.
     */
    function approve(address spender, uint256 amount) external returns (bool);

    /**
     * @dev Moves `amount` tokens from `sender` to `recipient` using the
     * allowance mechanism. `amount` is then deducted from the caller's
     * allowance.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {Transfer} event.
     */
    function transferFrom(
        address sender,
        address recipient,
        uint256 amount
    ) external returns (bool);

    /**
     * @dev Emitted when `value` tokens are moved from one account (`from`) to
     * another (`to`).
     *
     * Note that `value` may be zero.
     */
    event Transfer(address indexed from, address indexed to, uint256 value);

    /**
     * @dev Emitted when the allowance of a `spender` for an `owner` is set by
     * a call to {approve}. `value` is the new allowance.
     */
    event Approval(
        address indexed owner,
        address indexed spender,
        uint256 value
    );

    function mint(address account, uint256 amount) external; // 铸币

    function burn(uint256 _value) external returns (bool success); // 燃烧
}

contract BridgeAdmin {

    address public admin;

    address[] public managers;

    // 需要多签数量
    uint public sigNum;

    // 设置多签数量
    function setSigNum(uint num) public onlyAdmin {
        sigNum = num;
    }

    struct Token {
        bool isRun; // 是否运行
        bool isMain; // 是否主链
        address local; // 本链地址
    }


    // 代币列表[侧链][侧链代币地址] = 代币信息
    mapping(uint => mapping(address => Token)) public tokens;

    // 主网币列表[主网][是主链] = 主网代币信息
    mapping(uint => mapping(bool => Token)) public natives;

    // 主网ID[质押hash] = 管理员地址[]
    mapping(uint => mapping(bytes => address[])) public multiSigs;

    event managerAdded(address _address);

    event managerDelete(address _address);

    event adminChanged(address _address);

    modifier onlyAdmin() {
        require(msg.sender == admin, "Bridge Admin: only use admin to call");
        _;
    }

    modifier onlyManager {
        bool isManager = false;
        for (uint i = 0; i < managers.length; i++) {
            if (managers[i] == msg.sender) {
                isManager = true;
                break;
            }
        }
        require(isManager, "Bridge Admin: only manager can call this function");
        _;
    }


    // 添加管理员
    function addManager(address payable manager) public onlyAdmin {
        managers.push(manager);
    }

    // 设置最高级管理员
    function setAdmin(address payable newAdmin) public onlyAdmin {
        admin = newAdmin;
    }

    // 跨出多签
    function withdrawMultiSign(uint toChainId, bytes memory depositHash) internal returns (bool) {
        address[] storage signs = multiSigns[toChainId][depositHash];
        bool isSign = false;
        for (uint i = 0; i < signs.length; i++) {
            if (signs[i] == msg.sender) {
                isSign=true;
            }
        }

        if (!isSign) {
            multiSigns[toChainId][depositHash].push(msg.sender);
        }

        return multiSigns[toChainId][depositHash].length >= sigNum;
    }

    // 代币是否支持跨链
    function tokenCanBridge(uint toChainId, address toToken) view internal returns (bool){
        return tokens[toChainId][toToken].isRun;
    }

    // 地址是否为合约
    function isContract(address addr) internal view returns (bool) {
        uint256 size;
        assembly {size := extcodesize(addr)}
        return size > 0;
    }

    // 侧链,侧链代币,本链代币,是否运行,本链是否主链
    function tokenInsert(uint toChainId, address toToken, address fromToken, bool isRun, bool isMain) external onlyAdmin {
        tokens[toChainId][toToken] = Token({
        isRun : isRun,
        isMain : isMain,
        local : fromToken
        });
    }

    // 添加支持的主网币
    function nativeInsert(uint toChainId, bool isRun, bool isMain, address fromToken) external onlyAdmin {
        if (isMain) {
            require(fromToken == address(0), "main native have not token");
        } else {
            require(isContract(fromToken), "address is not a contract");
            require(fromToken != address(0), "minor native must have token");
        }
        natives[toChainId][isMain] = Token({
        isMain : isMain,
        isRun : isRun,
        local : fromToken
        });
    }

    // 设置代币状态
    function setTokenIsRun(uint toChainId, address toToken, bool state) public {
        require(
            msg.sender == admin,
            "No operation permission"
        );

        tokens[toChainId][toToken].isRun = state;
    }

    // 设置主网币状态
    function setNativeIsRun(uint toChainId, bool isMain, bool state) public {
        require(
            msg.sender == admin,
            "No operation permission"
        );
        natives[toChainId][isMain].isRun = state;
    }


    // 资产转账
    function tokenTransfer(address fromToken, address recipient, uint256 value) public onlyAdmin {
        IERC20 token = IERC20(fromToken);
        token.transfer(recipient, value);
    }

    // 主网币转账
    function nativeTransfer(address payable recipient, uint256 value) public onlyAdmin {
        require(address(this).balance > value, "not enough native token");
        recipient.transfer(value);
    }
}

contract Bridge is BridgeAdmin {

    address public owner;

    event Deposit(uint toChainId, address fromToken, address toToken, address recipient, uint256 value);

    event DepositNative(uint toChainId, bool isMain, address recipient, uint256 value);

    event WithdrawDone(uint toChainId, address fromToken, address toToken, address recipient, uint256 value);

    event WithdrawNativeDone(uint fromChainId, address recipient, bool isMain, uint256 value);

    modifier onlyOwner {
        require(msg.sender == owner, "only owner can call this function");
        _;
    }



    modifier canBridge(uint chainId, address toToken) {
        require(tokenCanBridge(chainId, toToken), "token is can not use bridge");
        _;
    }


    constructor() {
        admin = msg.sender;
        owner = msg.sender;
        managers.push(msg.sender); // new manager
        sigNum=1;// multi sign limit number
    }

    function setOwner(address payable newOwner) public onlyOwner {
        owner = newOwner;
    }

    // token 跨出
    function deposit(
        uint toChainId,
        address toToken,
        uint256 value
    ) public canBridge(toChainId, toToken) {
        Token storage local = tokens[toChainId][toToken];
        IERC20 token = IERC20(local.local);
        token.transferFrom(msg.sender, address(this), value);
        if (!local.isMain) {
            // 侧链 燃烧
            token.burn(value);
        }
        emit Deposit(toChainId, local.local, toToken, msg.sender, value);
    }

    // token 跨入
    function withdraw(
        uint toChainId,
        address toToken,
        address recipient,
        uint256 value,
        bytes memory hash
    ) public onlyManager {
        bool isMultiSig = withdrawMultiSign(toChainId, hash);
        if (isMultiSig) {
            Token storage local = tokens[toChainId][toToken];
            IERC20 token = IERC20(local.local);
            if (local.isMain) {
                // 主链 转账
                token.transfer(recipient, value);
            } else {
                // 侧链 铸币
                token.mint(recipient, value);
            }
            emit WithdrawDone(toChainId, local.local, toToken, recipient, value);
        }

    }

    // 主网币跨出
    function depositNative(uint toChainId, bool isMain, uint256 value) public payable {
        Token storage native = natives[toChainId][isMain];
        require(native.isRun, "chain is not support");

        if (native.isMain) {
            // 主链跨出
            require(msg.value == value, "value is error");
            require(msg.value > 0, "value is zero");
        } else {
            // 侧链跨出
            IERC20 token = IERC20(native.local);
            token.transferFrom(msg.sender, address(this), value);
        }
        emit DepositNative(toChainId, isMain, msg.sender, value);
    }

    // 主网币跨入
    function withdrawNative(uint toChainId, address payable recipient, bool isMain, uint256 value, bytes memory hash) public onlyManager {

        bool isMultiSig = withdrawMultiSign(toChainId, hash);
        if (isMultiSig) {
            Token storage native = natives[toChainId][isMain];
            require(native.isRun, "chain is not support");
            if (native.isMain) {
                // 主链跨入
                require(address(this).balance > value, "not enough native token");
                recipient.transfer(value);
            } else {
                // 侧链跨入
                IERC20 token = IERC20(native.local);
                token.mint(recipient, value);
            }
            emit WithdrawNativeDone(toChainId, recipient, isMain, value);
        }
    }

    receive() external payable {}
}
