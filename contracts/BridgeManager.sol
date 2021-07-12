// SPDX-License-Identifier: SimPL-2.0
pragma solidity ^0.7.0;

interface Bridge {
    function withdraw(uint toChainId, address toToken, address recipient, uint256 value) external;

    function withdrawNative(uint toChainId, address payable recipient, bool isMain, uint256 value) external;
}

contract BridgeManager {

    address public owner;

    address public bridgeAddress;

    // 全部管理员
    address[] public managers;

    // 需要多签数量
    uint public signLimit;

    // multiSigns[networkID][txHash] = managers[]
    mapping(uint => mapping(bytes => address[])) public multiSigns;

    constructor(uint _signLimit, address _bridgeAddress) {
        owner = msg.sender;
        managers.push(msg.sender);
        signLimit = _signLimit;
        bridgeAddress = _bridgeAddress;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Bridge Manager: only use owner to call");
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
        require(isManager, "Bridge Manager: only manager can call this function");
        _;
    }

    function setOwner(address payable _owner) public onlyOwner {
        owner = _owner;
    }

    function setBridgeAddress(address _bridgeAddress) public onlyOwner {
        bridgeAddress = _bridgeAddress;
    }

    // 设置多签数量
    function setSignLimit(uint num) public onlyOwner {
        signLimit = num;
    }

    function managerAdd(address _address) public onlyOwner {
        managers.push(_address);
    }

    function managerDel(uint index) public onlyOwner {
        managers[index] = address(0);
    }


    function multiSign(uint toChainId, bytes memory hash) internal returns (bool) {
        address[] storage signs = multiSigns[toChainId][hash];
        bool isSign = false;
        for (uint i = 0; i < signs.length; i++) {
            if (signs[i] == msg.sender) {
                isSign = true;
            }
        }

        if (!isSign) {
            multiSigns[toChainId][hash].push(msg.sender);
        }
        return multiSigns[toChainId][hash].length >= signLimit;
    }

    function withdraw(uint toChainId, address toToken, address recipient, uint256 value, bytes memory hash) public onlyManager {
        bool isMultiSign = multiSign(toChainId, hash);
        if (isMultiSign) {
            Bridge bridge = Bridge(bridgeAddress);
            bridge.withdraw(toChainId, toToken, recipient, value);
        }

    }

    function withdrawNative(uint toChainId, address payable recipient, bool isMain, uint256 value, bytes memory hash) public onlyManager {
        bool isMultiSign = multiSign(toChainId, hash);
        if (isMultiSign) {
            Bridge bridge = Bridge(bridgeAddress);
            bridge.withdrawNative(toChainId, recipient, isMain, value);
        }
    }

}