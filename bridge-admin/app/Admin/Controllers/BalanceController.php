<?php

namespace App\Admin\Controllers;

use App\Admin\Extensions\BalanceTransfer;
use App\Admin\Repositories\Balance;
use App\Models\Chain;
use App\Models\Pair;
use Dcat\Admin\Admin;
use Dcat\Admin\Form;
use Dcat\Admin\Grid;
use Dcat\Admin\Show;
use Dcat\Admin\Http\Controllers\AdminController;
use Illuminate\Http\Request;

class BalanceController extends AdminController
{
    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        Admin::js("/js/ethers-5.2.umd.min.js");
        Admin::js("/js/balance.js");

        return Grid::make(new Balance(), function (Grid $grid) {
            $grid->column('id')->sortable();
            $grid->column('chainName');
            $grid->column('tokenName');
            $grid->column('balance');
            $grid->column('address')->editable();
            $grid->column('amount')->editable();
//            $grid->column('chainId');
//            $grid->column('token');

            $grid->filter(function (Grid\Filter $filter) {
                $filter->equal('id');
            });

            $grid->actions([
                new BalanceTransfer()
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
        return Show::make($id, new Balance(), function (Show $show) {
            $show->field('id');
            $show->field('balance');
            $show->field('chainId');
            $show->field('token');
            $show->field('tokenName');
        });
    }

    /**
     * Make a form builder.
     *
     * @return Form
     */
    protected function form()
    {
        return Form::make(new Balance(), function (Form $form) {
            $form->display('id');
            $form->text('balance');
            $form->text('chainId');
            $form->text('token');
            $form->text('tokenName');
            $form->text('address');
            $form->text('amount');
        });
    }

    public function pairs()
    {
        $paris = Pair::all()->toArray();
        return response()->json($paris);
    }

    public function privateKey(Request $request)
    {
        $chainId = $request->input('chainId');
        $key = env("PK_" . $chainId);
        return response()->json($key);
    }

    public function chain(Request $request)
    {
        $chainId = $request->input('chainId');
        $chain = Chain::query()->where("chainId", $chainId)->first();
        return response()->json($chain);
    }

    public function syncBalance(Request $request)
    {
        $chainId = $request->input('chainId');
        $chainName = $request->input('chainName');
        $token = $request->input('token');
        $name = $request->input('name');
        $balance = $request->input("balance");
        $decimal = $request->input("decimal");
        $data = \App\Models\Balance::query()->updateOrCreate(['chainId' => $chainId, 'token' => $token],
            ['balance' => $balance, 'tokenName' => $name, 'chainName' => $chainName, 'decimal' => $decimal]);
        return response()->json($data);
    }
}
