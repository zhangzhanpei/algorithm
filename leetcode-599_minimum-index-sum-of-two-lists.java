/**
 * 给定两个数组求交集，且交集元素的下标之和最小
 */
public static String[] findRestaurant(String[] list1, String[] list2) {
    int sum = Integer.MAX_VALUE;
    List<String> ret = new ArrayList<>();
    HashMap<String, Integer> map1 = new HashMap<>();
    //第一个数组进map
    for (int i = 0; i < list1.length; i++) map1.put(list1[i], i);
    for (int j = 0; j < list2.length; j++) {
        if (map1.containsKey(list2[j])) { //交集元素
            if (map1.get(list2[j]) + j < sum) {
                ret.clear();
                ret.add(list2[j]);
                sum = map1.get(list2[j]) + j;
                continue;
            }
            if (map1.get(list2[j]) + j == sum) {
                ret.add(list2[j]);
            }
        }
    }
    return ret.toArray(new String[ret.size()]);
}
