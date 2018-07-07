/**
 * 最小生成树，一个连通所有顶点且权值之和最小的图
 */
public static void prim(int[][] city) {
    // 以第 0 个城市为起点，初始化各城市距离
    int[] minCost = Arrays.copyOf(city[0], city[0].length);

    int sum = 0;
    for (int i = 0; i < city.length; i++) {
        // 查找能连通的最近的城市
        int min = -1;
        for (int j = 0; j < city.length; j++) {
            if (minCost[j] > 0 && minCost[j] < 99999) {
                if (min == -1 || minCost[j] < minCost[min]) {
                    min = j;
                }
            }
        }

        // 如果找不到了就退出
        if (min == -1) {
            break;
        }
        
        // 加上该条边的权值
        sum += minCost[min];
        // 权值置 0 则查找 min 城市的所有出边时不会往回走
        minCost[min] = 0;
        
        // 已加入的顶点看作一个整体，现新加入 min 顶点，可能有更短的距离或新路线能到达原整体能到达或不能到达的城市
        // 这里需要更新最短距离数组
        for (int k = 0; k < city.length; k++) {
            if (city[min][k] < minCost[k]) {
                minCost[k] = city[min][k];
            }
        }
    }

    // 输出最小生成树的权值总和
    System.out.println(sum);
}
