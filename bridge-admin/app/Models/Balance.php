<?php

namespace App\Models;

use Dcat\Admin\Traits\HasDateTimeFormatter;

use Illuminate\Database\Eloquent\Model;

class Balance extends Model
{
    use HasDateTimeFormatter;

    protected $table = 'balance';
    public $timestamps = false;
    protected $guarded = [];

}
