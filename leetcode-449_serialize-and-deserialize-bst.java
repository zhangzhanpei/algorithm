/**
 * 二叉有序树的序列化与反序列化
 */
public class Codec {

    // 前序遍历一棵二叉树，节点值用逗号拼接成一个字符串
    public String serialize(TreeNode root) {
        String ret = "";
        if (root == null) {
            return ret;
        }
        ret += String.valueOf(root.val) + ",";
        ret += serialize(root.left);
        ret += serialize(root.right);
        return ret;
    }

    // 字符串以逗号分隔并构建成一个二叉有序树
    public TreeNode deserialize(String data) {
        if (data == null || "".equals(data)) {
            return null;
        }

        String[] values = data.split(",");

        TreeNode root = new TreeNode(Integer.valueOf(values[0]));
        for (int i = 1; i < values.length; i++) {
            TreeNode curr = root;
            int tmp = Integer.valueOf(values[i]);
            while (true) {
                if (tmp < curr.val ) {
                    if (curr.left != null) {
                        curr = curr.left;
                    } else {
                        curr.left = new TreeNode(tmp);
                        break;
                    }
                } else {
                    if (curr.right != null) {
                        curr = curr.right;
                    } else {
                        curr.right = new TreeNode(tmp);
                        break;
                    }
                }
            }
        }
        return root;
    }
}
