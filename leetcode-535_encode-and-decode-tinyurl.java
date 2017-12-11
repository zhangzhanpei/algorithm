/**
 * 短链接生成与解析
 */
public class Codec {
    //存储urls
    private ArrayList<String> urls = new ArrayList<>();

    //将一个长连接编码成短链接返回
    public String encode(String longUrl) {
        this.urls.add(longUrl);
        int index = this.urls.size() - 1;
        return "http://tinyurl.com/" + Integer.toString(index);
    }

    //短链接还原成长连接
    public String decode(String shortUrl) {
        int lastIdx = shortUrl.lastIndexOf("/") + 1;
        String key = shortUrl.substring(lastIdx);
        return this.urls.get(Integer.parseInt(key));
    }
}
