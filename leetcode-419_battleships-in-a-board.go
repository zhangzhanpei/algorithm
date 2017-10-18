/**
 * 给定一个由'.'和'X'组成的二维矩阵, '.'表示海水, 'X'表示战舰, 求战舰数量(战舰不相邻)
 */
package main

import "fmt"

func main() {
    
}

func countBattleships(board [][]byte) int {
    ret := 0
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            //如果X的左边和上边都是海水, 说明这是船头, 战舰数量+1
            if board[i][j] == 'X' && (i == 0 || board[i - 1][j] == '.') && (j == 0 || board[i][j - 1] == '.') {
                ret++
            }
        }
    }
    return ret
}
