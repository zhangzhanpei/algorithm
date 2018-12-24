/**
 * 给定一个int数组共2N个元素，其中有N+1个唯一元素，找出重复N次的元素
 */
public static int repeatedNTimes(int[] A) {
    Map<Integer, Boolean> map = new HashMap<>();
    for (Integer i : A) {
        if (map.containsKey(i)) {
            return i;
        }
        map.put(i, true);
    }
    return 0;
}
