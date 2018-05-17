/**
 * 一个机器人位于坐标(0, 0)处, 给定一个字符串(U上D下L左R右)表示移动过程, 判断最后机器人是否回到原点
 */
package main

import "fmt"

func main() {
    fmt.Println(judgeCircle("UD"))
}

func judgeCircle(moves string) bool {
    var x, y int
    for _, c := range moves {
        switch c {
            case 'U':
                y++
            case 'D':
                y--
            case 'L':
                x--
            case 'R':
                x++
        }
    }
    if x == 0 && y == 0 {
        return true
    }
    return false
}
