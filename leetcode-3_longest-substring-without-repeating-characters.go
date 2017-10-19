/**
 * 给定一个字符串, 找到不重复字串的最大长度
 */
package main

import "fmt"

func main() {
    fmt.Println(lengthOfLongestSubstring("jbpnbwwd"))
}

func lengthOfLongestSubstring(s string) int {
    byteArr := []byte{}
    max := 0
    for i := 0; i < len(s); i++ { //逐个元素作为起点
        for j := i; j < len(s); j++ {
            if inArray(byteArr, s[j]) { //如果该字符已在子串中, 判断子串长度
                if len(byteArr) > max {
                    max = len(byteArr)
                }
                byteArr = []byte{} //子串清空, 下一个字符作为起点
                break
            } else {
                byteArr = append(byteArr, s[j]) //如果不在子串中, 字符加入子串
            }
        }
    }

    //处理不重复子串在最后的情况(不重复子串出现在最后时, 会没有更新到max)
    if len(byteArr) > max {
        max = len(byteArr)
    }
    return max
}

//判断字符是否在子串中
func inArray(arr []byte, char byte) bool {
    for _, v := range arr {
        if v == char {
            return true
        }
    }
    return false
}
