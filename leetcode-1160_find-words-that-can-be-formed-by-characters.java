/**
 * 给定一个字符串数组words和一个字符串chars，判断数组中字符串可以由chars中的字符组成的长度之和
 */
public static int countCharacters(String[] words, String chars) {
    int sum = 0;
    int[] charArr = new int[26];
    // 将字符串分拆成字符和词频写入一个数组
    for (Character c : chars.toCharArray()) {
        int idx = c - 'a';
        charArr[idx]++;
    }
    // 逐个单词判断
    for (String s : words) {
        if (s.length() > chars.length()) {
            continue;
        }
        int[] tmp = charArr.clone();
        boolean flag = true;
        for (Character c : s.toCharArray()) {
            int idx = c - 'a';
            if (tmp[idx] == 0) {
                flag = false;
                break;
            } else {
                tmp[idx]--;
            }
        }

        if (flag) {
            sum += s.length();
        }
    }

    return sum;
}
