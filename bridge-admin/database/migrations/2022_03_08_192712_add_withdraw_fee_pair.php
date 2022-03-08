<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class AddWithdrawFeePair extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::table('pair', function (Blueprint $table) {
            $table->string("withdraw_fee")->default("0")->comment("已提现手续费");
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
            $table->dropColumn(["withdraw_fee"]);
        });
    }
}
