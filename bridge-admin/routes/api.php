<?php

use App\Models\Chain;
use App\Models\Pair;
use Illuminate\Database\Eloquent\Builder;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/
Route::get('bridges', function () {
    $bridges = Chain::query()->pluck("bridge", "chainId");
    return response()->json($bridges);
});

Route::get('pairs', function () {
    $all = Pair::all();
    $pairs = $all->groupBy("fromChain")->all();
    return response()->json($pairs);
});

Route::get("items", function (Request $request) {
    $chainId = $request->input("chain");
    $fromChain = Chain::query()->where("chainId", $chainId)->first();
    if (!$fromChain) {
        return response()->json([
            'code' => 500,
            'msg' => "功能暂不支持该链"
        ]);
    }
    $token = $request->input("token");
    $items = Pair::query()->where("fromChain", $fromChain->chainId)->where("isStop",0)->when($token, function (Builder $builder, $fromToken) {
        return $builder->where("fromToken", $fromToken);
    })->orderBy("sort")->get();
    foreach ($items as $item) {
        $item->fromChainData = Chain::query()->where("chainId", $item->fromChain)->first();
        $item->toChainData = Chain::query()->where("chainId", $item->toChain)->first();
    }
    $grouped = $items->groupBy('name');

    $gweis = Chain::query()->pluck("gwei", "chainId");

    return response()->json([
        'code' => 0,
        'data' => $grouped->all(),
        'gweis' => $gweis
    ]);
});


