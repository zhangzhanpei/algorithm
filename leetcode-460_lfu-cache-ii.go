/**
 * 实现一个LFU缓存[最近最低频率使用], 该缓存包括get和set方法
 * 使用hash table辅助存储, 复杂度O(1)
 */
package main

import "fmt"
import "container/list"

func main() {
    lfu := Constructor(2)
    lfu.Put(1, 1)
    lfu.Put(2, 2)
    fmt.Println(lfu.Get(1))
    lfu.Put(3, 3)
    fmt.Println(lfu.Get(2))
    fmt.Println(lfu.Get(3))
    lfu.Put(4, 4)
    fmt.Println(lfu.Get(1))
    fmt.Println(lfu.Get(3))
    fmt.Println(lfu.Get(4))
}

//key-value结构
type Node struct {
    K int //键
    V int //值
    F int //访问次数
}

//cache结构
type LFUCache struct {
    Cap int
    Linklist *list.List
    Container map[int]*list.Element
}

//构造函数
func Constructor(capacity int) LFUCache {
    return LFUCache{Cap:capacity, Linklist: list.New(), Container: map[int]*list.Element{}}
}

//从cache中获取指定key的值
func (this *LFUCache) Get(key int) int {
    if item, ok := this.Container[key]; ok {
        node := item.Value.(Node)
        node.F++
        item.Value = node
        this.Swap(item)
        return node.V
    }
    return -1
}

//写入key-value到cache中
func (this *LFUCache) Put(key int, value int) {
    //当缓存容量为0时无法写入
    if this.Cap == 0 {
        return
    }
    if item, ok := this.Container[key]; ok {
        node := item.Value.(Node)
        node.V = value
        node.F++
        item.Value = node
        this.Swap(item)
        return
    }

    //如果key不在cache中, 先判断cache是否还有容量, 如果没有容量需要先删除双链表尾部一个节点,然后将新节点插入到双链表尾部, 并向前调整节点位置
    if this.Linklist.Len() >= this.Cap {
        lastItem := this.Linklist.Back();
        lastNode := lastItem.Value.(Node)
        this.Linklist.Remove(lastItem)
        delete(this.Container, lastNode.K)
    }
    node := Node{K: key, V: value, F: 0}
    element := this.Linklist.PushBack(node)
    this.Swap(element)
    this.Container[key] = element
}

//向前调整节点位置
func (this *LFUCache) Swap(element *list.Element) {
    prevElement := element.Prev()
    for prevElement != nil {
        prevNode := prevElement.Value.(Node)
        node := element.Value.(Node)
        if node.F < prevNode.F {
            return
        }
        this.Linklist.MoveBefore(element, prevElement)
        prevElement = prevElement.Prev()
    }
}
