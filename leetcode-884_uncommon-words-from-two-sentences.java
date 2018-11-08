/**
 * 给定两个句子，找出只出现一次的单词
 */
class Solution {
    public String[] uncommonFromSentences(String A, String B) {
        Map<String, Integer> map = new HashMap<>();
        String[] tmpA = A.split(" ");
        for (String s : tmpA) {
            map.put(s, map.getOrDefault(s, 0) + 1);
        }
        String[] tmpB = B.split(" ");
        for (String s : tmpB) {
            map.put(s, map.getOrDefault(s, 0) + 1);
        }
        List<String> list = new ArrayList<>();
        for(Map.Entry<String, Integer> entry : map.entrySet()) {
            if (entry.getValue() == 1) {
                list.add(entry.getKey());
            }
        }
        return list.toArray(new String[list.size()]);
    }
}
