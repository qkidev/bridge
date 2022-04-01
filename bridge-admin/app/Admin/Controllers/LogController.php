<?php

namespace App\Admin\Controllers;

use App\Admin\Extensions\CheckWithdraw;
use App\Admin\Repositories\Log;
use App\Models\Chain;
use App\Models\Pair;
use Dcat\Admin\Admin;
use Dcat\Admin\Form;
use Dcat\Admin\Grid;
use Dcat\Admin\Show;
use Dcat\Admin\Http\Controllers\AdminController;
use Dcat\Admin\Widgets\Card;
use Illuminate\Http\Request;
use function PHPUnit\Framework\isNull;

class LogController extends AdminController
{
    const StatusLabel = ['否', '是'];

    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        Admin::js("/js/ethers-5.2.umd.min.js");
        return Grid::make(new Log(), function (Grid $grid) {
            $grid->model()->orderByDesc("id");
            $grid->column('id')->sortable();
            $grid->column('pairId')->using(Pair::getPairs());
            $grid->column('withdrawTime')->sortable();
            $grid->column('withdrawSubmit')
                ->if(function ($column) {
                    // 获取当前行其他字段值
                    $withdrawTime = $this->withdrawTime;
                    return empty($withdrawTime);
                })
                ->select(self::StatusLabel)
                ->else()
                ->using(self::StatusLabel);
            $grid->column('is_fail')
                ->if(function ($column) {
                    // 获取当前行其他字段值
                    $withdrawTime = $this->withdrawTime;
                    return empty($withdrawTime);
                })
                ->select(self::StatusLabel)
                ->else()
                ->using(self::StatusLabel);

            $grid->column('content', "详情")
                ->display('详情') // 设置按钮名称
                ->expand(function () {
                    // 返回显示的详情
                    // 这里返回 content 字段内容，并用 Card 包裹起来
                    $card = new Card();
                    $content = <<<EOF
                       接收者: $this->recipient<br/><br/>
                       数额	: $this->value<br/><br/>
                       手续费: $this->fee<br/><br/>
                       实际到账: $this->amount<br/><br/>
                       跨出hash: $this->depositHash<br/><br/>
                       到账hash: $this->withdrawHash<br/><br/>
                       备注: $this->remark<br/><br/>

EOF;
                    $card->content($content);
                    return "<div style='padding:10px 10px 0'>$card</div>";
                });
            $grid->column('depositTime')->sortable();
            $grid->column('overMax')->using(self::StatusLabel);

            $grid->filter(function (Grid\Filter $filter) {
                $filter->panel()->expand();
                $filter->equal("depositHash")->width(6);
                $filter->equal("withdrawHash")->width(6);
                $filter->like("recipient")->width(3);
                $filter->equal('pairId')->select(Pair::getPairs())->width(3);
                $filter->equal('fromChain')->select(Chain::getChains())->width(3);
                $filter->equal('toChain')->select(Chain::getChains())->width(3);
                $filter->equal('withdrawSubmit')->select(self::StatusLabel)->width(3);
                $filter->equal('is_fail')->select(self::StatusLabel)->width(3);
                $filter->equal('overMax')->select(self::StatusLabel)->width(3);
            });

            $grid->actions([
                new CheckWithdraw()
            ]);
            $grid->disableEditButton();
            $grid->disableCreateButton();
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
        return Show::make($id, new Log(), function (Show $show) {
            $show->field('id');
            $show->field('depositHash');
            $show->field('depositTime');
            $show->field('fromChain');
            $show->field('pairId');
            $show->field('recipient');
            $show->field('toChain');
            $show->field('value');
            $show->field('withdrawHash');
            $show->field('withdrawTime');
        });
    }

    /**
     * Make a form builder.
     *
     * @return Form
     */
    protected function form()
    {
        return Form::make(new Log(), function (Form $form) {
            $form->display('id');
            $form->text('depositHash');
            $form->text('depositTime');
            $form->text('fromChain');
            $form->text('pairId');
            $form->text('recipient');
            $form->text('toChain');
            $form->text('value');
            $form->text('withdrawHash');
            $form->text('withdrawTime');
            $form->select('withdrawSubmit')->options(self::StatusLabel);
            $form->select('is_fail')->options(self::StatusLabel);
        });
    }

    public function withdraw(Request $request)
    {
        $id = $request->input('logId');
        $hash = $request->input('hash');
        $microtime = time();
        \App\Models\Log::query()->where("id", $id)
            ->update(['withdrawTime' => $microtime, 'withdrawHash' => $hash]);
        return response()->json(['code' => 0, 'msg' => 'success']);
    }
}
