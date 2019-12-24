/**
 * 给定一个字符串，求最长回文子串
 * dp[i][j] 代表下标 i 到 j 的字串是否回文，如果 dp[i][j] 是回文的，那么字串 dp[i + 1][j - 1] 也是回文的
 */
public static String longestPalindrome(String s) {
    if (s.isEmpty()) {
        return s;
    }
    int len = s.length();
    boolean[][] dp = new boolean[len][len];
    int left = 0, right = 0;
    // 这里的 i 必须是有大到小，因为下面会用到 dp[i + 1] 的值
    for (int i = len - 2; i >= 0; i--) {
        // 单个字符的字串肯定是回文的
        dp[i][i] = true;
        for (int j = i + 1; j < len; j++) {
            // 这里 j - i < 3 包括两种情况(aa, aba)
            // 如果 j - i >= 3 说明是 4 个以上字符，通过字串 dp[i + 1][j - 1] 来判断是否回文
            dp[i][j] = s.charAt(i) == s.charAt(j) && (j - i < 3 || dp[i + 1][j - 1]);
            // 如果是回文字串，记录是否更大的长度
            if (dp[i][j] && right - left < j - i) {
                left = i;
                right = j;
            }
        }
    }
    return s.substring(left, right + 1);
}
