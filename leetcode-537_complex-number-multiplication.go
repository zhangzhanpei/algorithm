/**
 * 虚数乘法
 * (c+di)(e+fi) = (ce+df)+(cf+de)i
 */
package main

import "fmt"
import "strings"
import "strconv"

func main() {
    complexNumberMultiply("1+-1i", "1+-1i")
}

func complexNumberMultiply(a string, b string) string {
    index := strings.Index(a, "+")
    c, _ := strconv.Atoi(a[:index]) //字符串a的实部
    d, _ := strconv.Atoi(a[index + 1:len(a) - 1]) //字符串a的虚部
    index = strings.Index(b, "+")
    e, _ := strconv.Atoi(b[:index]) //字符串b的实部
    f, _ := strconv.Atoi(b[index + 1:len(b) - 1]) //字符串b的虚部
    real := c * e - d * f //求实部
    complex := c * f + d * e //求虚部
    return strconv.Itoa(real) + "+" + strconv.Itoa(complex) + "i"
}
