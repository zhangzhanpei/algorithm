/**
 * 最小堆方案
 */
public static void topK(int[] nums, int k) {
    // 优先队列，不提供 Comparator 时默认最小堆排序
    PriorityQueue<Integer> queue = new PriorityQueue<>(k);
    // 先让前 k 个元素进队列，此时第一个元素就是最小的
    for (int i = 0; i < k; i++) {
        queue.add(nums[i]);
    }
    // 遍历剩下的元素，如果比最小堆堆顶元素大，删除堆顶元素，新元素进堆
    for (int j = k; j < nums.length; j++) {
        if (queue.peek() < nums[j]) {
            queue.poll();
            queue.add(nums[j]);
        }
    }
    // 此时优先队列中剩下的就是最大的 k 个元素，打印输出
    for (int x = 0; x < k; x++) {
        System.out.println(queue.poll());
    }
}
