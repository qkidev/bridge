<?php

namespace App\Admin\Repositories;

use App\Models\Token as Model;
use Dcat\Admin\Repositories\EloquentRepository;

class Token extends EloquentRepository
{
    /**
     * Model.
     *
     * @var string
     */
    protected $eloquentClass = Model::class;
}
