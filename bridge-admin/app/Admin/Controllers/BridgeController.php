<?php

namespace App\Admin\Controllers;

use App\Admin\Repositories\Bridge;
use App\Models\Chain;
use Dcat\Admin\Form;
use Dcat\Admin\Grid;
use Dcat\Admin\Show;
use Dcat\Admin\Http\Controllers\AdminController;

class BridgeController extends AdminController
{
    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        return Grid::make(new Bridge(), function (Grid $grid) {
            $grid->column('id')->sortable();
            $grid->column('address');
            $grid->column('chainId')->using(
                Chain::query()->pluck("title", "id")->toArray()
            );

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
        return Show::make($id, new Bridge(), function (Show $show) {
            $show->field('id');
            $show->field('address');
            $show->field('chainId');
        });
    }

    /**
     * Make a form builder.
     *
     * @return Form
     */
    protected function form()
    {
        return Form::make(new Bridge(), function (Form $form) {
            $form->display('id');
            $form->text('address');
            $form->select('chainId')->options(
                Chain::query()->pluck("title", "id")
            );
        });
    }
}
