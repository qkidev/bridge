<?php

namespace App\Admin\Repositories;

use App\Models\Chain as Model;
use Dcat\Admin\Repositories\EloquentRepository;

class Chain extends EloquentRepository
{
    /**
     * Model.
     *
     * @var string
     */
    protected $eloquentClass = Model::class;
}
