<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateBalanceTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('balance', function (Blueprint $table) {
            $table->increments('id');
            $table->string('address')->nullable()->comment('地址');
            $table->string('amount')->nullable()->comment('数额');
            $table->string('balance')->nullable()->comment('余额');
            $table->integer('chainId')->nullable()->comment('链id');
            $table->string('chainName')->nullable()->comment('链名称');
            $table->string('decimal')->nullable()->comment('精度');
            $table->string('token')->nullable()->comment('代币名称');
            $table->string('tokenName')->nullable()->comment('代币地址');
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('balance');
    }
}
