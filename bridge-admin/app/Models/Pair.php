<?php

namespace App\Models;

use Dcat\Admin\Traits\HasDateTimeFormatter;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Support\Facades\Cache;

class Pair extends Model
{
    use HasDateTimeFormatter;

    public $timestamps = false;
    protected $table = 'pair';

    public static function getPairs()
    {
        return Cache::remember('bridgePairs', 11400, function () {
            $pairs = [];
            $chainTitle = Chain::getChains();
            Pair::query()->get()->map(function ($item) use (&$pairs, $chainTitle) {
                $str = $item->name;
                $str .= " (" . $chainTitle[$item->fromChain];
                $str .= "=>" . $chainTitle[$item->toChain] . ")";
                $pairs[$item->id] = $str;
            });
            return $pairs;
        });
    }

}
