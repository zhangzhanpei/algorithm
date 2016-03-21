<?php
//选择排序
//先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置，然后，再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。直到所有元素均排序完毕。
//相比冒泡排序，选择排序少了很多次的交换数组元素，因为一趟只找出最小（大）元素的下标，然后只交换一次
function selectionSort(array $arr)
{
    $len = count($arr);
    for ($i = 0; $i < $len; $i++) {
        $index = $i;
        for ($j = $i; $j < $len; $j++) {
            if ($arr[$j] < $arr[$index]) {
                $index = $j;
            }
        }
        if ($index != $i) {
            $tmp         = $arr[$i];
            $arr[$i]     = $arr[$index];
            $arr[$index] = $tmp;
        }
    }
    return $arr;
}

$arr = [3, 7, 1, 15, 4, -5, 168];
var_dump(selectionSort($arr));
