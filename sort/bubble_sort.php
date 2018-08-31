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

$a = [4, 3, 5, 8, 1, 6];

bubble($a);

var_dump($a);
