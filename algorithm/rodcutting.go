package main

import (
	"fmt"
)

func main() {
	length := 10
	price := []int{0, 1, 5, 8, 9, 17, 17, 17, 20, 24, 30}
	//fmt.Print(cutRodRec(price, length), "\n")
	fmt.Print(cutRodDp(price, length), "\n")
}

//求最大值
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}

}

//递归求解
func cutRodRec(price []int, length int) int {
	if length == 0 {
		return 0
	}
	q := -1
	for i := 1; i <= length; i++ {
		q = max(q, price[i]+cutRodRec(price, length-i))
	}
	return q
}

//动态规划求解，
func cutRodDp(price []int, length int) int {
	r := make([]int, length+1) // a.k.a the memoization array
	r[0] = 0                   // cost of 0 length rod is 0
	//先求出r[1],然后求r[2],有两种情况即r[1] + p[1]
	//    和r[0] + p[2]

	for j := 1; j <= length; j++ { // for each length (subproblem)
		q := -1
		for i := 1; i <= j; i++ {
			q = max(q, price[i]+r[j-i]) // avoiding recursive call
		}
		//
		r[j] = q
	}

	return r[length]
}
