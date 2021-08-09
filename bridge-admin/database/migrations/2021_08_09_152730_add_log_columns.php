<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class AddLogColumns extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        if (!Schema::hasColumn('log', 'fee')) {
            Schema::table('log', function (Blueprint $table) {
                $table->string('fee')->default("")->comment('手续费');
            });
        }
        if (!Schema::hasColumn('log', 'amount')) {
            Schema::table('log', function (Blueprint $table) {
                $table->string('amount')->default("")->comment('到账数额');
            });
        }
        if (!Schema::hasColumn('log', 'withdrawSubmit')) {
            Schema::table('log', function (Blueprint $table) {
                $table->tinyInteger('withdrawSubmit')->default(0)->comment('是否提交了提现调用');
            });
        }
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
