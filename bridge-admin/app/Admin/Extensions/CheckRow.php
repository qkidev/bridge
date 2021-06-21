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

    Dcat.loading({
    color: Dcat.color.danger,
    })
    const pairId = $(this).data('pair-id')
    const recipient = $(this).data('recipient')
    const value = $(this).data('value')
    let pks = $pks
    let pairs = $pairs
    let chains = $chains
    const abi = [
    "event WithdrawDone(uint toChainId, address fromToken, address toToken, address recipient, uint256 value)",
    "event WithdrawNativeDone(uint fromChainId, address recipient,bool isMain, uint256 value)",
    "function withdraw(uint toChainId,address toToken,address recipient,uint256 value)",
    "function withdrawNative(uint fromChainId, address payable recipient, bool isMain, uint256 value)"
    ]
    const pair = pairs[pairId][0]
    const chain = chains[pair['toChain']][0]

    const provider = new ethers.providers.JsonRpcProvider(chain.url)
    const wallet = new ethers.Wallet(pks[pair['toChain']], provider)
    const bridge = new ethers.Contract(chain['bridge'], abi, wallet)

    const final = ethers.utils.parseUnits(value.toString(), pair['decimal'])
    if (pair.isNative){
    }else{
       bridge['withdraw'](pair['fromChain'],pair['fromToken'],recipient,final).then(tx=>{
           console.log(tx.hash)
            tx.wait()
            console.log('done')
       })
    }

    let index = setInterval(()=>{
        Dcat.loading(false)
        Dcat.success("审核操作已经完成")
        clearInterval(index)
    },5000)
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
            'data-recipient' => $this->row->recipient,
            'data-value' => $this->row->value,
            'class' => 'grid-check-row'
        ]);

        return parent::html();
    }
}
