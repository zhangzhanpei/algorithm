/**
 * 实现一个LRU缓存[最近最少使用], 该缓存包括get和set方法
 * 使用hash table辅助, 实现O(1)复杂度
 */
package main

import "fmt"
import "container/list"

func main() {
    lru := Constructor(2)
    lru.Put(1, 1)
    lru.Put(2, 2)
    fmt.Println(lru.Get(1))
    lru.Put(3, 3)
    fmt.Println(lru.Get(2))
    lru.Put(4, 4)
    fmt.Println(lru.Get(1))
    fmt.Println(lru.Get(3))
    fmt.Println(lru.Get(4))
}

//key-value结构
type Node struct {
    K int
    V int
}

//lru缓存结构
type LRUCache struct {
    Cap int //lru缓存的容量
    Linklist *list.List //双向链表
    Container map[int]*list.Element //map, 保存的值是链表节点
}

//初始化lru缓存
func Constructor(capacity int) LRUCache {
    return LRUCache{Cap:capacity, Linklist: list.New(), Container: map[int]*list.Element{}}
}

//传入key到lru缓存中获取value
func (this *LRUCache) Get(key int) int {
    //从map中直接按key取, 如果取到链表节点, 则将该节点移到链表头部
    if item, ok := this.Container[key]; ok {
        this.Linklist.MoveToFront(item)
        node := item.Value.(Node)
        return node.V
    }
    return -1
}

//插入key-value键值对到lru缓存中
func (this *LRUCache) Put(key int, value int) {
    //如果key已经在缓存中, 则用新值覆盖
    if item, ok := this.Container[key]; ok {
        node := item.Value.(Node)
        node.V = value
        item.Value = node
        this.Linklist.MoveToFront(item) //链表节点赋新值并移到头部
        this.Container[key] = item //存入map中
        return
    }

    //如果key不在cache中, 则插入到双链表头部
    node := Node{K: key, V: value}
    element := this.Linklist.PushFront(node)
    this.Container[key] = element
    if this.Linklist.Len() > this.Cap { //检查双链表长度, 如果超出cache初始化时定义的容量, 则删除双链表尾部一个节点[因为尾部节点是最近最少使用的]
        lastElement := this.Linklist.Back();
        this.Linklist.Remove(lastElement)
        lastNode := lastElement.Value.(Node)
        delete(this.Container, lastNode.K) //取到尾部被删除节点的key到map中删除相应键值对
    }
}
