<?php

namespace App\Models;

use Dcat\Admin\Traits\HasDateTimeFormatter;

use Illuminate\Database\Eloquent\Model;

class Token extends Model
{
	use HasDateTimeFormatter;
    protected $table = 'token';
    public $timestamps = false;

}
