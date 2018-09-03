<?php
// 小顶堆排序函数 $index 为父节点索引
function heap(&$arr, $index)
{
    // 计算子节点的索引 (按照从 0 开始的索引算，因为数组下标从 0 开始)
    $left = $index * 2 + 1;
    $right = $index * 2 + 2;
    // 没有子节点时返回
    if ( ! array_key_exists($left, $arr)) {
        return;
    }
    // 存在右子节点或者右节点小于左节点时，使用右节点作为和父节点交换的节点，否则使用左节点
    $key = array_key_exists($right, $arr) && ($arr[$right] < $arr[$left]) ?
        $right : $left;

    // 子节点和父节点判断大小，如果父节点大于子节点，则交换
    if ($arr[$index] > $arr[$key]) {
        $tmp = $arr[$index];
        $arr[$index] = $arr[$key];
        $arr[$key] = $tmp;
        // 将这个大的值继续往下排序
        heap($arr, $key);
    }
}

// reuslt
$a = [3, 1, 4, 5, 7, 2, 8];
// 构建小顶堆
// 获取最后一个有子节点的节点 (完全二叉树的计算方式, 根节点从 0 开始)
$index = (int)floor(count($a) / 2) - 1;
// 对每一个有子节点的节点构成的小顶堆都进行排序调整
// 不然的话只对 0 节点调整形成的不是一个堆结构，没法进行进一步的排序
for ($i = $index; $i >= 0; $i--) { 
    heap($a, $i);
}
print_r($a);
// 堆顶换值，进行排序
$a[0] = 12;
heap($a, 0);
print_r($a);

// time

// top K 
$numArr = [];
for($i=0;$i<50000;$i++){
    $numArr[] = $i;    
}
shuffle($numArr);

$top = array_splice($numArr, 0, 10);
// 构建小顶堆
$index = (int)floor(count($top) / 2) - 1;
for ($i=$index; $i >= 0; $i--) { 
    heap($top, $i);
}

$t1 = microtime(true);
// 遍历剩余数据，保持小顶堆中为最大数据
for ($i=count($top); $i < count($numArr); $i++) { 
    // 和堆顶比较，如果大于堆顶，则排序，否则丢弃
    if ($top[0] < $numArr[$i]) {
        // 替换
        $top[0] = $numArr[$i];
        // 重新排序
        heap($top, 0); 
    }
}

$t2 = microtime(true);

echo '耗时'.round($t2 - $t1, 3).'秒';

print_r($top);
