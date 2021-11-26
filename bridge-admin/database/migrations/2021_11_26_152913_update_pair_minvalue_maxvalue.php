<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class UpdatePairMinvalueMaxvalue extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::table('pair', function (Blueprint $table) {
            $table->dropColumn(["minValue", "maxValue"]);
            $table->decimal('minValue',20,8,true);
            $table->decimal('maxValue',20,8,true);
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
