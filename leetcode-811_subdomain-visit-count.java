/**
 * 给定一个字符串数组表示每个域名被访问的次数，统计各级域名访问的次数
 */
public List<String> subdomainVisits(String[] cpdomains) {
    List<String> ret = new ArrayList<>();
    Map<String, Integer> map = new HashMap<>();
    for (String url : cpdomains) {
        // 元素格式如 ["9001 discuss.leetcode.com"]，切出次数与域名
        int count = Integer.parseInt(url.split(" ")[0]);
        String[] tmp = url.split(" ")[1].split("\\.");
        for (int i = 0; i < tmp.length; i++) {
            String domain = "";
            // 拼出各级域名
            for (int j = i; j < tmp.length; j++) {
                domain += "." + tmp[j];
            }
            domain = domain.substring(1);

            // 有则次数相加，无则初始化次数
            if (map.containsKey(domain)) {
                map.put(domain, map.get(domain) + count);
            } else {
                map.put(domain, count);
            }
        }
    }

    // 将结果加入列表返回
    map.forEach((k, v) -> ret.add(v + " " + k));
    return ret;
}
