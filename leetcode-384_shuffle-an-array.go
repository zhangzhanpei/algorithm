//打乱一个数组
import "math/rand"
import "time"

type Solution struct {
    Origin []int
    Shuffled []int
}

//构造函数
func Constructor(nums []int) Solution {
    l := len(nums)
    o := make([]int, l, l)
    copy(o, nums)
    return Solution{Origin: o, Shuffled: nums}
}

//返回原来的数组
func (this *Solution) Reset() []int {
    return this.Origin
}

//返回打乱后的数组
func (this *Solution) Shuffle() []int {
    l := len(this.Shuffled)
    for l > 0 {
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        j := r.Intn(l)
        this.Shuffled[l - 1], this.Shuffled[j] = this.Shuffled[j], this.Shuffled[l - 1]
        l--
    }
    return this.Shuffled
}
