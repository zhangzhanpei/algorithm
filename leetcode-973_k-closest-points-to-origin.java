import java.util.Comparator;
import java.util.PriorityQueue;

/**
 * 给定一个二维int数组，返回K个距[0,0]坐标最近的元素
 */
class Solution {
    public static int[][] kClosest(int[][] points, int K) {
        // 使用优先队列存储每个坐标距[0,0]的距离，优先队列默认为最小堆，这里取负则距离越大，负数越小越排在堆顶方便移除
        PriorityQueue<int[]> pq = new PriorityQueue<>(Comparator.comparing(arr -> -(arr[0] * arr[0] + arr[1] * arr[1])));
        for (int[] p : points) {
            pq.add(p);
            // 移除堆顶元素
            if (pq.size() > K) {
                pq.poll();
            }
        }

        // 剩下的就是负数较大，即正数距离较小的坐标
        int[][] ret = new int[K][2];
        while (K-- > 0) {
            ret[K] = pq.poll();
        }
        return ret;
    }
}
