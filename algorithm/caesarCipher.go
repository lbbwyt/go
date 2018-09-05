package main

import (
	"bytes"
	"flag"
	"fmt"
	"strings"
)

func main() {
	ciperKey := flag.Int("c", 1, "位移量")
	input := flag.String("i", "sdf", "输入字符串")
	flag.Parse()
	//是flag.Xxx()返回的值是变量的内存地址,要获取值时要通过在变量前加*(星号)获取.
	if *ciperKey > 26 || *ciperKey < -26 {
		flag.PrintDefaults()
	} else {
		fmt.Println(caesarCipher(*input, *ciperKey))
	}
}

func caesarCipher(input string, key int) string {
	var outputBuffer bytes.Buffer
	for _, r := range strings.ToLower(input) {
		newByte := int(r)
		if newByte >= 'a' && newByte <= 'z' {
			newByte += key
			if newByte > 'z' {
				newByte -= 26
			} else if newByte < 'a' {
				newByte += 26
			}
		}
		outputBuffer.WriteString(string(newByte))
	}
	return outputBuffer.String()
}
