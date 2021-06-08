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

    struct Token {
        bool isRun; // 是否运行
        bool isMain; // 是否主链
        address local; // 本链地址
    }

    // 代币列表[侧链][侧链代币地址] = 代币信息
    mapping(string => mapping(address => Token)) public tokens;

    event adminChanged(address _address);

    modifier onlyAdmin() {
        require(msg.sender == admin, "only use admin to call");
        _;
    }

    function setAdmin(address payable newAdmin) public onlyAdmin {
        admin = newAdmin;
    }

    // 代币是否支持跨链
    function tokenCanBridge(string memory chain, address token) view internal returns (bool){
        return tokens[chain][token].isRun;
    }


    // 侧链,侧链代币,本链代币,是否运行,本链是否主链
    function tokenInsert(string memory chain, address remote, address local, bool isRun, bool isMain) external onlyAdmin {
        tokens[chain][remote] = Token({
        isRun : isRun,
        isMain : isMain,
        local : local
        });
    }

    // 设置代币状态
    function setTokenIsRun(string memory chain, address remote, bool state) public {
        require(
            msg.sender == admin,
            "No operation permission"
        );

        tokens[chain][remote].isRun = state;
    }

    // 资产转账
    function tokenTransfer(address local, address recipient, uint256 value) public onlyAdmin {
        IERC20 token = IERC20(local);
        token.transfer(recipient, value);
    }
}

contract Bridge is BridgeAdmin {

    address public owner;

    event Deposit(string chain, address remote, address recipient, uint256 value);

    event WithdrawDone(address local,address remote, address recipient, uint256 value);

    modifier onlyOwner {
        require(msg.sender == owner, "only owner can call this function");
        _;
    }

    modifier canBridge(string memory chain, address _token) {
        require(tokenCanBridge(chain, _token), "token is can not use bridge");
        _;
    }

    constructor() {
        admin = msg.sender;
        owner = msg.sender;
    }

    function setOwner(address payable newOwner) public onlyOwner {
        owner = newOwner;
    }

    // 本链兑换外链代币
    function deposit(
        string memory chain,
        address remote,
        uint256 value
    ) public canBridge(chain, remote) {
        Token storage local = tokens[chain][remote];
        IERC20 token = IERC20(local.local);
        token.transferFrom(msg.sender, address(this), value);
        // uint balance = token.balanceOf(address(this));
        if (!local.isMain) {
            // 侧链 燃烧
            token.burn(value);
        }
        emit Deposit(chain, local.local, msg.sender, value);
    }

    // 外链兑换本链代币 [管理员]
    function withdraw(
        string memory chain,
        address remote,
        address recipient,
        uint256 value
    ) public onlyOwner {
        Token storage local = tokens[chain][remote];
        IERC20 token = IERC20(local.local);
        if (local.isMain) {
            // 主链 转账
            token.transfer(recipient, value);
        } else {
            // 侧链 铸币
            token.mint(recipient, value);
        }
        emit WithdrawDone(local.local,remote, recipient, value);
    }
}
