/**
 * 设计一个数据结构包含 insert(val), remove(val), getRandom() 三个方法, 并且这三个方法的时间复杂度为 O(1)
 * 使用数组保存 val, 使用 map 保存元素插入时的位置
 * 删除元素时先删掉 map 中的位置映射, 然后把目标元素与最后一个元素交换, 删除最后一个元素 
 */
type RandomizedSet struct {
    Nums []int
    Hm map[int]int
}

//初始化数据结构
func Constructor() RandomizedSet {
    return RandomizedSet{Nums : []int{}, Hm : map[int]int{}}
}

//插入元素
func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.Hm[val]; ok {
        return false
    }
    this.Nums = append(this.Nums, val) //插入到数组尾部
    this.Hm[val] = len(this.Nums) - 1 //记录元素位置
    return true
}

//删除元素
func (this *RandomizedSet) Remove(val int) bool {
    if idx, ok := this.Hm[val]; ok {
        last := len(this.Nums) - 1
        this.Hm[this.Nums[last]] = idx //记录尾部元素交换到的新位置
        delete(this.Hm, val) //删除 val 元素的位置记录
        this.Nums[idx], this.Nums[last] = this.Nums[last], this.Nums[idx] //val 和尾部元素交换位置
        this.Nums = this.Nums[:last] //删掉 val
        return true
    } else {
        return false
    }
}

//获取随机元素
func (this *RandomizedSet) GetRandom() int {
    return this.Nums[rand.Intn(len(this.Nums))]
}
