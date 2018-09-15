package main

import (
	"fmt"
	"strings"
)
//判断字符串中字符是否全都不同
//使用golang内置方法strings.Count,可以用来判断在一个字符串中包含的另外一个字符串的数量。
func isUniqueString(s string) bool {

	if len(s)>128{
		return  false
	}
	for _,v := range s {
		if v > 127 {
			return false
		}
		if strings.Count(s,string(v)) > 1 {
			return false
		}
	}
	return true
}


//交替打印数字和字母
func main() {
	v:=isUniqueString("adncd")
	fmt.Println(v)
}