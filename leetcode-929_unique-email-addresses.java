/**
 * 给定一个 email 地址数组，返回 unique email 数量
 * email 地址分为 local name 和 domain name 两部分，其中 local name 部分 . 号可去掉，+ 号后面的字符忽略
 */
class Solution {
    public int numUniqueEmails(String[] emails) {
        Set<String> set = new HashSet<>();
        for (String email : emails) {
            String[] tmp = email.split("\\+|@");
            String realEmail = tmp[0].replace(".", "") + "@" + tmp[tmp.length - 1];
            set.add(realEmail);
        }
        return set.size();
    }
}
