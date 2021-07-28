<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreatePairTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('pair', function (Blueprint $table) {
            $table->increments('id');
            $table->decimal('bridgeFee')->nullable()->comment('跨链桥手续费');
            $table->unsignedInteger('decimal')->default('0')->nullable()->comment('精度');
            $table->unsignedInteger('fromChain')->default('0')->nullable()->comment('源链id');
            $table->string('fromToken')->nullable()->comment('源token');
            $table->string('icon')->nullable()->comment('图标地址');
            $table->unsignedSmallInteger('isMain')->default('0')->nullable()->comment('是否是主主链');
            $table->unsignedSmallInteger('isNative')->default('0')->nullable()->comment('是否是主网币');
            $table->unsignedInteger('limit')->default('0')->nullable()->comment('审核阈值');
            $table->string('maxValue')->default('0.00000000')->nullable()->comment('最大数额');
            $table->decimal('minValue')->nullable()->comment('最小数额');
            $table->string('name')->nullable()->comment('token名称');
            $table->unsignedInteger('sort')->default('0')->nullable()->comment('排序');
            $table->string('title')->nullable()->comment('标题');
            $table->integer('toChain')->nullable()->comment('目标链');
            $table->decimal('tokenFee')->nullable()->comment('目标token交易手续费');
            $table->string('toToken')->nullable()->comment('目标token');
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('pair');
    }
}
