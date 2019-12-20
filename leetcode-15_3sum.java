/**
 * 给定一个 int 数组, 其中存在三个元素组成的三元组 (a, b, c)，使得 a + b + c = 0。找出所有这样的三元组
 */
public List<List<Integer>> threeSum(int[] nums) {
    List<List<Integer>> ret = new LinkedList<>();
    if (nums == null) {
        return ret;
    }
    // 升序排序
    Arrays.sort(nums);
    int left, right, tmp;
    for (int i = 0; i < nums.length; i++) {
        // 如果第一个元素就大于0，后面的也是大于0，和不可能为0
        if (nums[i] > 0) {
            break;
        }
        // 重复元素跳过
        if (i > 0 && nums[i] == nums[i - 1]) {
            continue;
        }
        left = i + 1;
        right = nums.length - 1;
        while (left < right) {
            tmp = nums[i] + nums[left] + nums[right];
            if (tmp > 0) {
                right--;
            } else if (tmp < 0) {
                left++;
            } else {
                List<Integer> list = Arrays.asList(nums[i], nums[left], nums[right]);
                ret.add(list);
                left++;
                right--;
                // 如果新的左元素与原来的相等，跳过
                while (left < right && nums[left] == nums[left - 1]) {
                    left++;
                }
                // 如果新的右元素与原来的相等，跳过
                while (left < right && nums[right] == nums[right + 1]) {
                    right--;
                }
            }
        }
    }
    return ret;
}
