import java.util.Collections;
import java.util.PriorityQueue;

/**
 * 给定一组 int 整数，求中数
 * 如果元素数量是偶数，中数为中间两个元素的均值
 * 如果元素数量是奇数，中数为中间那个元素
 */
class MedianFinder {

    private PriorityQueue<Integer> small = new PriorityQueue<>(Collections.reverseOrder()); // 大根堆保存小的一半元素，堆顶为小的那一半中最大那个元素
    private PriorityQueue<Integer> large = new PriorityQueue<>(); // 小根堆保存大的一半元素，堆顶为大的那一半中最小那个元素

    public void addNum(int num) {
        large.add(num); // 先将 num 放入大元素堆
        small.add(large.poll()); // 取出大元素堆最小的放到小元素堆
        if (large.size() < small.size()) { // 如果此时大元素堆的元素数量小于小元素堆数量，从小元素堆取出最大的元素放入大元素堆，保证大元素堆的数量等于小元素堆或比小元素堆多一个元素
            large.add(small.poll());
        }
    }

    public double findMedian() {
        return large.size() > small.size() ? large.peek() : (large.peek() + small.peek()) / 2.0;
    }

    public static void main(String[] args) {
        MedianFinder mf = new MedianFinder();
        mf.addNum(3);
        mf.addNum(2);
        mf.addNum(1);
        System.out.println(mf.findMedian());
    }
}
