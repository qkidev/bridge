<?php

namespace App\Admin\Controllers;

use App\Admin\Repositories\Pair;
use App\Models\Chain;
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
        return Grid::make(new Pair(), function (Grid $grid) {
            $grid->column('id')->sortable();
            $grid->column('fromChain')->display(function () {
                return Chain::query()->where("id", $this->fromChain)->value("name");
            });
            $grid->column('toChain')->display(function () {
                return Chain::query()->where("id", $this->toChain)->value("name");
            });
            $grid->column('name');
            $grid->column('title');
            $grid->column('fromPair');
            $grid->column('toPair');
            $grid->column('tokenFee');
            $grid->column('bridgeFee');
            $grid->column('decimal');
            $grid->column('icon');
            $grid->column('isMain');
            $grid->column('isNative');
            $grid->column('minValue');

            $grid->column('sort');


            $grid->filter(function (Grid\Filter $filter) {
                $filter->equal('id');

            });
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
            $form->select('fromChain')->options(
                Chain::query()->pluck("name", "id")
            )->required();
            $form->select('toChain')->options(
                Chain::query()->pluck("name", "id")
            )->required();
            $form->text('name')->required();
            $form->text('title')->required();
            $form->text('decimal')->required()->default(0);
            $form->text('fromPair');
            $form->text('toPair')->required();
            $form->text('tokenFee')->default(0)->required();
            $form->text('bridgeFee')->default(0)->required();
            $form->switch('isMain')->default(0)->required();
            $form->switch('isNative')->default(0)->required();
            $form->text('minValue')->default(0)->required();
            $form->image('icon');
            $form->text('sort')->default(0);

        });
    }
}
