/**
 * 给定一个正整数 n, 只能使用 copy 和 paste 操作, 求得到长度为 n 的字符串最少需要几步
 * n = 1, return 0
 * n = 2, return 2 //复制一次粘贴一次
 * n = 3, return 3 //复制一次粘贴两次
 * n = 4, return 4 //复制一次粘贴三次, 或复制一次粘贴一次再复制一次再粘贴一次
 * ...
 * dp[i]表示长度为i时最小步数, 则 dp[i] = min(dp[i], dp[j] + i / j)
 * n = 6 时, 复制一次粘贴一次共两步得到 AA, 6 / 2 = 3 即还需 3 步操作(复制一次, 粘贴两次)得到 AAAAAA
 * 或先得到 AAA, 6 / 3 = 2 即还需 2 步操作(复制一次, 粘贴一次)得到 AAAAAA
 */
public int minSteps(int n) {
    int[] dp = new int[n + 1];
    for (int i = 2; i <= n; i++) {
        dp[i] = i;
        for (int j = i - 1; j > 1; j--) {
            if (i % j == 0) {
                dp[i] = Math.min(dp[i], dp[j] + i / j);
            }
        }
    }
    return dp[n];
}
