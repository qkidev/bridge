<?php

use Illuminate\Routing\Router;
use Illuminate\Support\Facades\Route;
use Dcat\Admin\Admin;

Admin::routes();

Route::group([
    'prefix' => config('admin.route.prefix'),
    'namespace' => config('admin.route.namespace'),
    'middleware' => config('admin.route.middleware'),
], function (Router $router) {

    $router->get('/', 'HomeController@index');
    $router->get('/withdraw', 'LogController@withdraw');
    $router->get('/getPairs','BalanceController@pairs');
    $router->get('/getPrivateKey','BalanceController@privateKey');
    $router->get('/getChain','BalanceController@chain');
    $router->get('/syncBalance','BalanceController@syncBalance');



    $router->resource('chains', ChainController::class);
    $router->resource('pairs', PairController::class);
    $router->resource('logs', LogController::class);
    $router->resource('settings', SettingController::class);
    $router->resource('balances', BalanceController::class);
});
