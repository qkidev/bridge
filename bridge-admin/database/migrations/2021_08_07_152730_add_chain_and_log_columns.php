<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class AddChainAndLogColumns extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::table('chain', function (Blueprint $table) {
            $table->bigInteger('syncNumber')->default(0);
            $table->integer('syncNumber')->default(5000);
        });

        Schema::table('log', function (Blueprint $table) {
            $table->string('fee')->default("");
            $table->string('amount')->default("");
            $table->tinyInteger('withdrawSubmit')->default(0);
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        //
    }
}
