/**
 * 实现一个LRU缓存[最近最少使用], 该缓存包括get和set方法
 * 复杂度O(n)
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
    Cap int
    Linklist *list.List
}

//初始化lru缓存
func Constructor(capacity int) LRUCache {
    return LRUCache{Cap:capacity, Linklist: list.New()}
}

//传入key到lru缓存中获取value
func (this *LRUCache) Get(key int) int {
    for item := this.Linklist.Front(); item != nil; item = item.Next() {
        node := item.Value.(Node)
        //如果在cache中找到key, 则返回相应的value, 并将该节点移到双链表头部
        //将使用到的key节点移到头部, 那么最近最少使用的节点会自然地落到双链表尾部方便删除腾出容量
        if node.K == key {
            this.Linklist.MoveToFront(item)
            return node.V
        }
    }
    //如果key不在cache中, 返回-1
    return -1
}

//插入key-value键值对到lru缓存中
func (this *LRUCache) Put(key int, value int) {
    for item := this.Linklist.Front(); item != nil; item = item.Next() {
        node := item.Value.(Node)
        if node.K == key { //如果key已经在cache中则覆盖新值, 并将该节点移动到双链表头部
            node.V = value
            item.Value = node
            this.Linklist.MoveToFront(item)
            return
        }
    }

    //如果key不在cache中, 则插入到双链表头部, 并检查双链表长度, 如果超出cache初始化时定义的容量, 则删除双链表尾部一个节点[因为尾部节点是最近最少使用的]
    node := Node{K: key, V: value}
    this.Linklist.PushFront(node)
    if this.Linklist.Len() > this.Cap {
        this.Linklist.Remove(this.Linklist.Back())
    }
}
