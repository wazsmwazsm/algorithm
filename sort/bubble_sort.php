<?php

function bubble(&$arr)
{
    $count = count($arr);
    for ($i=0; $i < $count - 1; $i++) { 
        for ($j=0; $j < $count - $i - 1; $j++) { 
            if ($arr[$j] > $arr[$j+1]) {
                $tmp = $arr[$j];
                $arr[$j] = $arr[$j+1];
                $arr[$j+1] = $tmp;
            }
        }
    }
}

// result
$a = [3, 1, 4, 5, 7, 2, 8];
bubble($a);
print_r($a);

// time
$numArr = [];
for($i=0;$i<50000;$i++){
    $numArr[] = $i;    
}
shuffle($numArr);

$t1 = microtime(true);
bubble($numArr);
$t2 = microtime(true);

echo '耗时'.round($t2 - $t1, 3).'秒';
