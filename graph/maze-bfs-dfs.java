import java.util.LinkedList;
import java.util.Queue;

/**
 * 给定一个邻接矩阵表示迷宫，为1的位置不可通过，从（0, 0）处出发，抵达（3, 2）处营救小狗，求最短路径
 */
class Solution {
    private int min = Integer.MAX_VALUE;
    private int[][] next = {{0, -1}, {1, 0}, {0, 1}, {-1, 0}};
    private boolean[][] mark = new boolean[5][4];
    private int[][] map = {
            {0, 0, 1, 0},
            {0, 0, 0, 0},
            {0, 0, 1, 0},
            {0, 1, 0, 0},
            {0, 0, 0, 1},
    };

    // 存储节点坐标，便于放入队列
    class Node {
        public int x, y, step;
        Node(int x, int y, int step) {
            this.x = x;
            this.y = y;
            this.step = step;
        }
    }

    /**
     * 广度优先搜索
     */
    public void bfs() {
        // 起点加入队列
        Queue<Node> queue = new LinkedList<>();
        queue.add(new Node(0, 0, 0));
        while (!queue.isEmpty()) {
            Node node = queue.poll();
            // 取出节点，如果是目的地，更新最小步数
            if (node.x == 3 && node.y == 2) {
                if (node.step < min) {
                    min = node.step;
                }
                continue;
            }

            // 标记该节点已访问
            mark[node.x][node.y] = true;

            // 上下左右四个方向的节点加入队列
            for (int k = 0; k < 4; k++) {
                int nx = node.x + next[k][0];
                int ny = node.y + next[k][1];
                if (nx < 0 || nx >= 5 || ny < 0 || ny >= 4) {
                    continue;
                }
                if (!mark[nx][ny] && map[nx][ny] == 0) {
                    queue.add(new Node(nx, ny, node.step + 1));
                }
            }
        }
    }

    /**
     * 深度优先搜索
     */
    public void dfs(int x, int y, int step) {
        if (x == 3 && y == 2) {
            if (step < min) {
                min = step;
            }
            return;
        }

        // 上下左右四个方向继续尝试
        for (int k = 0; k < 4; k++) {
            int nx = x + next[k][0];
            int ny = y + next[k][1];
            if (nx < 0 || nx >= 5 || ny < 0 || ny >= 4) {
                continue;
            }
            if (!mark[nx][ny] && map[nx][ny] == 0) {
                // 标记该节点已访问
                mark[nx][ny] = true;
                // 以下一个节点为开始继续搜索
                dfs(nx, ny, step + 1);
                // 回溯时取消刚尝试的节点
                mark[nx][ny] = false;
            }
        }
    }
}
