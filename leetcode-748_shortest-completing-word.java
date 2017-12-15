/**
 * 给定一个目标字符串和一个字符串数组，从数组中找到一个最短的字符串，这个字符串中的字符能组成目标字符串
 * 大小写忽略
 */
public String shortestCompletingWord(String licensePlate, String[] words) {
    licensePlate = licensePlate.toLowerCase(); //同一使用小写
    int[] target = new int[26];
    for (int i = 0; i < licensePlate.length(); i++) { //统计目标字符串各个字符的数量
        if (Character.isLetter(licensePlate.charAt(i))) {
            int idx = licensePlate.charAt(i) - 'a';
            target[idx]++;
        }
    }
    int min = Integer.MAX_VALUE;
    String ret = null;
    for (int j = 0; j < words.length; j++) { //从数组中找到符合条件的最短字符串
        int[] tmp = target.clone();
        if (isMatch(words[j].toLowerCase(), tmp)) {
            if (words[j].length() < min) {
                min = words[j].length();
                ret = words[j];
            }
        }
    }
    return ret;
}

public boolean isMatch(String a, int[] target) {
    for (int i = 0; i < a.length(); i++) { //逐个字符减去目标字符串的字符计数
        if (Character.isLetter(a.charAt(i))) {
            int idx = a.charAt(i) - 'a';
            target[idx]--;
        }
    }
    for (int j = 0; j < 26; j++) { //如果有字符没减完说明无法组成目标字符串
        if (target[j] > 0) {
            return false;
        }
    }
    return true;
}
