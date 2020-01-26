<?php

namespace App\DB;

class AccessLog
{
    protected static $table = 'access_log';

    public static function table()
    {
        return \DB::table(static::$table);
    }
}
