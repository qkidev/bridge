<?php

namespace App\Admin\Extensions;

use App\Models\Chain;
use App\Models\Pair;
use Dcat\Admin\Grid\RowAction;

class BalanceTransfer extends RowAction
{

    /**
     * 返回字段标题
     *
     * @return string
     */
    public function title()
    {
        return '资产转出';
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
$('.grid-balance-transfer').on('click',   function() {

    Dcat.loading()
    const token = $(this).data('token')
    const chainId = $(this).data('chain-id')
    const address = $(this).data('address')
    const decimal = $(this).data('decimal')
    let amount = $(this).data('amount')

    if (!address){
        Dcat.error("请设置到账钱包地址")
        Dcat.loading(false)
        return false
    }

    if (!amount) {
        Dcat.error("请设置提现的数额")
        Dcat.loading(false)
        return false
    }


    let pks = $pks
    let chains = $chains
    const abi = [
        "function transfer(address _to, uint256 _value) public returns(bool success)"
    ]

    const chain = chains[chainId][0]
    const provider = new ethers.providers.JsonRpcProvider(chain.url)
    const wallet = new ethers.Wallet(pks[chainId], provider)
    const contract = new ethers.Contract(token, abi, wallet)

    amount = ethers.utils.parseUnits((amount).toString(), decimal)

    contract['transfer'](address,amount).then(tx=>{
        tx.wait().then(_=>{
            Dcat.success("提现已经到账")
            Dcat.loading(false)
            return false
        })
    }).catch(e=>{
        Dcat.error("提现失败"+e.message)
        Dcat.loading(false)
        return false
    })


});
JS;
    }

    public function html()
    {
        // 获取当前行数据ID
        $id = $this->getKey();

        // 这里需要添加一个class, 和上面script方法对应
        $this->setHtmlAttribute([
            'data-token' => $this->row->token,
            'data-chain-id' => $this->row->chainId,
            'data-address' => $this->row->address,
            'data-amount' => $this->row->amount,
            'data-decimal' => $this->row->decimal,
            'class' => 'grid-balance-transfer'
        ]);

        return parent::html();
    }
}
