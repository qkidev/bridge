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
}

// 1.部署跨链桥合约
// 2.脚本开始监听跨链桥合约代币转账接收
// 3.脚本接收跨链桥兑外链代币事件
// 4.脚本调用外链的充值事件
contract Bridge {
    address public owner;
    using SafeMath for uint256;

    // 发送代币到外网
    event Sended(string, address, address, uint256);
    // 代币-账号-余额
    mapping(address => mapping(address => uint256)) public balances;

    // 账号
    struct Account {
        address owner; // 所有者
        string network; // 主网
        string name; // 名称
        address account;
    }

    // 全部账号
    Account[] private accounts;

    // 资产
    struct Assert {
        string name; // 名称
        address token; // 地址
        uint256 balance; // 兑换仓剩余数额
    }

    Assert[] public asserts;

    modifier onlyOwner {
        require(msg.sender == owner, "Only owner can call this function.");
        _;
    }

    constructor() {
        owner = msg.sender;
    }

    // 对比字符串
    function compareStr(string memory _str1, string memory _str2)
        public
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

    // 给自己添加一个外网账号
    function account_insert(
        string memory _network,
        string memory _name,
        address _account
    ) public {
        require(_account != address(0), "Account address cannot be empty");
        accounts.push(Account(msg.sender, _network, _name, _account));
    }

    function set_owner(address _address) public onlyOwner {
        owner = _address;
    }

    // 兑换外网资产到指定地址
    function exchange_to_account(
        string memory _network,
        address _token,
        address _account,
        uint256 _value
    ) public {
        address _rAccount = address(0);
        for (uint256 i = 0; i < accounts.length; i++) {
            if (
                accounts[i].owner == msg.sender &&
                accounts[i].account == _account
            ) {
                _rAccount = accounts[i].account;
            }
        }
        require(_rAccount != address(0), "Invalid receive address");
        for (uint256 i = 0; i < asserts.length; i++) {
            if (
                asserts[i].token == _token
            ) {
                asserts[i].balance = asserts[i].balance.sub(_value);
                break;
            }
        }
        balances[_token][msg.sender] = balances[_token][msg.sender].sub(_value);
        emit Sended(_network, _token, _account, _value);
    }

    // 兑换外网资产
    function exchange(
        string memory _network,
        address _token,
        uint256 _value
    ) public {
        uint256 _balance = 0;
        for (uint256 i = 0; i < asserts.length; i++) {
            if (
                asserts[i].token == _token
            ) {
                _balance = asserts[i].balance;
                break;
            }
        }
        require(
            _balance > _value,
            "The reserve of the exchange pool is insufficient and cannot be exchanged"
        );
        
        balances[_token][msg.sender] = balances[_token][msg.sender].sub(_value);
        emit Sended(_network, _token, msg.sender, _value);
    }

    // 外链充值入本链发送代币[管理员]
    function send(
        address _token,
        address _address,
        uint256 _value
    ) public onlyOwner {
        IERC20 token = IERC20(_token);
        token.transfer(_address, _value);
        for (uint256 i = 0; i < asserts.length; i++) {
            if (
                asserts[i].token == _token
            ) {
                asserts[i].balance = asserts[i].balance.sub(_value);
                break;
            }
        }
        
    }

    // 充值操作[管理员]
    function recharge(
        address _token,
        address _address,
        uint256 _value
    ) public onlyOwner {
        for (uint256 i = 0; i < asserts.length; i++) {
            if (
                asserts[i].token == _token
            ) {
                asserts[i].balance = asserts[i].balance.add(_value);
                break;
            }
        }
        balances[_token][_address] = balances[_token][_address].add(_value);
    }

    // 添加兑换资产
    function assert_insert(
        string memory _name,
        address _token,
        uint256 _value
    ) public onlyOwner {
        asserts.push(
            Assert({
                name: _name,
                token: _token,
                balance: _value
            })
        );
    }

    // 支持的兑换资产数量
    function asserts_length() public view returns (uint256) {
        return asserts.length;
    }
}
