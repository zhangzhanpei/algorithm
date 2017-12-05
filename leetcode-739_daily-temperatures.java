/**
 * 给定一个int数组表示每日天气，求每天需要几天温度才变暖
 */
public int[] dailyTemperatures(int[] temperatures) {
    int[] ret = new int[temperatures.length];
    Stack<Integer> s = new Stack<>();
    for (int i = 0; i < temperatures.length; i++) {
        while (!s.isEmpty() && temperatures[i] > temperatures[s.peek()]) { //每天向后找到第一个大于当天气温的元素
            int idx = s.pop(); //当天出栈
            ret[idx] = i - idx; //记录天数
        }
        s.push(i); //如果找不到后面变暖，则天数不记录，默认为0
    }
    return ret;
}
