/**
 * 给定一个int数组表示运动员的成绩, 返回每个运动员的名次
 * 第一名返回"Gold Medal"
 * 第二名返回"Silver Medal"
 * 第三名返回"Bronze Medal"
 * 其他直接返回名次
 */
package main

import "fmt"
import "sort"
import "strconv"

func main() {
    findRelativeRanks([]int{10, 3, 8, 9, 4})
}

func findRelativeRanks(nums []int) []string {
    score := make([]int, len(nums))
    copy(score, nums)
    ret := []string{}
    scoreMap := map[int]int{}
    sort.Sort(sort.Reverse(sort.IntSlice(nums)))
    for k, v := range nums {
        scoreMap[v] = k
    }
    for _, val := range score {
        switch scoreMap[val] {
            case 0:
                ret = append(ret, "Gold Medal")
            case 1:
                ret = append(ret, "Silver Medal")
            case 2:
                ret = append(ret, "Bronze Medal")
            default:
                ret = append(ret, strconv.Itoa(scoreMap[val] + 1))
        }
    }
    return ret
}
