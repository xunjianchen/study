<?php
/**
 * Created by PhpStorm.
 * User: cjx
 * Date: 2022/4/28
 * Time: 11:15
 */

//Match 是一个表达式，它可以储存到变量中亦可以直接返回。
//Match 分支仅支持单行，它不需要一个 break; 语句。
//Match 使用严格比较。
function getStr($str)
{
    return match ($str) {
        1 => (function () {
            return 1;
        })(),
        'a' =>(function(){
            return 'a';
        })(),
        default =>'error'
    };
}
//var_dump(getStr(2));
echo match (8.0) {
    '8.0' => "Oh no!",
    8.0 => "This is what I expected",
};

// php7
/*$country =  null;
if ($session !== null) {
    $user = $session->user;
    if ($user !== null) {
        $address = $user->getAddress();
        if ($address !== null) {
            $country = $address->country;
        }
    }
}*/

// php8   Nullsafe 运算符
//$country = $session?->user?->getAddress()?->country;