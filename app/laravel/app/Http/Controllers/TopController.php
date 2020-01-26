<?php

namespace App\Http\Controllers;

use App\Domain\AccessLog as domain_AccessLog;

class TopController extends Controller
{
    public function index()
    {
        $view_data = [];

        domain_AccessLog::add();

        $view_data['pv'] = domain_AccessLog::getCount();

        return view('top', compact('view_data'));
    }
}
