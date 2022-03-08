<?php

namespace App\Admin\Controllers;

use App\Admin\Repositories\Chain;
use Dcat\Admin\Form;
use Dcat\Admin\Grid;
use Dcat\Admin\Http\Controllers\AdminController;
use Dcat\Admin\Show;
use Dcat\Admin\Widgets\Card;

class ChainController extends AdminController
{
    const StatusLabel = ['关闭', '开启'];

    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        return Grid::make(new Chain(), function (Grid $grid) {
            $grid->column('id')->sortable();
            $grid->column('name');
            $grid->column('title');
            $grid->column('icon')->image("", 50);
            $grid->column('chainId');
            $grid->column('sort')->sortable();
            $grid->column('content', "详情")
                ->display('详情') // 设置按钮名称
                ->expand(function () {
                    // 返回显示的详情
                    // 这里返回 content 字段内容，并用 Card 包裹起来
                    $card = new Card();
                    $content = <<<EOF
                       节点: $this->url<br/><br/>
                       Bridge	: $this->bridge<br/><br/>
                       BridgeManager: $this->bridge_manager<br/><br/>
                       同步高度: $this->syncNumber<br/><br/>
                       单次同步数量: $this->syncLimit<br/>

EOF;
                    $card->content($content);
                    return "<div style='padding:10px 10px 0'>$card</div>";
                });

//            $grid->column('url');
            $grid->column('status')->select(self::StatusLabel);
            $grid->column('gwei');
            $grid->column('manager_gwei');
//            $grid->column('syncNumber');
//            $grid->column('syncLimit');
//            $grid->column('bridge');
//            $grid->column('bridge_manager');

            $grid->disableFilterButton();
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
            $form->select('status')->options(self::StatusLabel);
            $form->text('bridge');
            $form->text('bridge_manager');
            $form->text('gwei');
            $form->text('manager_gwei');
            $form->text('syncNumber');
            $form->text('syncLimit');
        });
    }
}
