//给定一个整数n, 判断n是否为3的多次方
//因为3的20次方大于int的最大值2147483647, 所以n的最大值为3的19次方1162261467
package main

import "fmt"

func main() {
    fmt.Println(isPowerOfThree(27));
}

func isPowerOfThree(n int) bool {
    return (n > 0 && 1162261467 % n == 0);
}
