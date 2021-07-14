<?php

namespace App\Admin\Extensions;

use App\Models\Chain;
use App\Models\Pair;
use Dcat\Admin\Grid\RowAction;

class InsertPair extends RowAction
{

    /**
     * 返回字段标题
     *
     * @return string
     */
    public function title()
    {
        return '添加跨链对(智能合约)';
    }


    /**
     * 添加JS
     *
     * @return string
     * https://learnku.com/docs/dcat-admin/2.x/js-component/8087
     */
    protected function script()
    {
        $pk_owner = json_encode(env("PK_BRIDGE_OWNER"));

        $chains = json_encode(
            Chain::all()->groupBy('chainId')->toArray()
        );

        return <<<JS
$('.grid-insert-pair').on('click',   function() {

    Dcat.loading()

    let pk_owner = $pk_owner
    let chains = $chains
    console.log(pk_owner)


    const fromChainId = $(this).data("from-chain")
    const toChainId = $(this).data("to-chain")
    const fromToken = $(this).data("from-token")
    const toToken = $(this).data("to-token")
    const isMain = $(this).data("is-main")
    const isNative = $(this).data("is-native")
    const abi = [
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "chainId",
				"type": "uint256"
			},
			{
				"internalType": "address",
				"name": "toToken",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "deposit",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "toChainId",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "fromToken",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "toToken",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "recipient",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "Deposit",
		"type": "event"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "toChainId",
				"type": "uint256"
			},
			{
				"internalType": "bool",
				"name": "isMain",
				"type": "bool"
			},
			{
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "depositNative",
		"outputs": [],
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "toChainId",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "bool",
				"name": "isMain",
				"type": "bool"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "recipient",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "DepositNative",
		"type": "event"
	},
	{
		"inputs": [
			{
				"internalType": "address payable",
				"name": "recipient",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "nativeTransfer",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address payable",
				"name": "newAdmin",
				"type": "address"
			}
		],
		"name": "setAdmin",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "toChainId",
				"type": "uint256"
			},
			{
				"internalType": "bool",
				"name": "isMain",
				"type": "bool"
			},
			{
				"internalType": "bool",
				"name": "state",
				"type": "bool"
			}
		],
		"name": "setNativeIsRun",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "toChainId",
				"type": "uint256"
			},
			{
				"internalType": "address",
				"name": "toToken",
				"type": "address"
			},
			{
				"internalType": "bool",
				"name": "state",
				"type": "bool"
			}
		],
		"name": "setTokenIsRun",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "fromToken",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "recipient",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "tokenTransfer",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "toChainId",
				"type": "uint256"
			},
			{
				"internalType": "address",
				"name": "toToken",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "recipient",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "withdraw",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "toChainId",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "fromToken",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "toToken",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "recipient",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "WithdrawDone",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "fromChainId",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "recipient",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "bool",
				"name": "isMain",
				"type": "bool"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "WithdrawNativeDone",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "address",
				"name": "_address",
				"type": "address"
			}
		],
		"name": "adminChanged",
		"type": "event"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "toChainId",
				"type": "uint256"
			},
			{
				"internalType": "bool",
				"name": "isRun",
				"type": "bool"
			},
			{
				"internalType": "bool",
				"name": "isMain",
				"type": "bool"
			},
			{
				"internalType": "address",
				"name": "fromToken",
				"type": "address"
			}
		],
		"name": "nativeInsert",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address payable",
				"name": "newOwner",
				"type": "address"
			}
		],
		"name": "setOwner",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "toChainId",
				"type": "uint256"
			},
			{
				"internalType": "address",
				"name": "toToken",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "fromToken",
				"type": "address"
			},
			{
				"internalType": "bool",
				"name": "isRun",
				"type": "bool"
			},
			{
				"internalType": "bool",
				"name": "isMain",
				"type": "bool"
			}
		],
		"name": "tokenInsert",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "fromChainId",
				"type": "uint256"
			},
			{
				"internalType": "address payable",
				"name": "recipient",
				"type": "address"
			},
			{
				"internalType": "bool",
				"name": "isMain",
				"type": "bool"
			},
			{
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "withdrawNative",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"stateMutability": "payable",
		"type": "receive"
	},
	{
		"inputs": [],
		"name": "admin",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			},
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"name": "natives",
		"outputs": [
			{
				"internalType": "bool",
				"name": "isRun",
				"type": "bool"
			},
			{
				"internalType": "bool",
				"name": "isMain",
				"type": "bool"
			},
			{
				"internalType": "address",
				"name": "local",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "owner",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			},
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"name": "tokens",
		"outputs": [
			{
				"internalType": "bool",
				"name": "isRun",
				"type": "bool"
			},
			{
				"internalType": "bool",
				"name": "isMain",
				"type": "bool"
			},
			{
				"internalType": "address",
				"name": "local",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	}
]
    try {
      const fromChain = chains[fromChainId][0]
        const fromProvider = new ethers.providers.JsonRpcProvider(fromChain.url)
        const fromWallet = new ethers.Wallet(pk_owner, fromProvider)
        const fromBridge = new ethers.Contract(fromChain['bridge'], abi, fromWallet)
        if (isNative){
            fromBridge['natives'](toChainId,isMain).then(data=>{
                if (data.isRun){
                    Dcat.loading(false)
                    Dcat.confirm('跨链对在' + fromChain.name + '链已经存在了 要覆盖吗', null,  ()=> {
                        fromBridge['nativeInsert'](toChainId,true,isMain,fromToken).then(tx=>{
                             Dcat.success('跨链对在' + fromChain.name + '链部署完成')
                        })
                    })
                }else{
                    fromBridge['nativeInsert'](toChainId,true,isMain,fromToken).then(tx=>{
                         Dcat.success('跨链对在' + fromChain.name + '链部署完成')
                         Dcat.loading(false)
                    })
                }
            })
        }else {
             fromBridge['tokens'](toChainId,toToken).then(data=>{
                 if (data.local !== "0x0000000000000000000000000000000000000000"){
                     Dcat.loading(false)
                     Dcat.confirm('跨链对在' + fromChain.name + '链已经存在了 要覆盖吗', null,  ()=> {
                        fromBridge['tokenInsert'](toChainId,toToken,fromToken,true,isMain).then(tx=>{
                            tx.wait().then(res=>{
                                if (res.confirmations){
                                    Dcat.success('跨链对在' + fromChain.name + '链部署完成')
                                }
                            })
                        })
                    });
                 }else{
                        fromBridge['tokenInsert'](toChainId,toToken,fromToken,true,isMain).then(tx=>{
                            tx.wait().then(res=>{
                                if (res.confirmations){
                                    Dcat.success('跨链对在' + fromChain.name + '链部署完成')
                                    Dcat.loading(false)
                                }
                            })
                        })
                 }
            })
        }
    }catch (e) {
        Dcat.error(e.message)
        Dcat.loading(false)
        Dcat.reload()
    }

});
JS;
    }

    public function html()
    {
        // 获取当前行数据ID
        $id = $this->getKey();

        // 这里需要添加一个class, 和上面script方法对应
        $this->setHtmlAttribute([
            'data-id' => $id,
            'data-from-chain' => $this->row->fromChain,
            'data-to-chain' => $this->row->toChain,
            'data-from-token' => $this->row->fromToken,
            'data-to-token' => $this->row->toToken,
            'data-is-main' => $this->row->isMain,
            'data-is-native' => $this->row->isNative,
            'class' => 'grid-insert-pair'
        ]);

        return parent::html();
    }
}
