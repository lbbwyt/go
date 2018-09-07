package main

import (
	"fmt"
)

//编辑距离
//修改一个字符串成为另一个字符串，使得修改次数最少。一次修改操作包括：插入一个字符、删除一个字符、替换一个字符。

func minEditDistance(word1 string, word2 string) int {
	if word1 == "" && word2 == "" {
		return 0
	}
	w1 := []rune(word1)
	w2 := []rune(word2)

	m := len(w1)
	n := len(w2)
	//二维数组的初始化
	var dp [][]int
	for i := 0; i < m+1; i++ {
		temp := make([]int, n+1)
		dp = append(dp, temp)
	}
	//dp[i][j]表示word1中的前i个字符，修改为Word2的前j个字符，最少需要移动的次数。
	for i := 1; i < m+1; i++ {
		dp[i][0] = i
	}
	for i := 1; i < n+1; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if w1[i-1] == w2[j-1] {
				//相等
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = Min(dp[i-1][j-1], Min(dp[i-1][j], dp[i][j-1])) + 1
			}
		}
	}
	return dp[m][n]

}

func Max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func Min(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}

func main() {
	a := "intention"
	b := "execution"
	fmt.Println(minEditDistance(a, b))
}
