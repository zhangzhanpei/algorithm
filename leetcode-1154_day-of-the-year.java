/**
 * 给定一个日期字符串date，判断日期是该年的第几天
 * 日期格式如 2019-08-28
 */
public int dayOfYear(String date) {
    int sum = 0;
    int[] days = {31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};
    // 先切分成年月日
    String[] tmp = date.split("-");
    int y = Integer.parseInt(tmp[0]);
    // 看是否闰年
    if ((y % 100 != 0 && y % 4 == 0) || y % 400 == 0) {
        days[1] = 29;
    }
    int m = Integer.parseInt(tmp[1]);
    int d = Integer.parseInt(tmp[2]);
    // 根据月份累加日期
    for (int i = 0; i < m - 1; i++) {
        sum += days[i];
    }
    return sum + d;
}
