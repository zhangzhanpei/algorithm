/**
 * 给定一个包含重复元素的 int 数组和一个 target, 随机返回大小等于 target 的元素下标
 * 同 leetcode-382 使用水塘抽样法
 */
type Solution struct {
    Sli []int
}


func Constructor(nums []int) Solution {
    return Solution{Sli:nums}
}


func (this *Solution) Pick(target int) int {
    sli := this.Sli
    count := 0
    ret := -1
    for key, val := range sli {
        if val != target {
            continue
        }
        count++
        //检查当前下标是否可取, 但必须所有等于 target 的下标都要求一遍概率
        //如果第一次符合就返回的话, 那么每次都是返回第一个符合的下标, 不符合题目"随机"的要求
        if rand.Intn(count) == 0 {
            ret = key
        }
    }
    return ret
}
