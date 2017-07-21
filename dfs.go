//小狗被困在_map[3][2]位置, 从_map[0][0]出发去救它, 求最短路径
package main

import "fmt"

func main() {
    dfs(0, 0, 0)
    fmt.Println(min)
}

var _map = [][]int{ //地图, 1表示障碍物
    []int{0, 0, 1, 0},
    []int{0, 0, 0, 0},
    []int{0, 0, 1, 0},
    []int{0, 1, 0, 0},
    []int{0, 0, 0, 1},
}
var book = [5][4]int{} //用于标记哪些位置被走过
var min = 99999 //最短路径, 一开始初始化一个比较大的数, 用于比较
var next = [][]int{
    []int{0, -1}, //向上走, x+0不变, y-1
    []int{1, 0}, //向右走, x+1, y不变
    []int{0, 1}, //向下走, x+0不变, y+1
    []int{-1, 0}, //向左走, x-1, y不变
}

func dfs(x, y, step int) {
    if (x == 3 && y == 2) { //到达小狗位置
        if (step < min) {
            min = step;
        }
        return
    }

    for k := 0; k < 4; k++ { //尝试4个方向走
        //计算下一个位置坐标
        nx := x + next[k][0]
        ny := y + next[k][1]
        if (nx < 0 || nx >= 5 || ny < 0 || ny >= 4) { //是否越过边界
            continue
        }
        if (_map[nx][ny] == 0 && book[nx][ny] == 0) { //如果下一个位置能走并且没走过
            book[nx][ny] = 1 //标记该位置走过
            dfs(nx, ny, step + 1) //尝试下一个位置
            book[nx][ny] = 0 //取回刚才尝试的位置
        }
    }
    return
}
