/**
 * 给定两个数组 g 和 s, g 表示小朋友的胃口, s 表示每份饼干的数量
 * 先把两个数组由低到高排序, 逐个小朋友发饼干, 当饼干的数量满足当前小朋友时, 才发下一个小朋友
 */
func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    fmt.Println(g)
    sort.Ints(s)
    fmt.Println(s)
    ret := 0 //ret 表示轮到第几个小朋友
    for i := 0; ret < len(g) && i < len(s); i++ {
        if s[i] >= g[ret] {
            ret++
        }
    }
    return ret
}
