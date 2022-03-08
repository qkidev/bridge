<?php

namespace App\Models;

use Dcat\Admin\Traits\HasDateTimeFormatter;

use Illuminate\Database\Eloquent\Model;

class Log extends Model
{
	use HasDateTimeFormatter;
    protected $table = 'log';
    public $timestamps = false;

    protected $casts = [
        "depositTime" => "datetime",
        "withdrawTime" => "datetime",
    ];

}
