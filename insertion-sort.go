/**
 * 插入排序
 * 逐个遍历元素, 如果当前元素比前面的元素小, 一直向前交换直到位置正确
 * 即前面有序部分比当前元素大的元素都后移一位, 把当前元素交换(插入)到正确位置
 */
package main

import "fmt"

func main() {
    insertionSort([]int{3, 7, 1, 15, 4, -5, 168})
}

func insertionSort(arr []int) []int {
    for i := 1; i < len(arr); i++ { //遍历数组元素
        for j := i; j >= 1 && arr[j] < arr[j - 1]; j-- { //如果当前元素比前面的小, 一直向前交换直到位置正确
            arr[j - 1], arr[j] = arr[j], arr[j - 1]
        }
    }
    return arr
}
