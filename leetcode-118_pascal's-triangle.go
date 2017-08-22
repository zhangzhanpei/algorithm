/**
 * 输出一个杨辉三角
 * 1
 * 1 1
 * 1 2 1
 * 1 3 3 1
 * 1 4 6 4 1
 */
package main

import "fmt"

func main() {
    generate(5)
}

func generate(numRows int) [][]int {
    ret := [][]int{}
    for i := 1; i <= numRows; i++ {
        tmp := make([]int, i) //逐层生成数组, 首尾元素为1
        tmp[0] = 1
        tmp[len(tmp) - 1] = 1
        ret = append(ret, tmp)
    }
    //从第3层开始, 排除首尾元素, 使用上一层元素计算当前层的元素
    for h := 2; h < numRows; h++ {
        for j := 1; j < h; j++ {
            ret[h][j] = ret[h - 1][j - 1] + ret[h - 1][j]
        }
    }
    return ret
}
