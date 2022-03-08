<?php

namespace App\Models;

use Dcat\Admin\Traits\HasDateTimeFormatter;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Support\Facades\Cache;

class Pair extends Model
{
	use HasDateTimeFormatter;
    protected $table = 'pair';
    public $timestamps = false;

    public static function getPairs()
    {
        return Cache::remember('bridge_pairs', 11400, function () {
            $pairs = [];
            Pair::query()->get()->map(function ($item) use (&$pairs) {
                $str = $item->name;
                $str .= " (" . $this->chainTitle[$item->fromChain];
                $str .= "=>" . $this->chainTitle[$item->toChain] . ")";
                $pairs[$item->id] = $str;
            });
            return $pairs;
        });
    }

}
