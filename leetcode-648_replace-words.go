/**
 * 给定一个string数组和一个字符串句子, 如果句子中的单词以数组中的单词开头, 则替换成该单词
 */
package main

import "fmt"
import "strings"

func main() {
    replaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery")
}

func replaceWords(dict []string, sentence string) string {
    words := strings.Split(sentence, " ")
    for _, d := range dict {
        for key, word := range words { //切割句子, 逐个单词检查开头字符
            if strings.Index(word, d) == 0 {
                words[key] = d
            }
        }
    }
    return strings.Join(words, " ")
}
