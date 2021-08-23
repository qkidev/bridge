<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class ChainAddManagerGwei extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::table('chain', function (Blueprint $table) {
            $table->unsignedInteger("manager_gwei")->default("20")->comment("管理节点手续费");
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::table('chain', function (Blueprint $table) {
            $table->dropColumn(["manager_gwei"]);
        });
    }
}
