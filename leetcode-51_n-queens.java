/**
 * N 皇后问题
 */
class Solution {
    // 返回结果
    private List<List<String>> ret = new ArrayList<>();

    // N 皇后问题
    public List<List<String>> solveNQueens(int n) {
        char[][] chessboard = new char[n][n];
        dfs(chessboard, 0, n);
        return ret;
    }

    public void dfs(char[][] chessboard, int row, int n) {
        // 已经超出棋盘最后一行(即n个棋子已摆放完毕)
        if (row == n) {
            // 打印当前摆放结果
            ret.add(printChessboard(chessboard));
            return;
        }

        // 一行内逐个列尝试
        for (int col = 0; col < n; col++) {
            if (isValid(chessboard, row, col, n)) {
                chessboard[row][col] = 'Q';
                // 当前行放下棋子后，继续尝试下一行
                dfs(chessboard, row + 1, n);
                // 下一行尝试完后，回到当前行，收起刚才放下的棋子，for 循环尝试当前行的下一列
                chessboard[row][col] = '.';
            }
        }
    }

    public boolean isValid(char[][] chessboard, int row, int col, int n) {
        // 同一列是否已有棋子
        for (int i = 0; i < row; i++) {
            if (chessboard[i][col] == 'Q') {
                return false;
            }
        }

        // 左上方斜线上是否有棋子
        for (int i = row - 1, j = col - 1; i >= 0 && j >= 0; i = i - 1, j = j - 1) {
            if (chessboard[i][j] == 'Q') {
                return false;
            }
        }

        // 右上方斜线上是否有棋子
        for (int i = row - 1, j = col + 1; i >= 0 && j < n; i = i - 1, j = j + 1) {
            if (chessboard[i][j] == 'Q') {
                return false;
            }
        }

        return true;
    }

    // 打印棋子摆放结果
    public List<String> printChessboard(char[][] chessboard) {
        List<String> cb = new ArrayList<>();
        for (int i = 0; i < chessboard.length; i++) {
            for (int j = 0; j < chessboard[0].length; j++) {
                if (chessboard[i][j] == '\0') {
                    chessboard[i][j] = '.';
                }
            }
            cb.add(new String(chessboard[i]));
        }
        return cb;
    }
}
