/**
 * 给定一个字符串数组，按包含相同字符的字符串分组
 */
public List<List<String>> groupAnagrams(String[] strs) {
    Map<String, List<String>> map = new HashMap<>();
    for (String s : strs) {
        // 计算字符串中每个字符的数量
        int[] tmp = new int[26];
        for (int i = 0; i < s.length(); i++) {
            int idx = s.charAt(i) - 'a';
            tmp[idx]++;
        }
        // 转成字符串作为 hash key
        String key = Arrays.toString(tmp);
        List<String> list = new ArrayList<>();
        if (map.containsKey(key)) {
            list = map.get(key);
        }
        // 加入列表中
        list.add(s);
        map.put(key, list);
    }
    return new ArrayList<>(map.values());
}
