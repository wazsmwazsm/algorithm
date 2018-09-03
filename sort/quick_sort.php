<?php

function quick($arr)
{
    $length = count($arr);
    // 注意要留出口
    if($length <= 1){
        return $arr;
    }

    $left = [];
    $right = [];
    $key = $arr[0];

    for ($i=1; $i < $length; $i++) { 
        if ($arr[$i] < $key) {
            $left[] = $arr[$i];
        } else {
            $right[] = $arr[$i];
        }
    }

    $left = quick($left);
    $right = quick($right);

    return array_merge($left, [$key,], $right);
}

// result
$a = [3, 1, 4, 5, 7, 2, 8];
$a = quick($a);
print_r($a);

// time
$numArr = [];
for($i=0;$i<50000;$i++){
    $numArr[] = $i;    
}
shuffle($numArr);

$t1 = microtime(true);
$numArr = quick($numArr);
$t2 = microtime(true);

echo '耗时'.round($t2 - $t1, 3).'秒';
