<?php

namespace App\Models;

use Dcat\Admin\Traits\HasDateTimeFormatter;

use Illuminate\Database\Eloquent\Model;

class Setting extends Model
{
    public $timestamps = false;
	use HasDateTimeFormatter;
    protected $table = 'setting';

}
