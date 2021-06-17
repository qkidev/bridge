<?php

namespace App\Models;

use Dcat\Admin\Traits\HasDateTimeFormatter;

use Illuminate\Database\Eloquent\Model;

class Pair extends Model
{
	use HasDateTimeFormatter;
    protected $table = 'pair';
    public $timestamps = false;

}
