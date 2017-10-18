/**
 * 给定一面由二维数组组成的墙, 现需要在墙上打洞, 求至少要穿过的砖的数量
 * 一维数组的元素表示砖的长度, 要穿过最少的砖, 即尽量多穿过砖缝
 */
package main

import "fmt"

func main() {
    fmt.Println(leastBricks([][]int{
        {1,2,2,1},
        {3,1,2},
        {1,3,2},
        {2,4},
        {3,1,4},
        {1,3,1,1},
    }))
}

func leastBricks(wall [][]int) int {
    m := map[int]int{}
    //遍历每层砖, 使用map记录砖缝离左侧基线的距离
    for i := 0; i < len(wall); i++ {
        tmp := 0
        for j := 0; j < len(wall[i]) - 1; j++ {
            tmp += wall[i][j]
            m[tmp]++
        }
    }

    //找到出现次数最多的距离, 说明这个距离的砖缝最多
    maxFrequency := 0
    for _, v := range m {
        if v > maxFrequency {
            maxFrequency = v
        }
    }

    //砖的层数减去砖缝的层数, 即要穿过的砖的数量
    return len(wall) - maxFrequency
}
