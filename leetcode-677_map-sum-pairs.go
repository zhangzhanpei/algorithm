/**
 * 设计一个 MapSum 结构, 包含 insert 和 sum 方法
 * insert 方法给定 key 和 val, 如果 key 已存在, 覆盖 val
 * sum 方法给定一个前缀字符串, 如果 Map 中的 key 以该前缀开头, 则加上该 key 对应的 val
 */
type MapSum struct {
    m map[string]int
}

func Constructor() MapSum {
    return MapSum{m : map[string]int{}}
}

//插入复杂度O(1)
func (this *MapSum) Insert(key string, val int)  {
    this.m[key] = val
}

//求和复杂度O(n)
func (this *MapSum) Sum(prefix string) int {
    sum := 0
    for k, v := range this.m {
        if strings.HasPrefix(k, prefix) {
            sum += v
        }
    }
    return sum
}
