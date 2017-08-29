/**
 * 给定一个二维数组, 里层数组都是升序, 求第k小的元素
 * 使用归并排序合并所有一维数组取第k个元素
 */
package main

import "fmt"

func main() {
    matrix := [][]int{
       {1, 2},
       {1, 3},
    }
    fmt.Println(kthSmallest(matrix, 2))
}

func kthSmallest(matrix [][]int, k int) int {
    arr := []int{}
    for i := 0; i < len(matrix); i++ {
        arr = mergeArray(arr, matrix[i])
    }
    return arr[k - 1]
}

//合并两个有序数组得到一个大的有序数组(归并排序)
func mergeArray(a []int, b []int) []int {
    tmp := []int{}
    i , j := 0, 0
    for i < len(a) && j < len(b) {
        if a[i] <= b[j] {
            tmp = append(tmp, a[i])
            i++
        } else {
            tmp = append(tmp, b[j])
            j++
        }
    }
    for i < len(a) {
        tmp = append(tmp, a[i])
        i++
    }
    for j < len(b) {
        tmp = append(tmp, b[j])
        j++
    }
    return tmp
}
