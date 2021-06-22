<?php

namespace App\Admin\Extensions;

use App\Models\Chain;
use App\Models\Pair;
use Dcat\Admin\Grid\RowAction;

class CheckRow extends RowAction
{

    /**
     * 返回字段标题
     *
     * @return string
     */
    public function title()
    {
        return '审核';
    }


    /**
     * 添加JS
     *
     * @return string
     * https://learnku.com/docs/dcat-admin/2.x/js-component/8087
     */
    protected function script()
    {
        $pks = json_encode([
            '20181205' => env("PK_20181205"),
            '3' => env("PK_3")
        ]);

        $pairs = json_encode(
            Pair::all()->groupBy('id')->toArray()
        );

        $chains = json_encode(
            Chain::all()->groupBy('chainId')->toArray()
        );

        return <<<JS
$('.grid-check-row').on('click',   function() {

    const logId = $(this).data('id')
    const pairId = $(this).data('pair-id')
    const recipient = $(this).data('recipient')
    const hash = $(this).data('hash')
    const value = $(this).data('value')
    let pks = $pks
    let pairs = $pairs
    let chains = $chains
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

    if (hash){
        Dcat.error("已经跨链到账!")
        return false
    }

    try {
        Dcat.loading()

        const pair = pairs[pairId][0]
        const chain = chains[pair['toChain']][0]

        const provider = new ethers.providers.JsonRpcProvider(chain.url)
        const wallet = new ethers.Wallet(pks[pair['toChain']], provider)
        const bridge = new ethers.Contract(chain['bridge'], abi, wallet)

        const fee = Math.ceil(value * pair['bridgeFee'] / 100)
        const final = ethers.utils.parseUnits((value-fee).toString(), pair['decimal'])

        if (pair.isNative){
            // 主网币
            bridge['withdrawNative'](pair['fromChain'],recipient,!pair['isMain'],final).then(tx=>{
                tx.wait().then(res=>{
                    fetch("/admin/withdraw?logId="+logId+"&hash="+res.transactionHash).then()
                    Dcat.loading(false)
                    Dcat.success("审核成功跨链已到账")
                    Dcat.reload()
                })
            }).catch(e=>{
                Dcat.loading(false)
                console.log(e.message)
                Dcat.error(e.message)
            })
        }else{
            // 是代币
           bridge['withdraw'](pair['fromChain'],pair['fromToken'],recipient,final).then(tx=>{
                tx.wait().then(res=>{
                    fetch("/admin/withdraw?logId="+logId+"&hash="+res.transactionHash).then()
                    Dcat.loading(false)
                    Dcat.success("审核成功跨链已到账")
                    Dcat.reload()
                })
           })
        }
    }catch (e) {
        Dcat.loading(false)
        Dcat.error(e.message)
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
            'data-pair-id' => $this->row->pairId,
            'data-hash' => $this->row->withdrawHash,
            'data-recipient' => $this->row->recipient,
            'data-value' => $this->row->value,
            'class' => 'grid-check-row'
        ]);

        return parent::html();
    }
}
