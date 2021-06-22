<?php

namespace App\Admin\Controllers;

use App\Admin\Extensions\InsertPair;
use App\Admin\Repositories\Pair;
use Dcat\Admin\Admin;
use Dcat\Admin\Form;
use Dcat\Admin\Grid;
use Dcat\Admin\Show;
use Dcat\Admin\Http\Controllers\AdminController;

class PairController extends AdminController
{
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
            $grid->column('fromChain');
            $grid->column('toChain');
            $grid->column('name');
            $grid->column('title');
            $grid->column('fromToken');
            $grid->column('toToken');
            $grid->column('tokenFee');
            $grid->column('bridgeFee');
            $grid->column('decimal');
            $grid->column('icon')->image();
            $grid->column('isMain');
            $grid->column('isNative');
            $grid->column('minValue');
            $grid->column('limit')->editable();

            $grid->column('sort');


            $grid->filter(function (Grid\Filter $filter) {
                $filter->equal('id');

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
        return Form::make(new Pair(), function (Form $form) {
            $form->display('id');
            $form->select('fromChain')->required();
            $form->select('toChain')->required();
            $form->text('name')->required();
            $form->text('title')->required();
            $form->text('decimal')->required()->default(0);
            $form->text('fromToken')
                ->default("0x0000000000000000000000000000000000000000")
                ->placeholder("如果是主网币跨出请输入 0x0000000000000000000000000000000000000000");
            $form->text('toToken')->required();
            $form->text('tokenFee')
                ->placeholder("资产本身转账需要扣除的手续费")
                ->default(0)->required();
            $form->text('bridgeFee')
                ->placeholder("资产跨链需要扣除的手续费")
                ->default(0)->required();
            $form->switch('isMain')->default(0)->required();
            $form->switch('isNative')->default(0)->required();
            $form->text('minValue')->default(0)->required();
            $form->text('limit')->default(0)->required();
            $form->image('icon');
            $form->text('sort')->default(0);
        });
    }
}
