// SPDX-License-Identifier: SimPL-2.0
pragma solidity ^0.8.0;

interface Bridge {
    function withdraw(uint toChainId, address toToken, address recipient, uint256 value, bytes memory hash) external;

    function withdrawNative(uint toChainId, address payable recipient, bool isMain, uint256 value, bytes memory hash) external;
}

contract BridgeManager {

    address public owner;

    address public bridgeAddress;

    mapping(address => bool) public isManager;

    mapping(bytes => bool) public isComplete;


    // 需要多签数量
    uint public signLimit;

    // multiSigns[networkID][txHash] = managers[]
    mapping(uint => mapping(bytes => address[])) public multiSigns;

    constructor(uint _signLimit, address _bridgeAddress) {
        owner = msg.sender;
        isManager[msg.sender] = true;
        signLimit = _signLimit;
        bridgeAddress = _bridgeAddress;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Bridge Manager: only use owner to call");
        _;
    }

    modifier onlyManager {
        require(isManager[msg.sender], "Bridge Manager: only manager can call this function");
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
        isManager[_address] = true;
    }

    function managerDel(address _address) public onlyOwner {
        isManager[_address] = false;
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

        if (!isComplete[hash] && multiSigns[toChainId][hash].length >= signLimit) {
            isComplete[hash] = true;
            return true;
        } else {
            return false;
        }
    }

    function withdraw(uint toChainId, address toToken, address recipient, uint256 value, bytes memory depositHash) public onlyManager {
        bool isMultiSign = multiSign(toChainId, depositHash);
        if (isMultiSign) {
            Bridge bridge = Bridge(bridgeAddress);
            bridge.withdraw(toChainId, toToken, recipient, value, depositHash);
        }

    }

    function withdrawNative(uint toChainId, address payable recipient, bool isMain, uint256 value, bytes memory hash) public onlyManager {
        bool isMultiSign = multiSign(toChainId, hash);
        if (isMultiSign) {
            Bridge bridge = Bridge(bridgeAddress);
            bridge.withdrawNative(toChainId, recipient, isMain, value, hash);
        }
    }

}