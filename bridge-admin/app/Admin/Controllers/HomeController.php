<?php

namespace App\Admin\Controllers;

use App\Admin\Metrics\Examples;
use App\Http\Controllers\Controller;
use App\Models\Bridge;
use App\Models\Chain;
use App\Models\Token;
use Dcat\Admin\Http\Controllers\Dashboard;
use Dcat\Admin\Layout\Column;
use Dcat\Admin\Layout\Content;
use Dcat\Admin\Layout\Row;
use Dcat\Admin\Widgets\Lazy;

class HomeController extends Controller
{
    public function index(Content $content)
    {
        return $content
            ->header('Dashboard')
            ->description('Description...')
            ->body(function (Row $row) {
                $row->column(6, function (Column $column) {
                    $column->row(Dashboard::title());
                    $column->row(new Examples\Tickets());
                });

                $row->column(6, function (Column $column) {
                    $column->row(function (Row $row) {
                        $row->column(6, new Examples\NewUsers());
                        $row->column(6, new Examples\NewDevices());
                    });

                    $column->row(new Examples\Sessions());
                    $column->row(new Examples\ProductOrders());
                });
            });
    }

    public function pairs(Content $content)
    {
        $chains = Chain::all()->toArray();
        $bridges = Bridge::all()->toArray();
        $tokens = Token::all()->toArray();
        return $content->body(view('admin.pairs', [
            "chains" => $chains, "bridges" => $bridges, "tokens" => $tokens
        ]));
    }
}
