/**
 * 给定一个int数组, 取第三大的元素(重复元素不算)
 * [3, 2, 1] -> 1
 * [1, 2] -> 2 //没有3个元素时返回最大的元素
 * [2, 2, 3, 1] -> 1 //2重复只算一次
 */
package main

import "fmt"

func main() {
    fmt.Println(thirdMax([]int{3,2,-2147483648}))
}

func thirdMax(nums []int) int {
    max, mid, min := -2147483649, -2147483649, -2147483649
    for _, val := range nums {
        if val == max || val == mid || val == min { //重复元素不算
            continue
        }

        //维护3个最大的元素
        if val > min {
            min = val
        }
        if min > mid {
            min, mid = mid, min
        }
        if mid > max {
            mid, max = max, mid
        }
    }

    //最小元素没有改变时, 即数组元素不够3个, 此时返回最大的元素
    if min == -2147483649 {
        return max
    } else { //否则返回第三大的元素
        return min
    }
}
