/*
 * 给定一个string数组，删除某些列后使得每列都是升序排序
 */
public static int minDeletionSize(String[] A) {
    int ret = 0;
    int cols = A[0].length();
    // 遍历每一列
    for (int i = 0; i < cols; i++) {
        int prev = 'a';
        // 按行遍历
        for (String str : A) {
            // 如果下一个字符小于前一个字符，说明不是升序，该列需要删除
            if (str.charAt(i) < prev) {
                ret++;
                break;
            } else {
                prev = str.charAt(i);
            }
        }
    }
    return ret;
}
