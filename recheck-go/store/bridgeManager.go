// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BridgeManagerMetaData contains all meta data concerning the BridgeManager contract.
var BridgeManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_signLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"txHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Confirmation\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Managers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"confirmations\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"isConfirmed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"managerAdd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"managerDel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"multiSigns\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeAddress\",\"type\":\"address\"}],\"name\":\"setBridgeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"setSignLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"txHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isNative\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isMain\",\"type\":\"bool\"}],\"name\":\"submitTransaction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transactions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"txHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isNative\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isMain\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200202c3803806200202c8339818101604052810190620000379190620000f6565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508260068190555081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050620001c9565b600081519050620000d98162000195565b92915050565b600081519050620000f081620001af565b92915050565b60008060006060848603121562000112576200011162000190565b5b60006200012286828701620000df565b93505060206200013586828701620000c8565b92505060406200014886828701620000c8565b9150509250925092565b60006200015f8262000166565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600080fd5b620001a08162000152565b8114620001ac57600080fd5b50565b620001ba8162000186565b8114620001c657600080fd5b50565b611e5380620001d96000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c8063a3c573eb11610097578063bce83aae11610066578063bce83aae146102c6578063c69ed5f2146102e2578063ca12e593146102fe578063f3ae24151461032e57610100565b8063a3c573eb1461022c578063a97472df1461024a578063b7dae8501461027a578063b9cc52311461029657610100565b8063642f2eaf116100d3578063642f2eaf1461018b5780636486aa51146101c25780637f5a22f9146101f25780638da5cb5b1461020e57610100565b80630c4ecab41461010557806313af4035146101355780631e6db6cb14610151578063538cf76b1461016f575b600080fd5b61011f600480360381019061011a91906114ee565b61035e565b60405161012c9190611804565b60405180910390f35b61014f600480360381019061014a9190611494565b61038d565b005b61015961045e565b604051610166919061183f565b60405180910390f35b61018960048036038101906101849190611467565b610464565b005b6101a560048036038101906101a091906114c1565b610642565b6040516101b998979695949392919061190e565b60405180910390f35b6101dc60048036038101906101d791906114c1565b610779565b6040516101e99190611804565b60405180910390f35b61020c60048036038101906102079190611467565b610875565b005b610216610947565b60405161022391906117e9565b60405180910390f35b61023461096b565b60405161024191906117e9565b60405180910390f35b610264600480360381019061025f9190611619565b610991565b60405161027191906117e9565b60405180910390f35b610294600480360381019061028f9190611467565b610a02565b005b6102b060048036038101906102ab919061152e565b610b4e565b6040516102bd91906117e9565b60405180910390f35b6102e060048036038101906102db919061152e565b610b8d565b005b6102fc60048036038101906102f791906114c1565b610c25565b005b6103186004803603810190610313919061155b565b610e6f565b6040516103259190611804565b60405180910390f35b61034860048036038101906103439190611467565b611187565b6040516103559190611804565b60405180910390f35b60046020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461041b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104129061181f565b60405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60065481565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104e99061181f565b60405180910390fd5b60005b6002805490508110156105e6578173ffffffffffffffffffffffffffffffffffffffff166002828154811061052d5761052c611ccb565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156105d35760006002828154811061058a57610589611ccb565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b80806105de90611bf6565b9150506104f5565b506000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b600560205280600052604060002060009150905080600001549080600101805461066b90611b93565b80601f016020809104026020016040519081016040528092919081815260200182805461069790611b93565b80156106e45780601f106106b9576101008083540402835291602001916106e4565b820191906000526020600020905b8154815290600101906020018083116106c757829003601f168201915b5050505050908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154908060050160009054906101000a900460ff16908060050160019054906101000a900460ff16908060050160029054906101000a900460ff16905088565b6000806000905060005b60028054905081101561086957600460008581526020019081526020016000206000600283815481106107b9576107b8611ccb565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156108415760018261083e9190611a97565b91505b60065482141561085657600192505050610870565b808061086190611bf6565b915050610783565b5060009150505b919050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610903576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108fa9061181f565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60076020528260005260406000208280516020810182018051848252602083016020850120818352809550505050505081815481106109cf57600080fd5b90600052602060002001600092509250509054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a90576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a879061181f565b60405180910390fd5b6002819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b60028181548110610b5e57600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c1b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c129061181f565b60405180910390fd5b8060068190555050565b60006005600083815260200190815260200160002090506000610c4783610779565b9050808015610c6b5750600015158260050160029054906101000a900460ff161515145b15610e6a5760018260050160026101000a81548160ff0219169083151502179055508160050160009054906101000a900460ff1615610d7f57610d7a82600001548360050160019054906101000a900460ff168460030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168560040154866001018054610cf790611b93565b80601f0160208091040260200160405190810160405280929190818152602001828054610d2390611b93565b8015610d705780601f10610d4557610100808354040283529160200191610d70565b820191906000526020600020905b815481529060010190602001808311610d5357829003601f168201915b50505050506111a7565b610e69565b610e6882600001548360020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168460030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168560040154866001018054610de590611b93565b80601f0160208091040260200160405190810160405280929190818152602001828054610e1190611b93565b8015610e5e5780601f10610e3357610100808354040283529160200191610e5e565b820191906000526020600020905b815481529060010190602001808311610e4157829003601f168201915b5050505050611249565b5b5b505050565b6000808888888888604051602001610e8b95949392919061178e565b6040516020818303038152906040528051906020012090506004600082815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610f1057600191505061117c565b6040518061010001604052808a81526020018981526020018873ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff168152602001868152602001851515815260200184151581526020016000151581525060056000838152602001908152602001600020600082015181600001556020820151816001019080519060200190610fb69291906112eb565b5060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506080820151816004015560a08201518160050160006101000a81548160ff02191690831515021790555060c08201518160050160016101000a81548160ff02191690831515021790555060e08201518160050160026101000a81548160ff02191690831515021790555090505060016004600083815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507fa7b5dd6ee28af9c8cbb153395344e76caa21d51905900240357233bbaecfe29c898989898986336040516111569796959493929190611993565b60405180910390a161116781610779565b156111765761117581610c25565b5b60019150505b979650505050505050565b60036020528060005260406000206000915054906101000a900460ff1681565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663f405250f87868887876040518663ffffffff1660e01b815260040161120f95949392919061185a565b600060405180830381600087803b15801561122957600080fd5b505af115801561123d573d6000803e3d6000fd5b50505050505050505050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663ae57de0c87878787876040518663ffffffff1660e01b81526004016112b19594939291906118b4565b600060405180830381600087803b1580156112cb57600080fd5b505af11580156112df573d6000803e3d6000fd5b50505050505050505050565b8280546112f790611b93565b90600052602060002090601f0160209004810192826113195760008555611360565b82601f1061133257805160ff1916838001178555611360565b82800160010185558215611360579182015b8281111561135f578251825591602001919060010190611344565b5b50905061136d9190611371565b5090565b5b8082111561138a576000816000905550600101611372565b5090565b60006113a161139c84611a2e565b611a09565b9050828152602081018484840111156113bd576113bc611d2e565b5b6113c8848285611b51565b509392505050565b6000813590506113df81611daa565b92915050565b6000813590506113f481611dc1565b92915050565b60008135905061140981611dd8565b92915050565b60008135905061141e81611def565b92915050565b600082601f83011261143957611438611d29565b5b813561144984826020860161138e565b91505092915050565b60008135905061146181611e06565b92915050565b60006020828403121561147d5761147c611d38565b5b600061148b848285016113d0565b91505092915050565b6000602082840312156114aa576114a9611d38565b5b60006114b8848285016113e5565b91505092915050565b6000602082840312156114d7576114d6611d38565b5b60006114e58482850161140f565b91505092915050565b6000806040838503121561150557611504611d38565b5b60006115138582860161140f565b9250506020611524858286016113d0565b9150509250929050565b60006020828403121561154457611543611d38565b5b600061155284828501611452565b91505092915050565b600080600080600080600060e0888a03121561157a57611579611d38565b5b60006115888a828b01611452565b975050602088013567ffffffffffffffff8111156115a9576115a8611d33565b5b6115b58a828b01611424565b96505060406115c68a828b016113d0565b95505060606115d78a828b016113d0565b94505060806115e88a828b01611452565b93505060a06115f98a828b016113fa565b92505060c061160a8a828b016113fa565b91505092959891949750929550565b60008060006060848603121561163257611631611d38565b5b600061164086828701611452565b935050602084013567ffffffffffffffff81111561166157611660611d33565b5b61166d86828701611424565b925050604061167e86828701611452565b9150509250925092565b61169181611aff565b82525050565b6116a081611aed565b82525050565b6116b76116b282611aed565b611c3f565b82525050565b6116c681611b11565b82525050565b6116d581611b1d565b82525050565b60006116e682611a5f565b6116f08185611a6a565b9350611700818560208601611b60565b61170981611d3d565b840191505092915050565b600061171f82611a5f565b6117298185611a7b565b9350611739818560208601611b60565b80840191505092915050565b6000611752602683611a86565b915061175d82611d5b565b604082019050919050565b61177181611b47565b82525050565b61178861178382611b47565b611c63565b82525050565b600061179a8288611777565b6020820191506117aa8287611714565b91506117b682866116a6565b6014820191506117c682856116a6565b6014820191506117d68284611777565b6020820191508190509695505050505050565b60006020820190506117fe6000830184611697565b92915050565b600060208201905061181960008301846116bd565b92915050565b6000602082019050818103600083015261183881611745565b9050919050565b60006020820190506118546000830184611768565b92915050565b600060a08201905061186f6000830188611768565b61187c6020830187611688565b61188960408301866116bd565b6118966060830185611768565b81810360808301526118a881846116db565b90509695505050505050565b600060a0820190506118c96000830188611768565b6118d66020830187611697565b6118e36040830186611697565b6118f06060830185611768565b818103608083015261190281846116db565b90509695505050505050565b600061010082019050611924600083018b611768565b8181036020830152611936818a6116db565b90506119456040830189611697565b6119526060830188611697565b61195f6080830187611768565b61196c60a08301866116bd565b61197960c08301856116bd565b61198660e08301846116bd565b9998505050505050505050565b600060e0820190506119a8600083018a611768565b81810360208301526119ba81896116db565b90506119c96040830188611697565b6119d66060830187611697565b6119e36080830186611768565b6119f060a08301856116cc565b6119fd60c0830184611697565b98975050505050505050565b6000611a13611a24565b9050611a1f8282611bc5565b919050565b6000604051905090565b600067ffffffffffffffff821115611a4957611a48611cfa565b5b611a5282611d3d565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000611aa282611b47565b9150611aad83611b47565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611ae257611ae1611c6d565b5b828201905092915050565b6000611af882611b27565b9050919050565b6000611b0a82611b27565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015611b7e578082015181840152602081019050611b63565b83811115611b8d576000848401525b50505050565b60006002820490506001821680611bab57607f821691505b60208210811415611bbf57611bbe611c9c565b5b50919050565b611bce82611d3d565b810181811067ffffffffffffffff82111715611bed57611bec611cfa565b5b80604052505050565b6000611c0182611b47565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611c3457611c33611c6d565b5b600182019050919050565b6000611c4a82611c51565b9050919050565b6000611c5c82611d4e565b9050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b7f427269646765204d616e616765723a206f6e6c7920757365206f776e6572207460008201527f6f2063616c6c0000000000000000000000000000000000000000000000000000602082015250565b611db381611aed565b8114611dbe57600080fd5b50565b611dca81611aff565b8114611dd557600080fd5b50565b611de181611b11565b8114611dec57600080fd5b50565b611df881611b1d565b8114611e0357600080fd5b50565b611e0f81611b47565b8114611e1a57600080fd5b5056fea2646970667358221220f5b56747975035a24806fc701f240f89e9492966643c9dfa7648af8280d375f164736f6c63430008060033",
}

// BridgeManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeManagerMetaData.ABI instead.
var BridgeManagerABI = BridgeManagerMetaData.ABI

// BridgeManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeManagerMetaData.Bin instead.
var BridgeManagerBin = BridgeManagerMetaData.Bin

// DeployBridgeManager deploys a new Ethereum contract, binding an instance of BridgeManager to it.
func DeployBridgeManager(auth *bind.TransactOpts, backend bind.ContractBackend, _signLimit *big.Int, _bridgeAddress common.Address, _owner common.Address) (common.Address, *types.Transaction, *BridgeManager, error) {
	parsed, err := BridgeManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeManagerBin), backend, _signLimit, _bridgeAddress, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeManager{BridgeManagerCaller: BridgeManagerCaller{contract: contract}, BridgeManagerTransactor: BridgeManagerTransactor{contract: contract}, BridgeManagerFilterer: BridgeManagerFilterer{contract: contract}}, nil
}

// BridgeManager is an auto generated Go binding around an Ethereum contract.
type BridgeManager struct {
	BridgeManagerCaller     // Read-only binding to the contract
	BridgeManagerTransactor // Write-only binding to the contract
	BridgeManagerFilterer   // Log filterer for contract events
}

// BridgeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeManagerSession struct {
	Contract     *BridgeManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeManagerCallerSession struct {
	Contract *BridgeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BridgeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeManagerTransactorSession struct {
	Contract     *BridgeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BridgeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeManagerRaw struct {
	Contract *BridgeManager // Generic contract binding to access the raw methods on
}

// BridgeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeManagerCallerRaw struct {
	Contract *BridgeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeManagerTransactorRaw struct {
	Contract *BridgeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeManager creates a new instance of BridgeManager, bound to a specific deployed contract.
func NewBridgeManager(address common.Address, backend bind.ContractBackend) (*BridgeManager, error) {
	contract, err := bindBridgeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeManager{BridgeManagerCaller: BridgeManagerCaller{contract: contract}, BridgeManagerTransactor: BridgeManagerTransactor{contract: contract}, BridgeManagerFilterer: BridgeManagerFilterer{contract: contract}}, nil
}

// NewBridgeManagerCaller creates a new read-only instance of BridgeManager, bound to a specific deployed contract.
func NewBridgeManagerCaller(address common.Address, caller bind.ContractCaller) (*BridgeManagerCaller, error) {
	contract, err := bindBridgeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeManagerCaller{contract: contract}, nil
}

// NewBridgeManagerTransactor creates a new write-only instance of BridgeManager, bound to a specific deployed contract.
func NewBridgeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeManagerTransactor, error) {
	contract, err := bindBridgeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeManagerTransactor{contract: contract}, nil
}

// NewBridgeManagerFilterer creates a new log filterer instance of BridgeManager, bound to a specific deployed contract.
func NewBridgeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeManagerFilterer, error) {
	contract, err := bindBridgeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeManagerFilterer{contract: contract}, nil
}

// bindBridgeManager binds a generic wrapper to an already deployed contract.
func bindBridgeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeManager *BridgeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeManager.Contract.BridgeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeManager *BridgeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeManager.Contract.BridgeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeManager *BridgeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeManager.Contract.BridgeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeManager *BridgeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeManager *BridgeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeManager *BridgeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeManager.Contract.contract.Transact(opts, method, params...)
}

