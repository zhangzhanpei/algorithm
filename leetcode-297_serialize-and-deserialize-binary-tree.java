import java.util.LinkedList;
import java.util.Queue;

/**
 * 给定一个二叉树，实现序列化和反序列化方法
 */
class Solution {
    // 序列化
    public String serialize(TreeNode root) {
        if (root == null) {
            return "";
        }
        Queue<TreeNode> queue = new LinkedList<>();
        queue.offer(root);
        StringBuilder stringBuilder = new StringBuilder(root.val + ",");
        while (queue.size() > 0) {
            root = queue.poll();
            if (root.left != null) {
                queue.offer(root.left);
                stringBuilder.append(root.left.val + ",");
            } else {
                stringBuilder.append("null,");
            }

            if (root.right != null) {
                queue.offer(root.right);
                stringBuilder.append(root.right.val + ",");
            } else {
                stringBuilder.append("null,");
            }
        }

        return stringBuilder.toString();
    }

    // 反序列化
    public TreeNode deserialize(String data) {
        if (data.isEmpty()) {
            return null;
        }
        String[] nums = data.split(",");
        Queue<TreeNode> queue = new LinkedList<>();
        TreeNode root = new TreeNode(Integer.valueOf(nums[0]));
        queue.offer(root);
        int i = 0;
        while (i < nums.length - 1) {
            TreeNode node = queue.poll();
            i++;
            if (nums[i].equals("null")) {
                node.left = null;
            } else {
                node.left = new TreeNode(Integer.valueOf(nums[i]));
                queue.offer(node.left);
            }
            i++;
            if (nums[i].equals("null")) {
                node.right = null;
            } else {
                node.right = new TreeNode(Integer.valueOf(nums[i]));
                queue.offer(node.right);
            }
        }
        return root;
    }
}
