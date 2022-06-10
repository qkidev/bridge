<?php

namespace App\Admin\Extensions;

use App\Models\Chain;
use App\Models\Pair;
use Dcat\Admin\Grid\RowAction;

class CheckWithdraw extends RowAction
{

    /**
     * 返回字段标题
     *
     * @return string
     */
    public function title()
    {
        return '跨链审核';
    }


    /**
     * 添加JS
     *
     * @return string
     * https://learnku.com/docs/dcat-admin/2.x/js-component/8087
     */
    protected function script()
    {

        $pk_manager = json_encode(env("PK_BRIDGE_MANAGER"));

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
    const depositHash = $(this).data('deposit-hash')
    let pk_manager = $pk_manager
    let pairs = $pairs
    let chains = $chains
    const abiManager = [
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "_signLimit",
                "type": "uint256"
            },
            {
                "internalType": "address",
                "name": "_bridgeAddress",
                "type": "address"
            },
            {
                "internalType": "address",
                "name": "_owner",
                "type": "address"
            }
        ],
        "stateMutability": "nonpayable",
        "type": "constructor"
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
                "internalType": "bytes",
                "name": "txHash",
                "type": "bytes"
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
                "name": "amount",
                "type": "uint256"
            },
            {
                "indexed": false,
                "internalType": "bytes32",
                "name": "transactionId",
                "type": "bytes32"
            },
            {
                "indexed": false,
                "internalType": "address",
                "name": "sender",
                "type": "address"
            }
        ],
        "name": "Confirmation",
        "type": "event"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "name": "Managers",
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
        "inputs": [],
        "name": "bridgeAddress",
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
                "internalType": "bytes32",
                "name": "",
                "type": "bytes32"
            },
            {
                "internalType": "address",
                "name": "",
                "type": "address"
            }
        ],
        "name": "confirmations",
        "outputs": [
            {
                "internalType": "bool",
                "name": "",
                "type": "bool"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "bytes32",
                "name": "transactionId",
                "type": "bytes32"
            }
        ],
        "name": "executeTransaction",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "bytes32",
                "name": "transactionId",
                "type": "bytes32"
            }
        ],
        "name": "isConfirmed",
        "outputs": [
            {
                "internalType": "bool",
                "name": "",
                "type": "bool"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "",
                "type": "address"
            }
        ],
        "name": "isManager",
        "outputs": [
            {
                "internalType": "bool",
                "name": "",
                "type": "bool"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "_address",
                "type": "address"
            }
        ],
        "name": "managerAdd",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "_address",
                "type": "address"
            }
        ],
        "name": "managerDel",
        "outputs": [],
        "stateMutability": "nonpayable",
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
                "internalType": "bytes",
                "name": "",
                "type": "bytes"
            },
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "name": "multiSigns",
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
                "internalType": "address",
                "name": "_bridgeAddress",
                "type": "address"
            }
        ],
        "name": "setBridgeAddress",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address payable",
                "name": "_owner",
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
                "name": "num",
                "type": "uint256"
            }
        ],
        "name": "setSignLimit",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "signLimit",
        "outputs": [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "stateMutability": "view",
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
                "internalType": "bytes",
                "name": "txHash",
                "type": "bytes"
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
                "name": "amount",
                "type": "uint256"
            },
            {
                "internalType": "bool",
                "name": "isNative",
                "type": "bool"
            },
            {
                "internalType": "bool",
                "name": "isMain",
                "type": "bool"
            }
        ],
        "name": "submitTransaction",
        "outputs": [
            {
                "internalType": "bool",
                "name": "",
                "type": "bool"
            }
        ],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "bytes32",
                "name": "",
                "type": "bytes32"
            }
        ],
        "name": "transactions",
        "outputs": [
            {
                "internalType": "uint256",
                "name": "fromChainId",
                "type": "uint256"
            },
            {
                "internalType": "bytes",
                "name": "txHash",
                "type": "bytes"
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
                "name": "amount",
                "type": "uint256"
            },
            {
                "internalType": "bool",
                "name": "isNative",
                "type": "bool"
            },
            {
                "internalType": "bool",
                "name": "isMain",
                "type": "bool"
            },
            {
                "internalType": "bool",
                "name": "executed",
                "type": "bool"
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
        const wallet = new ethers.Wallet(pk_manager, provider)
        const bridgeManager = new ethers.Contract(chain['bridge_manager'], abiManager, wallet)

        const fee = (value * pair['bridgeFee'] / 100).toFixed(pair['decimal'])
        const final = ethers.utils.parseUnits((value-fee).toString(), pair['decimal'])

        if (pair.isNative){
            // 主网币
            bridgeManager['submitTransaction'](pair['fromChain'],depositHash,"0x0000000000000000000000000000000000000000",recipient,final,true,!pair['isMain']).then(tx=>{
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
            // is token
           bridgeManager['submitTransaction'](pair['fromChain'],depositHash,pair['fromToken'],recipient,final,false,!pair['isMain']).then(tx=>{
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
            'data-deposit-hash' => $this->row->depositHash,
            'class' => 'grid-check-row'
        ]);

        return parent::html();
    }
}
