/**
 * 给定一个int数组，去掉极值求平均数
 */
class Solution {
    public double average(int[] salary) {
        int min = salary[0];
        int max = salary[0];
        int sum = 0;
        for (int i : salary) {
            sum += i;
            min = Math.min(i, min);
            max = Math.max(i, max);
        }
        
        double total = sum - min - max;

        return total / (salary.length - 2);
    }
}
