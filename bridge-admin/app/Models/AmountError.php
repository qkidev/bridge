<?php

namespace App\Models;

use Dcat\Admin\Traits\HasDateTimeFormatter;

use Illuminate\Database\Eloquent\Model;

class AmountError extends Model
{
	use HasDateTimeFormatter;
    protected $table = 'amount_error';
    public $timestamps = false;

}
