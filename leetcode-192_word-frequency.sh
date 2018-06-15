# 给定一个文本，统计各单词出现次数
for i in `cat words.txt`; do echo $i; done | sort | uniq -c | sort -nrk1 | awk '{print $2, $1}'
