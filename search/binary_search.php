<?php

function binary_search($arr, $value)
{
    $max = count($arr);
    $min = 0;

    if ($max <= 0) {
        return false;
    }

    while ($min <= $max) {
        $mid = (int)(($min + $max) / 2);
        if ($arr[$mid] == $value) {
            return true;
        } else if ($arr[$mid] > $value) {
            $max = $mid - 1;
        } else {
            $min = $mid + 1;
        }
    }

    return false;
}

// result
$arr = [1,2,4,6,10,40,50,80,100,110];
var_dump(binary_search($arr, 50));
var_dump(binary_search($arr, 22));

// time
$numArr = [];
for($i=0;$i<500000;$i++){
    $numArr[] = $i;    
}

$t1 = microtime(true);
var_dump(binary_search($numArr, 9987));
$t2 = microtime(true);

echo '耗时'.round($t2 - $t1, 3).'秒';
