/**
 * 给定一组坐标点，判断是否连成一条直线
 * 逐个坐标判断与前一个坐标的斜率一致则为直线
 */
class Solution {
    public static void main(String[] args) {
        int[][] coordinates = {
                {1,2},
                {2,3},
                {3,4},
                {4,5},
                {5,6},
                {6,7}
        };
        System.out.println(checkStraightLine(coordinates));
    }

    public static boolean checkStraightLine(int[][] coordinates) {
        if (coordinates.length <= 2) {
            return true;
        }

        for (int i = 2; i < coordinates.length; i++) {
            int x0 = coordinates[i - 1][0] - coordinates[i - 2][0];
            int y0 = coordinates[i - 1][1] - coordinates[i - 2][1];
            int x1 = coordinates[i][0] - coordinates[i - 1][0];
            int y1 = coordinates[i][1] - coordinates[i - 1][1];
            // 斜率相等即 y1/x1 = y0/x0，为防止x0或x1为0，转换为乘法比较
            if (x0 * y1 != x1 * y0) {
                return false;
            }
        }

        return true;
    }
}
