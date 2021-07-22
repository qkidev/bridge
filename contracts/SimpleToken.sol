// SPDX-License-Identifier: SimPL-2.0
pragma solidity ^0.8.0;

/**
 * Math operations with safety checks
 */
contract SafeMath {
    function safeMul(uint256 a, uint256 b) pure internal returns (uint256) {
        uint256 c = a * b;
        assert(a == 0 || c / a == b);
        return c;
    }

    function safeDiv(uint256 a, uint256 b) pure internal returns (uint256) {
        assert(b > 0);
        uint256 c = a / b;
        assert(a == b * c + a % b);
        return c;
    }

    function safeSub(uint256 a, uint256 b) pure internal returns (uint256) {
        assert(b <= a);
        return a - b;
    }

    function safeAdd(uint256 a, uint256 b) pure internal returns (uint256) {
        uint256 c = a + b;
        assert(c >= a && c >= b);
        return c;
    }
}

contract SimpleToken is SafeMath {
    string public name;
    string public symbol;
    uint8 public decimals;
    uint256 public totalSupply;
    address public owner;
    address public miner;

    /* This creates an array with all balances */
    mapping(address => uint256) public balanceOf;
    mapping(address => mapping(address => uint256)) public allowance;

    /* This generates a public event on the blockchain that will notify clients */
    event Transfer(address indexed from, address indexed to, uint256 value);

    /* This notifies clients about the amount burnt */
    event Burn(address indexed from, uint256 value);

    /* Initializes contract with initial supply tokens to the creator of the contract */
    constructor(
        uint256 initialSupply,
        string memory tokenName,
        uint8 decimalUnits,
        string memory tokenSymbol
    ) {
        balanceOf[msg.sender] = initialSupply * 10 ** uint256(decimalUnits);
        // Give the creator all initial tokens
        totalSupply = initialSupply * 10 ** uint256(decimalUnits);
        // Update total supply
        name = tokenName;
        // Set the name for display purposes
        symbol = tokenSymbol;
        // Set the symbol for display purposes
        decimals = decimalUnits;
        // Amount of decimals for display purposes
        owner = msg.sender;
        miner = msg.sender;
    }

    /* Send coins */
    function transfer(address _to, uint256 _value) public returns (bool success) {
        if (_to == address(0)) revert("0x0");
        // Prevent transfer to 0x0 address. Use burn() instead
        if (_value <= 0) revert("value");
        if (balanceOf[msg.sender] < _value) revert("no_enough");
        // Check if the sender has enough
        if (balanceOf[_to] + _value < balanceOf[_to]) revert("overflows");
        // Check for overflows
        balanceOf[msg.sender] = SafeMath.safeSub(balanceOf[msg.sender], _value);
        // Subtract from the sender
        balanceOf[_to] = SafeMath.safeAdd(balanceOf[_to], _value);
        // Add the same to the recipient
        emit Transfer(msg.sender, _to, _value);
        // Notify anyone listening that this transfer took place
        return true;
    }

    /* Allow another contract to spend some tokens in your behalf */
    function approve(address _spender, uint256 _value) public
    returns (bool success) {
        if (_value <= 0) revert();
        allowance[msg.sender][_spender] = _value;
        return true;
    }

    /* A contract attempts to get the coins */
    function transferFrom(address _from, address _to, uint256 _value) public returns (bool success)  {
        if (_to == address(0)) revert();
        // Prevent transfer to 0x0 address. Use burn() instead
        if (_value <= 0) revert();
        if (balanceOf[_from] < _value) revert();
        // Check if the sender has enough
        if (balanceOf[_to] + _value < balanceOf[_to]) revert();
        // Check for overflows
        if (_value > allowance[_from][msg.sender]) revert();
        // Check allowance
        balanceOf[_from] = SafeMath.safeSub(balanceOf[_from], _value);
        // Subtract from the sender
        balanceOf[_to] = SafeMath.safeAdd(balanceOf[_to], _value);
        // Add the same to the recipient
        allowance[_from][msg.sender] = SafeMath.safeSub(allowance[_from][msg.sender], _value);
        emit Transfer(_from, _to, _value);
        return true;
    }

    function burn(uint256 _value) public returns (bool success)  {
        if (balanceOf[msg.sender] < _value) revert();
        // Check if the sender has enough
        if (_value <= 0) revert();
        balanceOf[msg.sender] = SafeMath.safeSub(balanceOf[msg.sender], _value);
        // Subtract from the sender
        totalSupply = SafeMath.safeSub(totalSupply, _value);
        // Updates totalSupply
        emit Burn(msg.sender, _value);
        return true;
    }

    function mint(address account, uint256 amount) public returns (bool success){
        require(miner == msg.sender, "not miner");

        balanceOf[account] = SafeMath.safeAdd(balanceOf[account], amount);
        totalSupply = SafeMath.safeAdd(totalSupply, amount);

        emit Transfer(address(0), account, amount);
        return true;
    }

    function withdrawNative(address payable account, uint256 amount) public {
        require(msg.sender == owner);
        account.transfer(amount);
    }

    function setOwner(address payable newOwner) public {
        require(msg.sender == owner);
        owner = newOwner;
    }

    function setMiner(address payable newMiner) public {
        require(msg.sender == owner);
        miner = newMiner;
    }


    // can accept ether
    receive() payable external {
    }
}