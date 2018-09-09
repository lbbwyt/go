package main

import (
	"fmt"
	"strings"
)

//Input:
//"I am a student."

//Output:
//"student. a am I"

//先旋转每个单词，再旋转整个字符串。

func ReverseSentence(str string) string {
	strs := strings.Split(str, " ")
	length := len(strs)
	i := 0
	j := length - 1
	for i <= j {
		strs[i], strs[j] = strs[j], strs[i]
		i++
		j--
	}
	return strings.Join(strs, " ")
}

func reverseWord(str string) string {
	chars := []rune(str)
	n := len(chars)
	var i, j int = 0, 0
	for j <= n {
		if j == n || chars[j] == ' ' {
			reverse(chars, i, j-1)
			i = j + 1
		}
		j++
	}
	return string(chars)

}

func reverse(c []rune, i int, j int) {
	for i < j {
		swap(c, i, j)
		i++
		j--
	}
}

func swap(c []rune, i int, j int) {
	c[i], c[j] = c[j], c[i]
}

func main() {
	fmt.Println(ReverseSentence(reverseWord("I am a student")))
}
