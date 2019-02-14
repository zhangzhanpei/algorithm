/*
 * 给定一个int数组和一个整数，两者相加，返回和的数组
 */
public List<Integer> addToArrayForm(int[] A, int K) {
    List<Integer> ret = new ArrayList<>();
    int j = A.length - 1;
    while (j >= 0 || K > 0) {
        if (j >= 0) {
            K += A[j];
        }
        ret.add(K % 10);
        K /= 10;
        j--;
    }
    Collections.reverse(ret);
    return ret;
}
