package main

import (
	"fmt"
)

//最长公共子序列
//定义一个二维数组 dp 用来存储最长公共子序列的长度，其中 dp[i][j] 表示 S1 的前 i 个字符与 S2 的前 j 个字符最长公共子序列的长度

func lengthOfLCS(nums1 []int, nums2 []int) int {
	n1 := len(nums1)
	n2 := len(nums2)
	//	dp := [n1+1][n2+1]int  就会显示no-constant array bound n,直接报错
	//  make([][]int, n1+1, n2+1) 二维的空数组，子数组长度为0 ,之后就会报数组越界的错误

	var dp [][]int
	for i := 0; i < n1+1; i++ {
		tmp := make([]int, n2+1)
		dp = append(dp, tmp)
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if nums1[i-1] == nums2[j-1] {
				//	相等
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = Max(dp[i-1][j], dp[i][j-1])
			}
		}

	}
	return dp[n1][n2]
}

func Max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func main() {
	a := []int{1, 3, 4, 74, 6}
	b := []int{1, 3, 4, 74, 6}
	fmt.Println(lengthOfLCS(a, b))
}
