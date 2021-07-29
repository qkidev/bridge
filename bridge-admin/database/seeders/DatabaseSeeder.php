<?php

namespace Database\Seeders;

use Illuminate\Database\Seeder;
use Illuminate\Support\Facades\DB;

class DatabaseSeeder extends Seeder
{
    /**
     * Seed the application's database.
     *
     * @return void
     */
    public function run()
    {
        DB::table('admin_menu')->insert([
            ["parent_id" => 0, "order" => 9, "title" => "主链", "icon" => "fa-bold", "uri" => "chains", "show" => "1"],
            ["parent_id" => 0, "order" => 10, "title" => "跨链对", "icon" => "fa-exchange", "uri" => "pairs", "show" => "1"],
            ["parent_id" => 0, "order" => 11, "title" => "跨链记录", "icon" => "fa-exchange", "uri" => "logs", "show" => "1"],
            ["parent_id" => 0, "order" => 12, "title" => "配置", "icon" => "fa-bars", "uri" => "settings", "show" => "1"],
            ["parent_id" => 0, "order" => 13, "title" => "跨链桥余额", "icon" => "fa-shopping-bag", "uri" => "balances", "show" => "1"],
        ]);
        DB::table("setting")->insert([
            ['name' => "withdraw-check", "value" => 0, "description" => "跨链是否需要后台审核 1是 0否"]
        ]);
        // \App\Models\User::factory(10)->create();
    }
}
