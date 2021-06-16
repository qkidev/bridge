<?php

namespace App\Models;

use Dcat\Admin\Traits\HasDateTimeFormatter;

use Illuminate\Database\Eloquent\Model;

class Chain extends Model
{
	use HasDateTimeFormatter;
    protected $table = 'chain';
    public $timestamps = false;

}
