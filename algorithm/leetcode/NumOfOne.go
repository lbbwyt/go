package main

import (
	"fmt"
)

//输入一个整数，输出该数二进制表示中 1 的个数。
//n       : 10110100
//n-1     : 10110011
//n&(n-1) : 10110000

//负数在计算机中使用补码表示的，即反码进1，11111011表示8位的-5，
func NumOfOne(n int) int {
	count := 0

	for n != 0 {
		count++
		n &= n - 1
	}

	return count
}

func main() {
	fmt.Println(NumOfOne(-1))
}
