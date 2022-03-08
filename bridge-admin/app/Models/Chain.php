<?php

namespace App\Models;

use Dcat\Admin\Traits\HasDateTimeFormatter;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Support\Facades\Cache;

class Chain extends Model
{
	use HasDateTimeFormatter;
    protected $table = 'chain';
    public $timestamps = false;

    public static function getChains()
    {
       return Cache::remember('bridge_network', 11400, function () {
            return Chain::query()->pluck("title", "chainId")->toArray();
        });
    }

}
