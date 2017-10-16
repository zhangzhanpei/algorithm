/**
 * 给定一个int数组, 返回与它度数相同的子集的最小长度
 * 出现次数最多的元素的出现次数就是数组的度数
 * 找到一个长度最小的子集, 度数与原数组相同
 */
package main

import "fmt"

func main() {
    fmt.Println(findShortestSubArray([]int{1,2,2,3,1,4,2}))
}

func findShortestSubArray(nums []int) int {
    //保存每个元素的开始出现位置和最后出现位置
    firstPos, lastPos, frequency := map[int]int{}, map[int]int{}, map[int]int{}
    for k, v := range nums {
        frequency[v]++
        if _, ok := firstPos[v]; !ok {
            firstPos[v] = k
        }
        if lastPos[v] != k {
            lastPos[v] = k
        }
    }

    //找到出现次数最多的元素
    maxFrequency := 0
    for _, v := range frequency {
        if v > maxFrequency {
            maxFrequency = v
        }
    }

    //找到子集最小长度
    degree := len(nums)
    for k, v := range frequency {
        if v == maxFrequency {
            tmp := lastPos[k] - firstPos[k] + 1
            if tmp < degree {
                degree = tmp
            }
        }
    }
    return degree
}
