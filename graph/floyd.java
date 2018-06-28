class Solution {
    public static void main(String[] args) {
        // 给定一个邻接矩阵表示城市间的距离，如0号城市到0号城市距离为0，0号到1号距离为2...
        // 使用 99999 代表无穷，即城市间不直接连通
        int[][] city = {
                {0, 2, 6, 4},
                {99999, 0, 3, 99999},
                {7, 99999, 0, 1},
                {5, 99999, 12, 0},
        };

        System.out.println(floyd(city, 3, 1));
        System.out.println(floyd(city, 2, 3));
    }

    /**
     * Floyd 算法，多源最短路径（点与点直接的所有距离）
     * @param city 城市间的距离
     * @param origin 起点城市
     * @param target 终点城市
     * @return
     */
    public static int floyd(int[][] city, int origin, int target) {
        for (int k = 0; k < city.length; k++) {
            for (int i = 0; i < city.length; i++) {
                for (int j = 0; j < city.length; j++) {
                    // 如果经过中间点 k 使得城市 i 到城市 j 的距离更短，那么 i 和 j 的最短距离缩小为 i 到 k 的距离加上 k 到 j 的距离
                    if (city[i][k] + city[k][j] < city[i][j]) {
                        city[i][j] = city[i][k] + city[k][j];
                    }
                }
            }
        }
        return city[origin][target];
    }
}
