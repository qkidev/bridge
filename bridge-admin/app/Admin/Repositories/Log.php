<?php

namespace App\Admin\Repositories;

use App\Models\Log as Model;
use Dcat\Admin\Repositories\EloquentRepository;

class Log extends EloquentRepository
{
    /**
     * Model.
     *
     * @var string
     */
    protected $eloquentClass = Model::class;
}
