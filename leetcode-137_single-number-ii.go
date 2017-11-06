/**
 * 给定一个int数组, 每个元素都出现3次, 但其中一个元素没有出现3次, 找到这个元素
 */
package main

import "fmt"

func main() {
    fmt.Println(singleNumber([]int{1,1,1,2,3,3,3}))
}

func singleNumber(nums []int) int {
    m := map[int]int{}
    for _, v := range nums { //统计每个元素出现的次数
        m[v]++
    }

    for k, v := range m { //找到没有出现3次的元素
        if v != 3 {
            return k
        }
    }
    return -1
}
