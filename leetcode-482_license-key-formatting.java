/**
 * 给定一个字符串和一个长度，按长度格式化license
 * Input: S = "5F3Z-2e-9-w", K = 4, Output: "5F3Z-2E9W"
 * Input: S = "2-5g-3-J", K = 2,    Output: "2-5G-3J" 
 */
public static String licenseKeyFormatting(String S, int K) {
    S = S.replaceAll("-", "").toUpperCase(); //去除-并转为大写
    StringBuffer sb = new StringBuffer();
    int j = 0;
    for (int i = S.length() - 1; i >= 0; i--) { //从后往前取字符
        j++;
        if (i > 0 && j % K == 0) { //每取K个字符则拼接一个-
            sb.append(S.charAt(i)).append("-");
        } else {
            sb.append(S.charAt(i));
        }
    }
    return sb.reverse().toString();
}
