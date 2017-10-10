/**
 * 给定一个int数组, 求等差数列的个数[等差数列元素是给定数组的至少3个连续元素]
 * input [1,2,3,4], return 3. 共有[1,2,3], [2,3,4], [1,2,3,4]三个等差数列
 * 规律如下
 * 输入数组       等差数列数量        与上一个输入数组的结果比较
 * [123]          1                   1-0=1
 * [1234]         3                   3-1=2
 * [12345]        6                   6-3=3
 * [123456]       10                  10-6=4
 * 每增加一个等差元素, 等差数列的数量会按照12345..的数量增加, 当下一个元素不是等差元素时, 数量又从0开始增加
 */
package main

import "fmt"

func main() {
    numberOfArithmeticSlices([]int{1,2,3,4,5})
}

func numberOfArithmeticSlices(A []int) int {
    count, add := 0, 0
    for i := 2; i < len(A); i++ {
        if A[i] - A[i - 1] == A[i - 1] - A[i - 2] {
            add++ //增加数量逐渐增大
            count += add
        } else {
            add = 0
        }
    }
    return count
}
