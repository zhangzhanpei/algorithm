/**
 * 给定一个ipv4地址，将其中的.转换成[.]
 */
class Solution {
    public static void main(String[] args) {
        String ip = "192.168.1.1";
        System.out.println(defangIPaddr(ip));
    }
    public static String defangIPaddr(String address) {
        return address.replaceAll("\\.", "[.]");
    }
}
