<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateChainTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('chain', function (Blueprint $table) {
            $table->increments('id');
            $table->string('bridge')->nullable()->comment('跨链桥合约地址');
            $table->string('bridge_manager')->nullable()->comment('管理员合约地址');
            $table->integer('chainId')->nullable()->comment('链id');
            $table->string('icon')->nullable()->comment('图标链接');
            $table->string('name')->nullable()->comment('名称');
            $table->integer('sort')->nullable()->comment('排序');
            $table->unsignedSmallInteger('status')->default('0')->nullable()->comment('状态');
            $table->string('title')->nullable()->comment('标题');
            $table->string('url')->nullable()->comment('链接');
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('chain');
    }
}
