/**
 * 给定一个int二维数组表示多个用户多个银行的存款，求总存款最高的用户有多少存款
 */
public int maximumWealth(int[][] accounts) {
    int max = 0;
    for (int[] banks : accounts) {
        int tmp = 0;
        for (int i : banks) {
            tmp += i;
        }
        max = Math.max(max, tmp);
    }
    return max;
}
