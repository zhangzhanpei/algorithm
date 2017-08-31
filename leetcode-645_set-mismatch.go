//给定一个n个元素的集合, 元素是1-n的所有值, 但是因为某些错误导致有一个元素重复, 找到重复元素和丢失的元素
package main

import "fmt"

func main() {
    fmt.Println(findErrorNums([]int{1, 2, 2, 4}))
}

func findErrorNums(nums []int) []int {
    m := map[int]bool{}
    n := len(nums)
    sum := (n * (n + 1)) / 2; //求数组的和
    duplicate := 0 //重复值
    for _, v := range nums {
        if m[v] {
            duplicate = v
        }
        sum -= v //因为有一个重复值, 减到最后的结果肯定不为0, 下面加回 duplicate 即得到丢失的值
        m[v] = true
    }
    return []int{duplicate, sum + duplicate}
}
