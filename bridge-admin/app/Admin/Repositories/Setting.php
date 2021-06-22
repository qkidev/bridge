<?php

namespace App\Admin\Repositories;

use App\Models\Setting as Model;
use Dcat\Admin\Repositories\EloquentRepository;

class Setting extends EloquentRepository
{
    /**
     * Model.
     *
     * @var string
     */
    protected $eloquentClass = Model::class;
}
