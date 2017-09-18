/**
 * 选择排序
 * 每一趟都在未排序部分选择最小(最大)元素, 交换到已排序部分的后面
 */
package main

import "fmt"

func main() {
    selectionSort([]int{3, 7, 1, 15, 4, -5, 168})
}

func selectionSort(arr []int) []int {
    n := len(arr)
    for i := 0; i < n; i++ {
        pos := i //pos 记录最小元素的下标
        for j := i; j < n; j++ { //在后面未排序部分找到最小元素
            if arr[j] < arr[pos] {
                pos = j
            }
        }
        if i != pos { //如果最开头的就是最小元素, 即已跟在有序部分后面, 此时不用交换
            arr[i], arr[pos] = arr[pos], arr[i]
        }
    }
    return arr
}
