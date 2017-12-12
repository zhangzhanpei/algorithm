/**
 * 实现一个单词查找树，包括insert、search、startsWith等方法
 * 一个Trie树从根节点到任意一个叶子节点经过的字符连起来就是一个单词
 */

//定义Trie树的节点
class TrieNode {
    public char val; //存储字符
    public boolean isWord; //叶子节点标记单词结束
    public TrieNode[] children = new TrieNode[26];

    TrieNode() {}

    TrieNode(char c) {
        TrieNode t = new TrieNode();
        t.val = c;
    }
}

//定义一个Trie树
public class Trie {
    private TrieNode root;

    public Trie() {
        this.root = new TrieNode();
        this.root.val = ' ';
    }

    //插入单词
    public void insert(String word) {
        TrieNode node = this.root; //从根节点开始
        for (int i = 0; i < word.length(); i++) {
            char c = word.charAt(i);
            if (node.children[c - 'a'] == null) { //如果根节点的儿子节点中没有这个字符
                node.children[c - 'a'] = new TrieNode(c); //则新建节点并加入子节点数组中
            }
            node = node.children[c - 'a']; //从子节点继续往下处理
        }
        node.isWord = true; //最后叶子节点标记单词完毕
    }

    //搜索单词
    public boolean search(String word) {
        TrieNode node = this.root; //从根节点开始
        for (int i = 0; i < word.length(); i++) {
            char c = word.charAt(i);
            if (node.children[c - 'a'] == null) return false; //如果有一个字符找不到说明Trie树中没有这个单词
            node = node.children[c - 'a'];
        }
        return node.isWord; //如果单词的字符都比较完毕了，如果到了叶子节点说明刚好是一个单词，否则这个单词只是前缀
    }

    //检查Trie树中是否有以prefix开头的单词
    public boolean startsWith(String prefix) {
        TrieNode node = this.root; //从根节点开始
        for (int i = 0; i < prefix.length(); i++) {
            char c = prefix.charAt(i);
            if (node.children[c - 'a'] == null) return false; //如果有一个字符找不到说明Trie树中没有这个前缀
            node = node.children[c - 'a'];
        }
        return true;
    }
}
