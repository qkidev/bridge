支持跨链桥合约需要实现在ERC20标准.

另外附加以下约定的功能

铸币  
function mint(address account, uint256 amount) public

燃烧代币  
function burn(uint256 _value) public returns (bool success)

设置铸币员  
function setMiner(address newMiner) public  