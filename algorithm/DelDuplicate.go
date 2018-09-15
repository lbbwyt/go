package main

import (
	"fmt"
	"reflect"
	"sort"
)

//交替打印数字和字母
func main() {
	b:=[]string {"a","b","a","c"}
	sort.Strings(b)
	a := DelDuplicate(b)
	fmt.Printf("%v",a)
}

func DelDuplicate(a interface{}) (ret []interface{}) {
	va := reflect.ValueOf(a)
	for i:=0;i<va.Len();i++ {
		if i>0 && reflect.DeepEqual(va.Index(i-1).Interface(),va.Index(i).Interface()) {
			continue
		}
		ret = append(ret, va.Index(i).Interface())
	}

	return  ret
}