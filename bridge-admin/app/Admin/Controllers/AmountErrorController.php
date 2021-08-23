<?php

namespace App\Admin\Controllers;

use App\Admin\Repositories\AmountError;
use Dcat\Admin\Form;
use Dcat\Admin\Grid;
use Dcat\Admin\Show;
use Dcat\Admin\Http\Controllers\AdminController;

class AmountErrorController extends AdminController
{
    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        return Grid::make(new AmountError(), function (Grid $grid) {
            $grid->column('id')->sortable();
            $grid->column('tokenName');
            $grid->column('mainChain');
            $grid->column('mainBalance');
            $grid->column('toChain');
            $grid->column('toBalance');
            $grid->column('errNum');
            $grid->column('errLimit')->editable();

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
        return Show::make($id, new AmountError(), function (Show $show) {
            $show->field('id');
            $show->field('mainChain');
            $show->field('toChain');
            $show->field('tokenName');
            $show->field('mainBalance');
            $show->field('toBalance');
            $show->field('errNum');
            $show->field('errLimit');
        });
    }

    /**
     * Make a form builder.
     *
     * @return Form
     */
    protected function form()
    {
        return Form::make(new AmountError(), function (Form $form) {
            $form->display('id');
            $form->text('errLimit');
        });
    }
}
