<?php

namespace App\Admin\Repositories;

use App\Models\Pair as Model;
use Dcat\Admin\Repositories\EloquentRepository;

class Pair extends EloquentRepository
{
    /**
     * Model.
     *
     * @var string
     */
    protected $eloquentClass = Model::class;
}
