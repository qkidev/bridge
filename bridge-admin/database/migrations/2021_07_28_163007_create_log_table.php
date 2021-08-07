<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateLogTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('log', function (Blueprint $table) {
            $table->increments('id');
            $table->string('depositHash')->nullable()->comment('质押hash');
            $table->string('depositTime')->nullable()->comment('质押时间');
            $table->string('fromChain')->nullable()->comment('源自链');
            $table->integer('pairId')->nullable()->comment('跨链对id');
            $table->string('recipient')->nullable()->comment('接收地址');
            $table->string('toChain')->nullable()->comment('目标token地址');
            $table->string('value')->nullable()->comment('数额');
            $table->string('withdrawHash')->nullable()->comment('提现hash');
            $table->string('withdrawTime')->nullable()->comment('提现时间');
            $table->string('fee')->nullable()->comment('手续费');
            $table->string('amount')->nullable()->comment('到账数额');
            $table->tinyInteger('withdrawSubmit')->default(0)->comment('是否提交了提现调用');
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('log');
    }
}
