/**
 * 给定两个string数组(words1, words2)和一个单词关系数组，求两个string数组是否相似
 */
package main

import "fmt"

func main() {
    words1 := []string{"great", "acting", "skills"}
    words2 := []string{"fine", "drama", "talent"}
    pairs := [][]string{{"great", "fine"}, {"acting","drama"}, {"skills","talent"}}
    fmt.Println(areSentencesSimilar(words1, words2, pairs))
}

func areSentencesSimilar(words1 []string, words2 []string, pairs [][]string) bool {
    if len(words1) != len(words2) { //相似必须单词数量一致
        return false
    }
    m := map[string][]string{}
    for _, v := range pairs { //遍历关系数组将相似单词写入map
        if _, ok := m[v[0]]; ok {
            m[v[0]] = append(m[v[0]], v[1])
        } else {
            m[v[0]] = []string{v[1]}
        }

        if _, ok2 := m[v[1]]; ok2 {
            m[v[1]] = append(m[v[1]], v[0])
        } else {
            m[v[1]] = []string{v[0]}
        }
    }

    for i := 0; i < len(words1); i++ {
        if words1[i] == words2[i] { //单词相同时直接就是相似
            continue
        }
        if stringArr, ok := m[words1[i]]; ok && inArray(stringArr, words2[i]) { //根据相似数组进行比较
            continue
        }
        return false
    }
    return true
}

func inArray(arr []string, value string) bool {
    for _, v := range arr {
        if v == value {
            return true
        }
    }
    return false
}
