<?php

use App\Http\Controllers\HelloWorldController;
use Illuminate\Support\Facades\Route;

Route::get('/', function () {
    return view('welcome');
});

Route::get('/hello', [HelloWorldController::class, 'helloWorld']);
