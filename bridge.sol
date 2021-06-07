// SPDX-License-Identifier: SimPL-2.0
pragma solidity ^0.7.0;

library SafeMath {
    /**
     * @dev Returns the addition of two unsigned integers, reverting on
     * overflow.
     *
     * Counterpart to Solidity's `+` operator.
     *
     * Requirements:
     *
     * - Addition cannot overflow.
     */
    function add(uint256 a, uint256 b) internal pure returns (uint256) {
        uint256 c = a + b;
        require(c >= a, "SafeMath: addition overflow");

        return c;
    }

    /**
     * @dev Returns the subtraction of two unsigned integers, reverting on
     * overflow (when the result is negative).
     *
     * Counterpart to Solidity's `-` operator.
     *
     * Requirements:
     *
     * - Subtraction cannot overflow.
     */
    function sub(uint256 a, uint256 b) internal pure returns (uint256) {
        return sub(a, b, "SafeMath: subtraction overflow");
    }

    /**
     * @dev Returns the subtraction of two unsigned integers, reverting with custom message on
     * overflow (when the result is negative).
     *
     * Counterpart to Solidity's `-` operator.
     *
     * Requirements:
     *
     * - Subtraction cannot overflow.
     */
    function sub(
        uint256 a,
        uint256 b,
        string memory errorMessage
    ) internal pure returns (uint256) {
        require(b <= a, errorMessage);
        uint256 c = a - b;

        return c;
    }

    /**
     * @dev Returns the multiplication of two unsigned integers, reverting on
     * overflow.
     *
     * Counterpart to Solidity's `*` operator.
     *
     * Requirements:
     *
     * - Multiplication cannot overflow.
     */
    function mul(uint256 a, uint256 b) internal pure returns (uint256) {
        // Gas optimization: this is cheaper than requiring 'a' not being zero, but the
        // benefit is lost if 'b' is also tested.
        // See: https://github.com/OpenZeppelin/openzeppelin-contracts/pull/522
        if (a == 0) {
            return 0;
        }

        uint256 c = a * b;
        require(c / a == b, "SafeMath: multiplication overflow");

        return c;
    }

    /**
     * @dev Returns the integer division of two unsigned integers. Reverts on
     * division by zero. The result is rounded towards zero.
     *
     * Counterpart to Solidity's `/` operator. Note: this function uses a
     * `revert` opcode (which leaves remaining gas untouched) while Solidity
     * uses an invalid opcode to revert (consuming all remaining gas).
     *
     * Requirements:
     *
     * - The divisor cannot be zero.
     */
    function div(uint256 a, uint256 b) internal pure returns (uint256) {
        return div(a, b, "SafeMath: division by zero");
    }

    /**
     * @dev Returns the integer division of two unsigned integers. Reverts with custom message on
     * division by zero. The result is rounded towards zero.
     *
     * Counterpart to Solidity's `/` operator. Note: this function uses a
     * `revert` opcode (which leaves remaining gas untouched) while Solidity
     * uses an invalid opcode to revert (consuming all remaining gas).
     *
     * Requirements:
     *
     * - The divisor cannot be zero.
     */
    function div(
        uint256 a,
        uint256 b,
        string memory errorMessage
    ) internal pure returns (uint256) {
        require(b > 0, errorMessage);
        uint256 c = a / b;
        // assert(a == b * c + a % b); // There is no case in which this doesn't hold

        return c;
    }

    /**
     * @dev Returns the remainder of dividing two unsigned integers. (unsigned integer modulo),
     * Reverts when dividing by zero.
     *
     * Counterpart to Solidity's `%` operator. This function uses a `revert`
     * opcode (which leaves remaining gas untouched) while Solidity uses an
     * invalid opcode to revert (consuming all remaining gas).
     *
     * Requirements:
     *
     * - The divisor cannot be zero.
     */
    function mod(uint256 a, uint256 b) internal pure returns (uint256) {
        return mod(a, b, "SafeMath: modulo by zero");
    }

    /**
     * @dev Returns the remainder of dividing two unsigned integers. (unsigned integer modulo),
     * Reverts with custom message when dividing by zero.
     *
     * Counterpart to Solidity's `%` operator. This function uses a `revert`
     * opcode (which leaves remaining gas untouched) while Solidity uses an
     * invalid opcode to revert (consuming all remaining gas).
     *
     * Requirements:
     *
     * - The divisor cannot be zero.
     */
    function mod(
        uint256 a,
        uint256 b,
        string memory errorMessage
    ) internal pure returns (uint256) {
        require(b != 0, errorMessage);
        return a % b;
    }
}

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

    struct White {
        string chain; // 侧链
        address remote; // 代币
        string symbol;
    }

    struct Token {
        bool isRun; // 是否运行
        bool isMain; // 是否主链
        address local; // 本链地址
    }

    White[] public whites;

    // 代币列表[侧链][侧链代币地址] = 代币信息
    mapping(string => mapping(address => Token)) public tokens;

    // 代币创建者列表[本链代币] = 创建者地址
    mapping(address => address) public tokenCreators;

    event adminChanged(address _address);

    modifier onlyAdmin() {
        require(msg.sender == admin, "only use admin to call");
        _;
    }

    // 对比字符串
    function compareStr(string memory _str1, string memory _str2)
    internal
    pure
    returns (bool)
    {
        if (bytes(_str1).length == bytes(_str2).length) {
            if (
                keccak256(abi.encodePacked(_str1)) ==
                keccak256(abi.encodePacked(_str2))
            ) {
                return true;
            }
        }
        return false;
    }

    // 代币创建者添加
    function tokenCreatorInsert(address local, address creator) public onlyAdmin {
        tokenCreators[local] = creator;
    }

    // 代币创建者删除
    function tokenCreatorDelete(address local) public onlyAdmin {
        tokenCreators[local] = address(0);
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
    function setTokenIsRun(string memory chain, address remote,bool state) public {
        Token storage local = tokens[chain][remote];

        // 验证管理员或者代币创建者
        require(
            msg.sender == admin || tokenCreators[local.local] == msg.sender,
            "No operation permission"
        );

        tokens[chain][remote].isRun = state;
    }

    // 白名单添加
    function whiteInsert(string memory chain, address remote, string memory symbol) public onlyAdmin {
        whites.push(White({
        chain : chain,
        remote : remote,
        symbol : symbol
        }));
    }

    // 白名单删除
    function whiteDelete(string memory chain, address remote) public onlyAdmin {
        for (uint i = 0; i < whites.length; i++) {
            if (compareStr(whites[i].chain, chain) && whites[i].remote == remote)
                delete whites[i];
        }
    }

    // 资产转账
    function tokenTransfer(address local,address recipient,uint256 value) public onlyAdmin {
        IERC20 token = IERC20(local);
        token.transfer(recipient,value);
    }
}

