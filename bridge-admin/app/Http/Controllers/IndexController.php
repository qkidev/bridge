<?php

namespace App\Http\Controllers;

use App\Models\Setting;
use Illuminate\Http\Request;

class IndexController extends Controller
{
    public function getLock(Request $request)
    {
        $chain = $request->input('chain');
        return Setting::query()->where("name", "lock_{$chain}")->value("value");
    }

    public function setLock(Request $request)
    {
        $chain = $request->input('chain');
        $number = $request->input('number');
        return Setting::query()->where("name", "lock_{$chain}")->update([
            'value' => $number
        ]);
    }

    public function pair(Request $request)
    {
        $fromChain = $request->input('fromChain');
        $toChain = $request->input('toChain');
        $fromToken = $request->input('fromToken');
        $toToken = $request->input('toToken');

    }
}
