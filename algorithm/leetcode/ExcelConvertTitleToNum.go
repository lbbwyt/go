package main

import (
	"fmt"
	"strconv"
)

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

//1 -> A
//2 -> B
//3 -> C
//...
//26 -> Z
//27 -> AA
//28 -> AB
func ExcelConvertToTitle(n int) string {
	if n == 0 {
		return ""
	}
	//	因为是从 1 开始计算的，而不是从 0 开始，因此需要对 n 执行 -1 操作。
	n--
	return ExcelConvertToTitle(n/26) + string(n%26+int('A'))
}

//将上面的反转
func ExcelConvertTitleToNum(s string) (res int) {
	res = 1
	//输入验证
	if s == "" {
		return 0
	}
	ss := []byte(s)
	length := len(ss)
	fmt.Println()
	for i := 0; i < length; i++ {
		res = i*26 + int(ss[i]+1)%int('A')
	}
	return
}

//10进制转2进制
func Dec2Binary(n int) string {
	if n < 0 {
		panic("必须大于0")
	}
	if n == 0 {
		return "0"
	}
	var s string = ""
	for q := n; q > 0; q = q / 2 {
		m := q % 2
		s = fmt.Sprintf("%v%v", m, s)
	}
	return s
}

var tenToAny map[int]string = map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f", 16: "g", 17: "h", 18: "i", 19: "j", 20: "k", 21: "l", 22: "m", 23: "n", 24: "o", 25: "p", 26: "q", 27: "r", 28: "s", 29: "t", 30: "u", 31: "v", 32: "w", 33: "x", 34: "y", 35: "z", 36: ":", 37: ";", 38: "<", 39: "=", 40: ">", 41: "?", 42: "@", 43: "[", 44: "]", 45: "^", 46: "_", 47: "{", 48: "|", 49: "}", 50: "A", 51: "B", 52: "C", 53: "D", 54: "E", 55: "F", 56: "G", 57: "H", 58: "I", 59: "J", 60: "K", 61: "L", 62: "M", 63: "N", 64: "O", 65: "P", 66: "Q", 67: "R", 68: "S", 69: "T", 70: "U", 71: "V", 72: "W", 73: "X", 74: "Y", 75: "Z"}

// 10进制转任意进制（如需大于76，需补齐tenToAny中的字符表示）
func decimalToAny(num, n int) string {
	new_num_str := ""
	var remainder int
	var remainder_string string
	for num != 0 {
		remainder = num % n
		if 76 > remainder && remainder > 9 {
			remainder_string = tenToAny[remainder]
		} else {
			remainder_string = strconv.Itoa(remainder)
		}
		new_num_str = remainder_string + new_num_str
		num = num / n
	}
	return new_num_str
}

func main() {
	fmt.Println(ExcelConvertTitleToNum("AB"))

	//	fmt.Println(ExcelConvertToTitle(27))
	//	//fmt.Println(string(27%26 + int('A')))

	//	fmt.Println(Dec2Binary(117))
	//	fmt.Println(decimalToAny(15, 16))
}
