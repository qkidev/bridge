<?php

namespace App\Admin\Repositories;

use App\Models\Balance as Model;
use Dcat\Admin\Repositories\EloquentRepository;

class Balance extends EloquentRepository
{
    /**
     * Model.
     *
     * @var string
     */
    protected $eloquentClass = Model::class;
}
