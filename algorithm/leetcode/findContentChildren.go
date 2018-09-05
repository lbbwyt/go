package main

import (
	"fmt"
	"sort"
)

//题目描述：每个孩子都有一个满足度，每个饼干都有一个大小，只有饼干的大小大于等于一个孩子的满足度，
//该孩子才会获得满足。求解最多可以获得满足的孩子数量。

//Input: [1,2], [1,2,3]
//Output: 2

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var i, j int
	i = 0
	j = 0
	for i < len(g) && j < len(s) {
		if g[i] < s[j] {
			i++
		}
		j++
	}
	return i
}
func main() {
	g := []int{1, 3}
	s := []int{1, 2, 3, 4}
	fmt.Printf("%v", findContentChildren(g, s))
}
