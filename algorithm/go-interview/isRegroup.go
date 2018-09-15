package main

import (
	"fmt"
	"strings"
)
//判断两个给定的字符串排序后是否一致
//只需要一次循环遍历s1中的字符在s2是否都存在即可。
func isRegroup(s1,s2 string) bool {

	for _,v := range s1 {
		if strings.Count(s1,string(v)) != strings.Count(s2,string(v)) {
			return false
		}
	}
	return true
}


//交替打印数字和字母
func main() {
	v:=isRegroup("adncd", "adcnd")
	fmt.Println(v)
}