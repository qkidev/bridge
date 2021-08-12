<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class AddFeeMinMaxToPair extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::table('pair', function (Blueprint $table) {
            $table->string("feeMin")->default("0")->comment("最低手续费");
            $table->string("feeMax")->default("0")->comment("最高手续费");
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
            $table->dropColumn(["feeMin", "feeMax"]);
        });
    }
}
