//设计一个 stack 支持 push, pop, top 和 getMin 方法 
type MinStack struct {
    Data []int
}

//初始化一个 stack
func Constructor() MinStack {
    return MinStack{Data: []int{}}
}

//元素进栈
func (this *MinStack) Push(x int) {
    this.Data = append(this.Data, x)
}

//元素出栈
func (this *MinStack) Pop() {
    if len(this.Data) == 0 { //栈中没有元素时直接返回
        return
    } else if len(this.Data) == 1 { //只有一个元素时出栈后栈为空
        this.Data = []int{}
    } else if len(this.Data) > 1 {
        this.Data = this.Data[:len(this.Data) - 1]
    }
}

//取栈顶元素
func (this *MinStack) Top() int {
    if len(this.Data) == 0 {
        return 0
    } else {
        return this.Data[len(this.Data) - 1]
    }
}

//取栈中最小的元素
func (this *MinStack) GetMin() int {
    min := this.Data[0] //栈中只有一个元素时该元素就是最小元素
    for _, v := range this.Data {
        if v < min {
            min = v
        }
    }
    return min
}
