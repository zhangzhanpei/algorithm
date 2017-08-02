//给定一个int数组表示价格, 求进价和售价使得利润最大
package main

import "fmt"

func main() {
    fmt.Println(maxProfit([]int{3, 2, 6, 5, 0, 3}))
}

func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }
    min := prices[0]
    max := 0
    for i := 0; i < len(prices); i++ {
        if prices[i] > min { //当售价大于进价时, 取最大的售价
            tmp := prices[i] - min
            if tmp > max {
                max = tmp
            }
        } else {
            min = prices[i]
        }
    }
    return max
}
