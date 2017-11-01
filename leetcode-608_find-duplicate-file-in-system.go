/**
 * 给定一个包含路径和内容的文件数组, 返回重复的文件
 */
package main

import "fmt"
import "strings"

func main() {
    findDuplicate([]string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"})
}

var m map[string][]string

func findDuplicate(paths []string) [][]string {
    m = map[string][]string{}
    ret := [][]string{}
    //遍历文件数组, 构建map
    for _, v := range paths {
        mappingFileAndContent(strings.Split(v, " "))
    }
    //如果一个内容有两个以上的文件, 则返回重复的文件
    for _, v := range m {
        if len(v) >= 2 {
            ret = append(ret, v)
        }
    }
    return ret
}

//以文件内容为key, 文件路径为value, 存入map
func mappingFileAndContent(item []string) {
    path := item[0]
    for i := 1; i < len(item); i++ {
        file, content := splitFileAndContent(item[i])
        if _, ok := m[content]; ok {
            m[content] = append(m[content], path + "/" + file)
        } else {
            m[content] = []string{path + "/" + file}
        }
    }
}

//将文件名和内容分割返回
func splitFileAndContent(s string) (file, content string) {
    tmp := strings.Split(strings.TrimRight(s, ")"), "(")
    file = tmp[0]
    content = tmp[1]
    return
}
