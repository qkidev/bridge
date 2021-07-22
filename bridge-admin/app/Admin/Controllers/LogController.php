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
use Illuminate\Http\Request;

class LogController extends AdminController
{
    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        Admin::js("/js/ethers-5.2.umd.min.js");
        return Grid::make(new Log(), function (Grid $grid) {
            $grid->model()->latest('id');
            $chainTitle = Chain::query()->pluck("title", "chainId")->toArray();
            $grid->column('id')->sortable();
            $grid->column('description')->display(function () use ($chainTitle) {
                $pair = Pair::query()->find($this->pairId);
                $fromChain = $chainTitle[$pair->fromChain];
                $toChain = $chainTitle[$pair->toChain];
                return "{$pair->name} 从 {$fromChain} 跨到 {$toChain}";
            });
            $grid->column('recipient');
            $grid->column('depositTime')->display(function () {
                return date("Y-m-d H:i:s", $this->depositTime);
            });
            $grid->column('value');
            $grid->column('withdrawTime')->display(function () {
                if ($this->withdrawTime) {
                    return date("Y-m-d H:i:s", $this->withdrawTime);
                }
            });
            $grid->filter(function (Grid\Filter $filter) {
                $filter->equal('id');
            });

            $grid->actions([
                new CheckWithdraw()
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
        });
    }

    public function withdraw(Request $request)
    {
        $id = $request->input('logId');
        $hash = $request->input('hash');
        $microtime = time();
        \App\Models\Log::query()->where("id", $id)
            ->update(['withdrawTime' => $microtime, 'withdrawHash' => $hash]);
        return response()->json(['code'=>0,'msg'=>'success']);
    }
}
