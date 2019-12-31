// kmp
public int strpos(String haystack, String needle) {
    if (haystack.length() < needle.length()) {
        return -1;
    }
    int[] next = getNext(needle);
    int j = 0;
    for (int i = 0; i < haystack.length(); i++) {
        while (j > 0 && haystack.charAt(i) != needle.charAt(j)) {
            j = next[j - 1];
        }
        if (haystack.charAt(i) == needle.charAt(j)) {
            j++;
        }
        if (j >= needle.length()) {
            return i - j + 1;
        }
    }

    return -1;
}

public int[] getNext(String needle) {
    int[] next = new int[needle.length()];
    int k = 0;
    for (int i = 1; i < needle.length(); i++) {
        while (k > 0 && needle.charAt(k) != needle.charAt(i)) {
            k = next[k - 1];
        }
        if (needle.charAt(k) == needle.charAt(i)) {
            k++;
        }
        next[i] = k;
    }
    return next;
}
