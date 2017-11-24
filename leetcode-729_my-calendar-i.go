/**
 * 设计一个日历表可预约事件日期，保证所有事件日期不重叠
 */
package main

import "fmt"

func main() {
    obj := Constructor()
    fmt.Println(obj.Book(10, 20))
    fmt.Println(obj.Book(15, 25))
    fmt.Println(obj.Book(20, 30))
}

type MyCalendar struct {
    date [][]int
}


func Constructor() MyCalendar {
    return MyCalendar{date: [][]int{}}
}


func (this *MyCalendar) Book(start int, end int) bool {
    for _, v := range this.date {
        if start < v[1] && end > v[0] {
            return false
        }
    }
    this.date = append(this.date, []int{start, end})
    return true
}
