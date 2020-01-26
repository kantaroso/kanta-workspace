<?php

namespace App\Domain;

use App\DB\AccessLog as db_AccessLog;
use Faker\Provider\zh_CN\DateTime;
use Illuminate\Http\Request;

class AccessLog
{

    public static function add()
    {
        $now = new \DateTime();
        $request = request();

        $insert_data = [
            'method' => $request->method(),
            'endpoint' => $request->path(),
            'query_string' => json_encode($request->toArray(), JSON_PRETTY_PRINT),
            'user_agent' => $request->userAgent(),
            'created_at' => $now,
            'updated_at' => $now,
        ];

        if (!db_AccessLog::table()->insert($insert_data) ) {
            \Log::error('AccessLog insert error.');
            \Log::error(print_r($insert_data, true));
        }

        return;
    }

    public static function getCount()
    {
        return count(db_AccessLog::table()->get());
    }
}
