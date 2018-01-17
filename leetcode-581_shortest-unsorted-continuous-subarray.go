/**
 * 给定一个无序int数组, 其中有一部分元素是无序的, 只要把这部分元素排序后则数组整体有序, 求这段无序数组的长度
 * 先排序数组, 然后和原数组比较, 从左边找到一个不同的元素, 从右边找到一个不同的元素
 */
package main

import "fmt"
import "sort"

func main() {
    fmt.Println(findUnsortedSubarray([]int{2,6,4,8,10,9,15}))
}

func findUnsortedSubarray(nums []int) int {
    arr := make([]int, len(nums))
    copy(arr, nums)
    sort.Ints(arr) //slice赋值并排序
    n := len(nums)
    i, j := 0, n - 1
    for ; i < n && arr[i] == nums[i]; i++ { //从左边开始比较

    }
    for ; j > i && arr[j] == nums[j]; j-- { //从右边开始比较

    }
    return j - i + 1
}