contract Bridge is BridgeAdmin {

    address public owner;

    using SafeMath for uint256;


    event Deposit(string, address, address, uint256);

    event WithdrawDone(address, address, uint256);

    modifier onlyOwner {
        require(msg.sender == owner, "only owner can call this function.");
        _;
    }

    modifier canBridge(string memory chain, address _token) {
        require(tokenCanBridge(chain, _token), "token is can not use bridge.");
        _;
    }

    // balanceOf[账号][代币]=余额
    mapping(address => mapping(address => uint256)) public balanceOf;

    constructor() {
        owner = msg.sender;
    }

    function setOwner(address payable newOwner) public onlyOwner {
        owner = newOwner;
    }

    // 中间节点充值代币
    function recharge(address recipient, address local, uint256 value) public onlyOwner {
        balanceOf[recipient][local] = balanceOf[recipient][local].add(value);
    }

    // 本链兑换外链代币
    function deposit(
        string memory chain,
        address remote,
        uint256 value
    ) public onlyOwner canBridge(chain, remote) {


        Token storage local = tokens[chain][remote];

        uint256 balance = balanceOf[msg.sender][local.local];

        require(
            balance > 0,
            "Sorry, your credit is running low"
        );

        if (!local.isMain) {
            // 侧链 燃烧
            IERC20 token = IERC20(local.local);
            token.burn(value);
        }

        balanceOf[msg.sender][local.local] = balanceOf[msg.sender][local.local].sub(value);
        emit Deposit(chain, remote, msg.sender, value);
    }

    // 外链兑换本链代币 [管理员]
    function withdrawToken(
        string memory chain,
        address remote,
        address recipient,
        uint256 value
    ) public onlyOwner {
        Token storage localToken = tokens[chain][remote];
        if (localToken.isMain) {
            // 主链 转账
            IERC20 token = IERC20(localToken.local);
            token.transfer(recipient, value);
        } else {
            // 侧链 铸币
            IERC20 token = IERC20(localToken.local);
            token.mint(recipient, value);
        }
        emit WithdrawDone(localToken.local, recipient, value);
    }
}
// 运行流程:
// 接收跨链请求,检查账号本链中目标资产是否足购,目标资产是否允许跨链,减去本链资产余额,发送跨链事件给中间节点,
// 中间节点接收跨链事件(入库),节点操作目标链合约转账目标资产给目标地址(入库),接收交易完成通知(入库).
// 节点定时查询记录,存在跨出未到账的情况尝试再次转账到地址.

// 系统流程:
// 主链部署代币,主链部署桥,侧链部署代币,侧链部署桥,侧链将代币的铸币权限给侧链桥,桥将代币的创建者添加到代币创建者列表中.
// 用户在主(侧)链转账代币给主链桥,用户在主(侧)链桥调用桥的跨链且提供(目标链,目标代币,数额),用户在侧(主)链查看接收到的代币

// 中间流程:
// 给操作账户在链中充值足购的主链币做手续费,监听桥接收代币事件,将代币充值给发送地址在桥本代币的余额.
// 监听桥的跨链事件,调用目标链中桥的提现合约且提供(链,发送链代币地址,接收账号,数额).

// 用户测试:
// 主链将代币存入桥,获取桥中白名单代币,选择跨链的代币,请求跨链传入参数

// 管理员测试:
// 部署桥,添加支持的代币,删除支持的代币,设置代币状态,添加白名单,删除白名单,添加代币创建者,删除代币创建者,转移代币到指定账号

// 代币创建者测试:
// 设置代币状态

// 节点测试:
// 查看事件是否完成,数据是否记录