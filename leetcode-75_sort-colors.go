/**
 * 给定一个int数组包含 0, 1, 2 表示红, 白, 蓝, 对数组进行排序
 */
package main

import "fmt"

func main() {
    fmt.Println(sortColors([]int{2, 0, 1}))
}

func sortColors(nums []int)  {
    sort.Ints(nums)
}
