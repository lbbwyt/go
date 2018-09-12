package main

import (
	"fmt"
	"math/rand"
	"time"
)


//洗牌算法
func Shuffle(vals []int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(vals) > 0 {
		fmt.Println(len(vals) )
		n := len(vals)
		randIndex := r.Intn(n)
		//将得到的下标对应的元素和最后一个数交换
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
		fmt.Printf("%v",vals)
	}
}



func main() {
	a := []int{1,2,3,4,5,6,7,8,9}
	Shuffle(a)
	fmt.Printf("%v",a)
}