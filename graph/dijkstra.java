import java.util.Arrays;

class Solution {
    public static void main(String[] args) {
        // 给定一个邻接矩阵表示城市间的距离，如0号城市到0号城市距离为0，0号到1号距离为2...
        // 使用 99999 代表无穷，即城市间不直接连通
        int[][] city = {
                {0, 2, 12, 99999, 3, 99999},
                {99999, 0, 3, 6, 99999, 99999},
                {99999, 99999, 0, 99999, 5, 99999},
                {99999, 99999, 4, 0, 99999, 15},
                {99999, 99999, 99999, 2, 0, 4},
                {99999, 99999, 99999, 99999, 99999, 0},
        };

        System.out.println(Arrays.toString(dijkstra(city)));
    }

    /**
     * Dijkstra 算法，单源最短路径（一个点到其他点的最短距离）
     * @param city 城市间的距离
     * @return
     */
    public static int[] dijkstra(int[][] city) {
        // 这里默认计算第0个城市到其他城市的最短距离，初始化距离数组
        int[] distance = new int[city.length];
        for (int i = 0; i < city.length; i++) {
            distance[i] = city[0][i];
        }

        // 标记第0号城市到第0号城市已确定是最短距离
        boolean[] mark = new boolean[city.length];
        mark[0] = true;

        // 计算0号城市到其他城市的最短距离
        for (int i = 1; i < city.length; i++) {
            // 从距离数组中的未确定城市中找到距离0号城市最近的城市
            int min = Integer.MAX_VALUE;
            int idx = -1;
            for (int j = 1; j < distance.length; j++) {
                if (!mark[j] && distance[j] < min) {
                    min = distance[j];
                    idx = j;
                }
            }

            // 已确定 idx 号城市的最短距离
            mark[idx] = true;

            // 以 idx 号城市为中间点，看能否使0号城市到其他城市距离变短
            for (int k = 0; k < city.length; k++) {
                if (city[idx][k] < 99999) {
                    if (distance[k] > distance[idx] + city[idx][k]) {
                        distance[k] = distance[idx] + city[idx][k];
                    }
                }
            }
        }

        return distance;
    }
}
