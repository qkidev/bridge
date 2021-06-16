<?php

namespace App\Admin\Repositories;

use App\Models\Bridge as Model;
use Dcat\Admin\Repositories\EloquentRepository;

class Bridge extends EloquentRepository
{
    /**
     * Model.
     *
     * @var string
     */
    protected $eloquentClass = Model::class;
}
