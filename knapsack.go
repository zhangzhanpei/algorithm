/**
 * 0-1 背包问题
 */
package main

import "fmt"
import "math"

func main() {
    cap := []int{3, 5, 2, 6, 4} //每个物品的体积
    val := []int{4, 4, 3, 5, 3} //每个物品的价值
    knapsack(cap, val, 12) //12即背包总容量
}

func knapsack(cap, val []int, capacity int) {
    dp := make([]int, capacity + 1) //dp[j] 表示j容量的背包能得到的最大价值

    for i := 0; i < len(cap); i++ {
        for j := capacity; j >= cap[i]; j-- {
            //dp[j - cap[i]] + val[i] 取第i件物品后, 容量减少, 价值增加
            //dp[j] 不取第i件物品
            dp[j] = int(math.Max(float64(dp[j - cap[i]] + val[i]), float64(dp[j])))
        }
        //站在第0件物品前, 无论背包容量多大, 都只能取这一件, 所以在j>=3时val=4
        //[0 0 0 4 4 4 4 4 4 4 4 4 4]
        //站在第1件物品前, 当j>=8时可取两件val=8, 当capacity<8时只能取一件val=4
        //[0 0 0 4 4 4 4 4 8 8 8 8 8] 
        //站在第二件物品前, 当j>=10时可取3件val=11, 当8=<j<10时可取两件val=8, 当5<=j<8时可取两件val=7, 容量继续减小...
        //[0 0 3 4 4 7 7 7 8 8 11 11 11] 
        //[0 0 3 4 4 7 7 7 8 9 11 12 12]
        //[0 0 3 4 4 7 7 7 8 10 11 12 12]
    }
}
