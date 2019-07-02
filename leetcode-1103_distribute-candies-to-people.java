/**
 * 分糖果
 * 给的糖果数和人数，从第一个人开始分，分发数量依次加一，到最后一个人后如有剩余，回到第一个人继续递增
 */
public int[] distributeCandies(int candies, int num_people) {
    int[] ret = new int[num_people];
    int n = 0;
    while (candies > 0) {
        for (int i = 0; i < num_people; i++) {
            n++;
            if (candies >= n) {
                ret[i] += n;
                candies -= n;
            } else {
                ret[i] += candies;
                candies = 0;
            }
        }
    }
    return ret;
}
