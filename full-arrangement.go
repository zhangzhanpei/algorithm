//给定一个字符数组, 返回字符的所有排列
package main

import "fmt"

var arr = []string{"a", "b", "c", "d"} //需要全排列的字符
var s = make([]string, 4, 4) //存储一种排列
var book = make(map[string]int) //标记某个字符是否已使用

func main() {
    fullArrangement(0) //从第一个位置开始尝试
}

func fullArrangement(step int) { //step指第几个位置
    if step == 4 { //如果到了第四个位置, 即所有字符已用完, 输入当前的排列
        fmt.Println(s)
        return //返回前一个位置
    }

    for i := 0; i < 4; i++ {
        if book[arr[i]] == 0 {
            s[step] = arr[i]
            book[arr[i]] = 1
            fullArrangement(step + 1)
            //当输出一种排列后返回, 取回当前step尝试的字符, 换成下一个字符试试.
            //即每一个step都会尝试所有的字符
            book[arr[i]] = 0
        }
    }
    return
}
