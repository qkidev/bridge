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
     */
    protected function script()
    {
        $pk = json_encode([
            '20181205' => env("PK_20181205")
        ]);

        $pairs = json_encode(
            Pair::all()->groupBy('id')->toArray()
        );

        $chains = json_encode(
            Chain::all()->groupBy('chainId')->toArray()
        );

        return <<<JS
$('.grid-check-withdraw').on('click', function () {
    alert('xxxxx')
    let pk = $pk
    let pairs = $pairs
    let chains = $chains
    const chainId = $(this).data('chain-id')
    const pairId = $(this).data('pair-id')
    const pair = pairs[pairId][0]
    console.log(pair)
    console.log(chains)
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
            'data-chain-id' => '20181205',
            'class' => 'grid-check-withdraw'
        ]);

        return parent::html();
    }
}
