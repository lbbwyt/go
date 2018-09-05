package main

import (
	"fmt"
)

//题目描述：有 N 阶楼梯，每次可以上一阶或者两阶，求有多少种上楼梯的方法。
/*定义一个数组 dp 存储上楼梯的方法数（为了方便讨论，数组下标从 1 开始），dp[i] 表示走到第 i 个楼梯的方法数目。
第 i 个楼梯可以从第 i-1 和 i-2 个楼梯再走一步到达，走到第 i 个楼梯的方法数为走到第 i-1 和第 i-2 个楼梯的方法数之和*/

func climbstairs(n int) int {
	if n <= 2 {
		return n
	}
	return climbstairs(n-1) + climbstairs(n-2)
}

func climb(n int) int {
	if n <= 2 {
		return n
	}
	pre1, pre2 := 1, 2
	for i := 2; i < n; i++ {
		temp := pre1 + pre2
		pre1 = pre2
		pre2 = temp
	}
	return pre2

}

func main() {
	fmt.Println(climb(7))
}
