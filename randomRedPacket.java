/**
 * 插板法分配红包金额
 * 100 块分成 20 个红包，金额精确到分即 10000 分，生成 19 个 0~10000 之间的随机位置(插板)，插板间隔就是每个红包的金额
 */
public static List<Double> randomRedPacket(double money, int count) {
    // 总金额精确到分
    int totalAmount = (int) (money * 100);
    // 利用 TreeMap 存储插板位置，这样插板位置就是有序的
    Map<Integer, Integer> map = new TreeMap<>();
    // 存储随机红包金额
    List<Double> redPacket = new ArrayList<>();
    Random random = new Random();
    for (int i = 0; i < count - 1; i++) {
        int idx = random.nextInt(totalAmount);
        // 如果插板位置冲突了则重新生成随机位置
        while (map.containsKey(idx)) {
            idx = random.nextInt(totalAmount);
        }
        map.put(idx, idx);
    }
    // 处理最后一个红包
    map.put(totalAmount, totalAmount);

    // 前一个插板位置
    int prev = 0;
    for (Integer k : map.keySet()) {
        // 插板间隔就是每个红包的金额
        redPacket.add((double) (k - prev) / 100);
        prev = k;
    }
    return redPacket;
}
