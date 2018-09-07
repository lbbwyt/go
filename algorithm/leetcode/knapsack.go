package main

import (
	"fmt"
)

//背包问题
//其中 dp[i][j] 表示前 i 件物品体积不超过 j 的情况下能达到的最大价值。
//第 i 件物品没添加到背包，总体积不超过 j 的前 i 件物品的最大价值就是总体积不超过 j 的前 i-1 件物品的最大价值，
//dp[i][j] = dp[i-1][j]。
//第 i 件物品添加到背包中，dp[i][j] = dp[i-1][j-w] + v。

//容量为 W 的背包, N为物品的件数
func knapsack(W, N int, weights []int, values []int) int {
	//构造二维数组
	var dp [][]int
	for i := 0; i < N+1; i++ {
		tmp := make([]int, W+1)
		dp = append(dp, tmp)
	}

	for i := 1; i <= N; i++ {
		//		物品的属性，体积 w 和价值 v
		w := weights[i-1]
		v := values[i-1]
		for j := 1; j <= W; j++ {

			if j >= w {
				//第 i 件物品添加到背包中
				//dp[i][j] 表示前 i 件物品体积不超过 j 的情况下能达到的最大价值。
				dp[i][j] = Max(dp[i-1][j], dp[i-1][j-w]+v)
			} else {
				//第 i 件物品没添加到背包
				dp[i][j] = dp[i-1][j]
			}
		}

	}
	return dp[N][W]
}

func Max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func main() {
	a := []int{1, 3, 4, 7, 6}
	b := []int{1, 3, 4, 7, 6}
	fmt.Println(knapsack(8, len(a), a, b))
}
