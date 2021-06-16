<?php

namespace App\Admin\Controllers;

use App\Admin\Repositories\Chain;
use Dcat\Admin\Form;
use Dcat\Admin\Grid;
use Dcat\Admin\Show;
use Dcat\Admin\Http\Controllers\AdminController;

class ChainController extends AdminController
{
    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        return Grid::make(new Chain(), function (Grid $grid) {
            $grid->column('id')->sortable();
            $grid->column('chainId');
            $grid->column('url');
            $grid->column('icon');
            $grid->column('name');
            $grid->column('sort');
            $grid->column('title');

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
        return Show::make($id, new Chain(), function (Show $show) {
            $show->field('id');
            $show->field('chainId');
            $show->field('url');
            $show->field('icon');
            $show->field('name');
            $show->field('sort');
            $show->field('title');
        });
    }

    /**
     * Make a form builder.
     *
     * @return Form
     */
    protected function form()
    {
        return Form::make(new Chain(), function (Form $form) {
            $form->display('id');
            $form->text('chainId');
            $form->text('icon');
            $form->text('name');
            $form->text('sort');
            $form->text('title');
            $form->text('url');
        });
    }
}
