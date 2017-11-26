/**
 * 给定一个int二维数组表示一张图片，给定一个开始坐标和一个新颜色，从开始坐标处填充图片
 */
package main

import "fmt"

func main() {
    fmt.Println(floodFill([][]int{
        {0, 0, 0},
        {0, 0, 0},
    }, 0, 0, 2))
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    oldColor := image[sr][sc]
    if oldColor == newColor { //如果开始坐标处的颜色与新颜色一致则不填充
        return image
    }
    fill(image, sr, sc, oldColor, newColor)
    return image
}

func fill(image [][]int, sr, sc, oldColor, newColor int) {
    //记录开始坐标处的旧颜色
    if sr < 0 || sr >= len(image) || sc < 0 || sc >= len(image[0]) || image[sr][sc] != oldColor {
        return
    }
    image[sr][sc] = newColor
    //上下左右四个方向查找相邻的旧颜色进行填充
    fill(image, sr, sc + 1, oldColor, newColor)
    fill(image, sr + 1, sc, oldColor, newColor)
    fill(image, sr, sc - 1, oldColor, newColor)
    fill(image, sr - 1, sc, oldColor, newColor)
}
