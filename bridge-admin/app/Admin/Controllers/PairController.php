<?php

namespace App\Admin\Controllers;

use App\Admin\Extensions\InsertPair;
use App\Admin\Repositories\Pair;
use App\Models\Chain;
use App\Models\Log;
use Dcat\Admin\Admin;
use Dcat\Admin\Form;
use Dcat\Admin\Grid;
use Dcat\Admin\Http\Controllers\AdminController;
use Dcat\Admin\Show;
use Dcat\Admin\Widgets\Card;

class PairController extends AdminController
{
    const StatusLabel = ['关闭', '开启'];
    const IsMain = ['否', '是'];
    const IsNative = ['否', '是'];

    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        Admin::js("/js/ethers-5.2.umd.min.js");
        return Grid::make(new Pair(), function (Grid $grid) {
            $grid->column('id')->sortable();
            $grid->column('name');
            $grid->column('累计手续费')->display(function () {
                $fee = Log::query()->where("pairId", $this->id)->sum("fee");
                $style = "";
                if ($fee > 0) {
                    $style = 'color: red';
                }
                return "<html><b style=\"{$style}\">{$fee}</b></html>";
            });

            $grid->column('fromChain')->using(Chain::getChains());
            $grid->column('toChain')->using(Chain::getChains());
            $grid->column('icon')->image("", 50, 50);
            $grid->column('content', "详情")
                ->display('详情') // 设置按钮名称
                ->expand(function () {
                    // 返回显示的详情
                    // 这里返回 content 字段内容，并用 Card 包裹起来
                    $card = new Card();
                    $content = <<<EOF
                       资产名称: $this->title<br/><br/>
                       跨出Token	: $this->fromToken<br/><br/>
                       跨入Token: $this->toToken<br/><br/>
                       转账费(%): $this->tokenFee<br/><br/>
                       跨链费(%): $this->bridgeFee<br/><br/>
                       币种精度: $this->decimal<br/><br/>

EOF;
                    $card->content($content);
                    return "<div style='padding:10px 10px 0'>$card</div>";
                });

            $grid->column('isStop')->select(self::StatusLabel);
            $grid->column('isMain')->select(self::IsMain);
            $grid->column('isNative')->select(self::IsNative);
            $grid->column('minValue')->sortable()->editable();
            $grid->column('feeMin')->sortable()->editable();
            $grid->column('feeMax')->sortable()->editable();
            $grid->column('limit')->sortable()->editable();
            $grid->column('sort')->sortable()->editable();

            $grid->filter(function (Grid\Filter $filter) {
                $filter->expand()->panel();
                $filter->like("name")->width(3);
                $filter->equal('id', "交易对")->select(\App\Models\Pair::getPairs())->width(3);
                $filter->equal('fromChain')->select(Chain::getChains())->width(3);
                $filter->equal('toChain')->select(Chain::getChains())->width(3);
                $filter->equal('isStop')->select(self::StatusLabel)->width(3);
                $filter->equal('isMain')->select(self::IsMain)->width(3);
                $filter->equal('isNative')->select(self::IsNative)->width(3);

                $filter->like("fromToken")->width(3);
                $filter->like("toToken")->width(3);
            });

            $grid->actions([
                new InsertPair()
            ]);
        });
    }

    /**
     * Make a show builder.
     *
     * @param mixed $id
     *
     * @return Show
     */
    protected function detail($id)
    {
        return Show::make($id, new Pair(), function (Show $show) {
            $show->field('id');
            $show->field('bridgeFee');
            $show->field('decimal');
            $show->field('fromChain');
            $show->field('fromPair');
            $show->field('icon');
            $show->field('isMain');
            $show->field('minValue');
            $show->field('name');
            $show->field('sort');
            $show->field('title');
            $show->field('toChain');
            $show->field('tokenFee');
            $show->field('toPair');
        });
    }

    /**
     * Make a form builder.
     *
     * @return Form
     */
    protected function form()
    {
        $chainIds = Chain::query()->pluck("title", "chainId")->toArray();

        return Form::make(new Pair(), function (Form $form) use ($chainIds) {
            $form->display('id');
            $form->select('fromChain')
                ->options($chainIds)
                ->required();
            $form->select('toChain')->options($chainIds)->required();
            $form->text('name')->required();
            $form->text('title')->required();
            $form->text('decimal')->required()->default(0);
            $form->text('fromToken')
                ->default("0x0000000000000000000000000000000000000000")
                ->placeholder("如果是主网币跨出请输入 0x0000000000000000000000000000000000000000");
            $form->text('toToken')->required();
            $form->text('tokenFee')
                ->placeholder("资产本身转账需要扣除的手续费")
                ->default(0);
            $form->text('bridgeFee')
                ->placeholder("资产跨链需要扣除的手续费")
                ->default(0);
            $form->switch('isMain')->default(0);
            $form->switch('isNative')->default(0);
            $form->switch('isStop')->default(0);
            $form->text('minValue')->default(0);
            $form->text('limit')->default(0);
            $form->text('icon');
            $form->text('sort')->default(0);
            $form->text('feeMax')->default(0);
            $form->text('feeMin')->default(0);
        });
    }
}
