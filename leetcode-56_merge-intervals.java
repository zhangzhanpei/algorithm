/**
 * 给定一个区间集合，合并所有区间的交集
 */
public static int[][] merge(int[][] intervals) {
    // 如果只有一组区间，直接返回
    if (intervals.length <= 1) {
        return intervals;
    }

    // 按每个区间的第一个元素排序
    Arrays.sort(intervals, new Comparator<int[]>() {
        @Override
        public int compare(int[] o1, int[] o2) {
            return o1[0] - o2[0];
        }
    });

    List<int[]> ret = new ArrayList<>();
    int[] tmp = intervals[0];
    ret.add(tmp);
    for (int[] item : intervals) {
        if (item[0] <= tmp[1]) { // 后一个区间的起点小于前一个区间的终点，说明存在交集
            tmp[1] = Math.max(item[1], tmp[1]); // 两个区间融合
        } else {
            tmp = item;
            ret.add(tmp);
        }
    }

    return ret.toArray(new int[ret.size()][]);
}
