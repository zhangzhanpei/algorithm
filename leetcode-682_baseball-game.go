/**
 * 给定一个string数组包括数字和"C", "D", "+"三个操作符, 求总成绩
 * + 表示本轮分数等于前两轮分数的和
 * D 表示本轮分数等于前一轮的双倍
 * C 表示前一轮的成绩作废并删除
 */
package main

import "fmt"
import "strconv"

func main() {
    calPoints([]string{"5","-2","4","C","D","9","+","+"})
}

func calPoints(ops []string) int {
    list := []int{}
    for _, val := range ops {
        switch val {
            case "C":
                list = list[:len(list) - 1] //最后一轮成绩作废并删除
            case "D":
                num := list[len(list) - 1] * 2 //本轮成绩为最后一轮成绩的双倍
                list = append(list, num)
            case "+":
                num := list[len(list) - 1] + list[len(list) - 2] //本轮成绩为前两轮成绩之和
                list = append(list, num)
            default:
                num, _ := strconv.Atoi(val)
                list = append(list, num)
        }
    }

    sum := 0 //最后求总成绩
    for _, val := range list {
        sum += val
    }
    return sum
}
