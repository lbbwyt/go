package main

import (
	"fmt"
)

//一趟快速排序的算法是：
//1）设置两个变量i、j，排序开始的时候：i=0，j=N-1；
//2）以第一个数组元素作为关键数据，赋值给key，即 key=A[0]；
//3）从j开始向前搜索，即由后开始向前搜索（j -- ），找到第一个小于key的值A[j]，A[i]与A[j]交换；
//4）从i开始向后搜索，即由前开始向后搜索（i ++ ），找到第一个大于key的A[i]，A[i]与A[j]交换；
//5）重复第3、4、5步，直到 I=J； (3,4步是在程序中没找到时候j=j-1，i=i+1，直至找到为止。找到并交换的时候i， j指针位置不变。
//另外当i=j这过程一定正好是i+或j-完成的最后令循环结束。）

func quickSort(array []int, low int, high int) {

	if low < high {
		pos := partition(array, low, high)
		quickSort(array, low, pos-1)
		quickSort(array, pos+1, high)
	}
}

func partition(array []int, low, high int) int {
	key := array[low]
	tmpLow := low
	tmpHigh := high
	for {
		//		从tmpHigh开始向前搜索，即由后开始向前搜索（tmpHigh -- ），找到第一个小于key的值A[tmpHigh]，
		for array[tmpHigh] > key {
			tmpHigh--
		}
		//找到大于key的元素，该元素的位置一定是low到tmpHigh+1之间。因为array[tmpHigh+1]必定大于key
		for array[tmpLow] <= key && tmpLow < tmpHigh {
			tmpLow++
		}

		if tmpLow >= tmpHigh {
			break
		}
		// swap(array[tmpLow], array[tmpHigh])
		array[tmpLow], array[tmpHigh] = array[tmpHigh], array[tmpLow]
		//		fmt.Println(array)
	}
	array[tmpLow], array[low] = array[low], array[tmpLow]
	return tmpLow
}

func main() {
	var sortArray = []int{3, 41, 24, 76, 11, 45, 3, 3, 64, 21, 69, 19, 36}
	fmt.Println(sortArray)
	quickSort(sortArray, 0, len(sortArray)-1)
	fmt.Println(sortArray)
}
