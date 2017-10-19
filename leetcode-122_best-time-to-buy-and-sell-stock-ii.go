//多次买进卖出, 求总利润
func maxProfit(prices []int) int {
    sum := 0
    for i := 0; i < len(prices) - 1; i++ {
        if (prices[i + 1] > prices[i]) {
            sum += prices[i + 1] - prices[i]
        }
    }
    return sum
}
