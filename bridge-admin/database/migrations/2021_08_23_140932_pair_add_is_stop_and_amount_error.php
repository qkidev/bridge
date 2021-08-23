<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class PairAddIsStopAndAmountError extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::table('pair', function (Blueprint $table) {
            $table->unsignedTinyInteger("isStop")->default(0)->comment("是否关停");
        });

        Schema::create('amount_error', function (Blueprint $table) {
            $table->increments('id');
            $table->string('mainChain', 64)->comment('主链网络');
            $table->string('toChain', 64)->comment('侧链网络');
            $table->string('tokenName', 128)->comment('代币名称');
            $table->string('mainBalance', 64)->comment('主链余额');
            $table->string('toBalance', 64)->comment('侧链余额');
            $table->string('errNum', 64)->default(0)->comment('余额差值');
            $table->string('errLimit', 64)->default(0)->comment('余额差值限制');
            $table->unique(['mainChain', 'toChain', 'tokenName']);
        });

    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::table('pair', function (Blueprint $table) {
            $table->dropColumn(["isStop"]);
        });

        Schema::drop("amount_error");
    }
}
