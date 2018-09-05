package main

import (
	"fmt"
)

//最长递增子序列
//在子序列中，当下标 ix > iy 时，Six > Siy，称子序列为原序列的一个 递增子序列 。
func lengthOfLIS(nums []int) int {
	length := len(nums)
	dp := make([]int, length)
	for i := 0; i < length; i++ {
		max := 1
		//求nums[i]之前的最长子序列
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				max = Max(max, dp[j]+1)
			}
		}
		dp[i] = max
	}

	//遍历dp求最大值，即为最长递增子序列的长度
	result := 0
	for ret := 0; ret < length; ret++ {
		result = Max(result, dp[ret])
	}
	return result

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
	fmt.Println(lengthOfLIS(a))
}
