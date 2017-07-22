//给定一个整数n, 判断n是否为4的多次方
//Good solution without good explanation,it's easy to find that power of 4 numbers have those 3 common features.
//First,greater than 0.
//Second,only have one '1' bit in their binary notation,so we use x&(x-1) to delete the lowest '1',and if then it becomes 0,it prove that there is only one '1' bit.
//Third,the only '1' bit should be locate at the odd location,for example,16.It's binary is 00010000.So we can use '0x55555555' to check if the '1' bit is in the right place.With this thought we can code it out easily!
package main

import "fmt"

func main() {
    fmt.Println(isPowerOfFour(64));
}

func isPowerOfFour(n int) bool {
    return num > 0 && (num & (num - 1)) == 0 && (num & 0x55555555) != 0;
}
