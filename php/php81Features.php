<?php
/**
 * Created by PhpStorm.
 * User: cjx
 * Date: 2022/4/28
 * Time: 11:15
 */
// D:\phpstudy_pro\Extensions\php\php8.1.5nts\php.exe .\php81Feture.php

// 枚举
enum Status
{
    case Draft;
    case Published;
    case Archived;
    public function getName(): string
    {
        return $this->name;
    }
}
//var_dump(Status::Draft->name);
function acceptStatus(Status $status)
{
    echo "Status : " . $status->getName() . PHP_EOL;
}

acceptStatus(Status::Draft);

// 属性 readonly
class BlogData
{
    public readonly Status $status;

    public function __construct(Status $status)
    {
        $this->status = $status;
    }
}

// 纤程
/*$fiber = new Fiber(function (): void {
    $value = Fiber::suspend('fiber');
    echo "Value used to resume fiber: ", $value, PHP_EOL;
});
$value = $fiber->start();
echo "Value from fiber suspending: ", $value, PHP_EOL;
$fiber->resume('test');*/