// Managers is a free data retrieval call binding the contract method 0xb9cc5231.
//
// Solidity: function Managers(uint256 ) view returns(address)
func (_BridgeManager *BridgeManagerCaller) Managers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "Managers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Managers is a free data retrieval call binding the contract method 0xb9cc5231.
//
// Solidity: function Managers(uint256 ) view returns(address)
func (_BridgeManager *BridgeManagerSession) Managers(arg0 *big.Int) (common.Address, error) {
	return _BridgeManager.Contract.Managers(&_BridgeManager.CallOpts, arg0)
}

// Managers is a free data retrieval call binding the contract method 0xb9cc5231.
//
// Solidity: function Managers(uint256 ) view returns(address)
func (_BridgeManager *BridgeManagerCallerSession) Managers(arg0 *big.Int) (common.Address, error) {
	return _BridgeManager.Contract.Managers(&_BridgeManager.CallOpts, arg0)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_BridgeManager *BridgeManagerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_BridgeManager *BridgeManagerSession) BridgeAddress() (common.Address, error) {
	return _BridgeManager.Contract.BridgeAddress(&_BridgeManager.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_BridgeManager *BridgeManagerCallerSession) BridgeAddress() (common.Address, error) {
	return _BridgeManager.Contract.BridgeAddress(&_BridgeManager.CallOpts)
}

// Confirmations is a free data retrieval call binding the contract method 0x0c4ecab4.
//
// Solidity: function confirmations(bytes32 , address ) view returns(bool)
func (_BridgeManager *BridgeManagerCaller) Confirmations(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "confirmations", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Confirmations is a free data retrieval call binding the contract method 0x0c4ecab4.
//
// Solidity: function confirmations(bytes32 , address ) view returns(bool)
func (_BridgeManager *BridgeManagerSession) Confirmations(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _BridgeManager.Contract.Confirmations(&_BridgeManager.CallOpts, arg0, arg1)
}

// Confirmations is a free data retrieval call binding the contract method 0x0c4ecab4.
//
// Solidity: function confirmations(bytes32 , address ) view returns(bool)
func (_BridgeManager *BridgeManagerCallerSession) Confirmations(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _BridgeManager.Contract.Confirmations(&_BridgeManager.CallOpts, arg0, arg1)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x6486aa51.
//
// Solidity: function isConfirmed(bytes32 transactionId) view returns(bool)
func (_BridgeManager *BridgeManagerCaller) IsConfirmed(opts *bind.CallOpts, transactionId [32]byte) (bool, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "isConfirmed", transactionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsConfirmed is a free data retrieval call binding the contract method 0x6486aa51.
//
// Solidity: function isConfirmed(bytes32 transactionId) view returns(bool)
func (_BridgeManager *BridgeManagerSession) IsConfirmed(transactionId [32]byte) (bool, error) {
	return _BridgeManager.Contract.IsConfirmed(&_BridgeManager.CallOpts, transactionId)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x6486aa51.
//
// Solidity: function isConfirmed(bytes32 transactionId) view returns(bool)
func (_BridgeManager *BridgeManagerCallerSession) IsConfirmed(transactionId [32]byte) (bool, error) {
	return _BridgeManager.Contract.IsConfirmed(&_BridgeManager.CallOpts, transactionId)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address ) view returns(bool)
func (_BridgeManager *BridgeManagerCaller) IsManager(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "isManager", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address ) view returns(bool)
func (_BridgeManager *BridgeManagerSession) IsManager(arg0 common.Address) (bool, error) {
	return _BridgeManager.Contract.IsManager(&_BridgeManager.CallOpts, arg0)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address ) view returns(bool)
func (_BridgeManager *BridgeManagerCallerSession) IsManager(arg0 common.Address) (bool, error) {
	return _BridgeManager.Contract.IsManager(&_BridgeManager.CallOpts, arg0)
}

// MultiSigns is a free data retrieval call binding the contract method 0xa97472df.
//
// Solidity: function multiSigns(uint256 , bytes , uint256 ) view returns(address)
func (_BridgeManager *BridgeManagerCaller) MultiSigns(opts *bind.CallOpts, arg0 *big.Int, arg1 []byte, arg2 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "multiSigns", arg0, arg1, arg2)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MultiSigns is a free data retrieval call binding the contract method 0xa97472df.
//
// Solidity: function multiSigns(uint256 , bytes , uint256 ) view returns(address)
func (_BridgeManager *BridgeManagerSession) MultiSigns(arg0 *big.Int, arg1 []byte, arg2 *big.Int) (common.Address, error) {
	return _BridgeManager.Contract.MultiSigns(&_BridgeManager.CallOpts, arg0, arg1, arg2)
}

// MultiSigns is a free data retrieval call binding the contract method 0xa97472df.
//
// Solidity: function multiSigns(uint256 , bytes , uint256 ) view returns(address)
func (_BridgeManager *BridgeManagerCallerSession) MultiSigns(arg0 *big.Int, arg1 []byte, arg2 *big.Int) (common.Address, error) {
	return _BridgeManager.Contract.MultiSigns(&_BridgeManager.CallOpts, arg0, arg1, arg2)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeManager *BridgeManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeManager *BridgeManagerSession) Owner() (common.Address, error) {
	return _BridgeManager.Contract.Owner(&_BridgeManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeManager *BridgeManagerCallerSession) Owner() (common.Address, error) {
	return _BridgeManager.Contract.Owner(&_BridgeManager.CallOpts)
}

// SignLimit is a free data retrieval call binding the contract method 0x1e6db6cb.
//
// Solidity: function signLimit() view returns(uint256)
func (_BridgeManager *BridgeManagerCaller) SignLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "signLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SignLimit is a free data retrieval call binding the contract method 0x1e6db6cb.
//
// Solidity: function signLimit() view returns(uint256)
func (_BridgeManager *BridgeManagerSession) SignLimit() (*big.Int, error) {
	return _BridgeManager.Contract.SignLimit(&_BridgeManager.CallOpts)
}

// SignLimit is a free data retrieval call binding the contract method 0x1e6db6cb.
//
// Solidity: function signLimit() view returns(uint256)
func (_BridgeManager *BridgeManagerCallerSession) SignLimit() (*big.Int, error) {
	return _BridgeManager.Contract.SignLimit(&_BridgeManager.CallOpts)
}

// Transactions is a free data retrieval call binding the contract method 0x642f2eaf.
//
// Solidity: function transactions(bytes32 ) view returns(uint256 fromChainId, bytes txHash, address toToken, address recipient, uint256 amount, bool isNative, bool isMain, bool executed)
func (_BridgeManager *BridgeManagerCaller) Transactions(opts *bind.CallOpts, arg0 [32]byte) (struct {
	FromChainId *big.Int
	TxHash      []byte
	ToToken     common.Address
	Recipient   common.Address
	Amount      *big.Int
	IsNative    bool
	IsMain      bool
	Executed    bool
}, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "transactions", arg0)

	outstruct := new(struct {
		FromChainId *big.Int
		TxHash      []byte
		ToToken     common.Address
		Recipient   common.Address
		Amount      *big.Int
		IsNative    bool
		IsMain      bool
		Executed    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FromChainId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TxHash = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.ToToken = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Recipient = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.IsNative = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.IsMain = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.Executed = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// Transactions is a free data retrieval call binding the contract method 0x642f2eaf.
//
// Solidity: function transactions(bytes32 ) view returns(uint256 fromChainId, bytes txHash, address toToken, address recipient, uint256 amount, bool isNative, bool isMain, bool executed)
func (_BridgeManager *BridgeManagerSession) Transactions(arg0 [32]byte) (struct {
	FromChainId *big.Int
	TxHash      []byte
	ToToken     common.Address
	Recipient   common.Address
	Amount      *big.Int
	IsNative    bool
	IsMain      bool
	Executed    bool
}, error) {
	return _BridgeManager.Contract.Transactions(&_BridgeManager.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x642f2eaf.
//
// Solidity: function transactions(bytes32 ) view returns(uint256 fromChainId, bytes txHash, address toToken, address recipient, uint256 amount, bool isNative, bool isMain, bool executed)
func (_BridgeManager *BridgeManagerCallerSession) Transactions(arg0 [32]byte) (struct {
	FromChainId *big.Int
	TxHash      []byte
	ToToken     common.Address
	Recipient   common.Address
	Amount      *big.Int
	IsNative    bool
	IsMain      bool
	Executed    bool
}, error) {
	return _BridgeManager.Contract.Transactions(&_BridgeManager.CallOpts, arg0)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xc69ed5f2.
//
// Solidity: function executeTransaction(bytes32 transactionId) returns()
func (_BridgeManager *BridgeManagerTransactor) ExecuteTransaction(opts *bind.TransactOpts, transactionId [32]byte) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "executeTransaction", transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xc69ed5f2.
//
// Solidity: function executeTransaction(bytes32 transactionId) returns()
func (_BridgeManager *BridgeManagerSession) ExecuteTransaction(transactionId [32]byte) (*types.Transaction, error) {
	return _BridgeManager.Contract.ExecuteTransaction(&_BridgeManager.TransactOpts, transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xc69ed5f2.
//
// Solidity: function executeTransaction(bytes32 transactionId) returns()
func (_BridgeManager *BridgeManagerTransactorSession) ExecuteTransaction(transactionId [32]byte) (*types.Transaction, error) {
	return _BridgeManager.Contract.ExecuteTransaction(&_BridgeManager.TransactOpts, transactionId)
}

// ManagerAdd is a paid mutator transaction binding the contract method 0xb7dae850.
//
// Solidity: function managerAdd(address _address) returns()
func (_BridgeManager *BridgeManagerTransactor) ManagerAdd(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "managerAdd", _address)
}

// ManagerAdd is a paid mutator transaction binding the contract method 0xb7dae850.
//
// Solidity: function managerAdd(address _address) returns()
func (_BridgeManager *BridgeManagerSession) ManagerAdd(_address common.Address) (*types.Transaction, error) {
	return _BridgeManager.Contract.ManagerAdd(&_BridgeManager.TransactOpts, _address)
}

// ManagerAdd is a paid mutator transaction binding the contract method 0xb7dae850.
//
// Solidity: function managerAdd(address _address) returns()
func (_BridgeManager *BridgeManagerTransactorSession) ManagerAdd(_address common.Address) (*types.Transaction, error) {
	return _BridgeManager.Contract.ManagerAdd(&_BridgeManager.TransactOpts, _address)
}

// ManagerDel is a paid mutator transaction binding the contract method 0x538cf76b.
//
// Solidity: function managerDel(address _address) returns()
func (_BridgeManager *BridgeManagerTransactor) ManagerDel(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "managerDel", _address)
}

// ManagerDel is a paid mutator transaction binding the contract method 0x538cf76b.
//
// Solidity: function managerDel(address _address) returns()
func (_BridgeManager *BridgeManagerSession) ManagerDel(_address common.Address) (*types.Transaction, error) {
	return _BridgeManager.Contract.ManagerDel(&_BridgeManager.TransactOpts, _address)
}

// ManagerDel is a paid mutator transaction binding the contract method 0x538cf76b.
//
// Solidity: function managerDel(address _address) returns()
func (_BridgeManager *BridgeManagerTransactorSession) ManagerDel(_address common.Address) (*types.Transaction, error) {
	return _BridgeManager.Contract.ManagerDel(&_BridgeManager.TransactOpts, _address)
}

// SetBridgeAddress is a paid mutator transaction binding the contract method 0x7f5a22f9.
//
// Solidity: function setBridgeAddress(address _bridgeAddress) returns()
func (_BridgeManager *BridgeManagerTransactor) SetBridgeAddress(opts *bind.TransactOpts, _bridgeAddress common.Address) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "setBridgeAddress", _bridgeAddress)
}

// SetBridgeAddress is a paid mutator transaction binding the contract method 0x7f5a22f9.
//
// Solidity: function setBridgeAddress(address _bridgeAddress) returns()
func (_BridgeManager *BridgeManagerSession) SetBridgeAddress(_bridgeAddress common.Address) (*types.Transaction, error) {
	return _BridgeManager.Contract.SetBridgeAddress(&_BridgeManager.TransactOpts, _bridgeAddress)
}

// SetBridgeAddress is a paid mutator transaction binding the contract method 0x7f5a22f9.
//
// Solidity: function setBridgeAddress(address _bridgeAddress) returns()
func (_BridgeManager *BridgeManagerTransactorSession) SetBridgeAddress(_bridgeAddress common.Address) (*types.Transaction, error) {
	return _BridgeManager.Contract.SetBridgeAddress(&_BridgeManager.TransactOpts, _bridgeAddress)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_BridgeManager *BridgeManagerTransactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_BridgeManager *BridgeManagerSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _BridgeManager.Contract.SetOwner(&_BridgeManager.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_BridgeManager *BridgeManagerTransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _BridgeManager.Contract.SetOwner(&_BridgeManager.TransactOpts, _owner)
}

// SetSignLimit is a paid mutator transaction binding the contract method 0xbce83aae.
//
// Solidity: function setSignLimit(uint256 num) returns()
func (_BridgeManager *BridgeManagerTransactor) SetSignLimit(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "setSignLimit", num)
}

// SetSignLimit is a paid mutator transaction binding the contract method 0xbce83aae.
//
// Solidity: function setSignLimit(uint256 num) returns()
func (_BridgeManager *BridgeManagerSession) SetSignLimit(num *big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.SetSignLimit(&_BridgeManager.TransactOpts, num)
}

// SetSignLimit is a paid mutator transaction binding the contract method 0xbce83aae.
//
// Solidity: function setSignLimit(uint256 num) returns()
func (_BridgeManager *BridgeManagerTransactorSession) SetSignLimit(num *big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.SetSignLimit(&_BridgeManager.TransactOpts, num)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xca12e593.
//
// Solidity: function submitTransaction(uint256 fromChainId, bytes txHash, address toToken, address recipient, uint256 amount, bool isNative, bool isMain) returns(bool)
func (_BridgeManager *BridgeManagerTransactor) SubmitTransaction(opts *bind.TransactOpts, fromChainId *big.Int, txHash []byte, toToken common.Address, recipient common.Address, amount *big.Int, isNative bool, isMain bool) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "submitTransaction", fromChainId, txHash, toToken, recipient, amount, isNative, isMain)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xca12e593.
//
// Solidity: function submitTransaction(uint256 fromChainId, bytes txHash, address toToken, address recipient, uint256 amount, bool isNative, bool isMain) returns(bool)
func (_BridgeManager *BridgeManagerSession) SubmitTransaction(fromChainId *big.Int, txHash []byte, toToken common.Address, recipient common.Address, amount *big.Int, isNative bool, isMain bool) (*types.Transaction, error) {
	return _BridgeManager.Contract.SubmitTransaction(&_BridgeManager.TransactOpts, fromChainId, txHash, toToken, recipient, amount, isNative, isMain)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xca12e593.
//
// Solidity: function submitTransaction(uint256 fromChainId, bytes txHash, address toToken, address recipient, uint256 amount, bool isNative, bool isMain) returns(bool)
func (_BridgeManager *BridgeManagerTransactorSession) SubmitTransaction(fromChainId *big.Int, txHash []byte, toToken common.Address, recipient common.Address, amount *big.Int, isNative bool, isMain bool) (*types.Transaction, error) {
	return _BridgeManager.Contract.SubmitTransaction(&_BridgeManager.TransactOpts, fromChainId, txHash, toToken, recipient, amount, isNative, isMain)
}

// BridgeManagerConfirmationIterator is returned from FilterConfirmation and is used to iterate over the raw logs and unpacked data for Confirmation events raised by the BridgeManager contract.
type BridgeManagerConfirmationIterator struct {
	Event *BridgeManagerConfirmation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeManagerConfirmationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeManagerConfirmation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeManagerConfirmation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeManagerConfirmationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeManagerConfirmationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeManagerConfirmation represents a Confirmation event raised by the BridgeManager contract.
type BridgeManagerConfirmation struct {
	FromChainId   *big.Int
	TxHash        []byte
	ToToken       common.Address
	Recipient     common.Address
	Amount        *big.Int
	TransactionId [32]byte
	Sender        common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterConfirmation is a free log retrieval operation binding the contract event 0xa7b5dd6ee28af9c8cbb153395344e76caa21d51905900240357233bbaecfe29c.
//
// Solidity: event Confirmation(uint256 fromChainId, bytes txHash, address toToken, address recipient, uint256 amount, bytes32 transactionId, address sender)
func (_BridgeManager *BridgeManagerFilterer) FilterConfirmation(opts *bind.FilterOpts) (*BridgeManagerConfirmationIterator, error) {

	logs, sub, err := _BridgeManager.contract.FilterLogs(opts, "Confirmation")
	if err != nil {
		return nil, err
	}
	return &BridgeManagerConfirmationIterator{contract: _BridgeManager.contract, event: "Confirmation", logs: logs, sub: sub}, nil
}

// WatchConfirmation is a free log subscription operation binding the contract event 0xa7b5dd6ee28af9c8cbb153395344e76caa21d51905900240357233bbaecfe29c.
//
// Solidity: event Confirmation(uint256 fromChainId, bytes txHash, address toToken, address recipient, uint256 amount, bytes32 transactionId, address sender)
func (_BridgeManager *BridgeManagerFilterer) WatchConfirmation(opts *bind.WatchOpts, sink chan<- *BridgeManagerConfirmation) (event.Subscription, error) {

	logs, sub, err := _BridgeManager.contract.WatchLogs(opts, "Confirmation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeManagerConfirmation)
				if err := _BridgeManager.contract.UnpackLog(event, "Confirmation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConfirmation is a log parse operation binding the contract event 0xa7b5dd6ee28af9c8cbb153395344e76caa21d51905900240357233bbaecfe29c.
//
// Solidity: event Confirmation(uint256 fromChainId, bytes txHash, address toToken, address recipient, uint256 amount, bytes32 transactionId, address sender)
func (_BridgeManager *BridgeManagerFilterer) ParseConfirmation(log types.Log) (*BridgeManagerConfirmation, error) {
	event := new(BridgeManagerConfirmation)
	if err := _BridgeManager.contract.UnpackLog(event, "Confirmation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
