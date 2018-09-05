package main

import (
	"fmt"
	"math"
)

//判断一个数是否为两个数的平方和，例如 5 = 12 + 22。
func main() {
	fmt.Println(judgeSquareSum(5))
}

//双指针
func judgeSquareSum(c int) (isOrNot bool) {
	i := 0
	//Sqrt returns the square root of x
	j := sqrt(c)
	for i < j {
		powSum := i*i + j*j
		if powSum == c {
			isOrNot = true
			return
		} else if powSum > c {
			j--
		} else {
			i++
		}
	}
	return
}

//求整型的平方根
func sqrt(c int) (result int) {
	temp := float64(c)
	fresult := math.Sqrt(temp)
	result = int(fresult)
	return
}
