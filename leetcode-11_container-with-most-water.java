/**
 * 给定一个 int 数组，任意两个元素表示桶壁的高度，求最多能装多少水
 */
public int maxArea(int[] height) {
    int water = 0;
    // 水量 = 底 * 高，其中底为两个元素的下标差，高为两个元素中的小者(短板)
    // 这里的元素从两边夹逼，是为了让底边由大到小
    int i = 0, j = height.length - 1;
    while (i < j) {
        water = Math.max(water, (j - i) * Math.min(height[i], height[j]));
        // 这里去掉低的桶壁，看能不能换个高的
        if (height[i] == height[j]) {
            i++;
            j--;
        } else if (height[i] < height[j]) {
            i++;
        } else {
            j--;
        }
    }
    return water;
}
