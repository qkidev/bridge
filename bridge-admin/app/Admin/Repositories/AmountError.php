<?php

namespace App\Admin\Repositories;

use App\Models\AmountError as Model;
use Dcat\Admin\Repositories\EloquentRepository;

class AmountError extends EloquentRepository
{
    /**
     * Model.
     *
     * @var string
     */
    protected $eloquentClass = Model::class;
}
