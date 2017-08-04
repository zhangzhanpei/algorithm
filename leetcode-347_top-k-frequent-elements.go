/**
 * 给定一个int数组, 取出现频率最高的 k 个元素
 */
package main

import "fmt"
import "container/list"

func main() {
    fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}

func topKFrequent(nums []int, k int) []int {
    //先把元素和出现的次数写进一个map
    m := map[int]int{}
    for _, val := range nums {
        m[val] += 1
    }

    //以次数为下标, 将出现该次数的所有元素存为链表放到数组中
    bucket := make([]*list.List, len(nums) + 1)
    for key, val := range m {
        if bucket[val] == nil {
            bucket[val] = list.New()
        }
        bucket[val].PushBack(key) //把元素插入到双向链表中
    }

    //下标越大即出现次数越多, 从数组后面往回读, 取前 k 个元素
    res := []int{}
    for i := len(bucket) - 1; i >= 0 && k > 0; i-- {
        if bucket[i] != nil {
            k -= bucket[i].Len()
            for e := bucket[i].Front(); e != nil; e = e.Next() {
                res = append(res, e.Value.(int)) //e.Value 是 interface{} 类型, 使用断言转成 int 类型
            }
        }
    }
    return res
}
