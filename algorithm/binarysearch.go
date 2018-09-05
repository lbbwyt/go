package main

import (
	"fmt"
	"sort"
)

func main() {
	//必须是有序的
	arr := []int{6, 3, 3, 5, 8, 7, 9}
	sort.Ints(arr)
	num := iterBinarySearch(arr, 6, 0, len(arr)-1)
	fmt.Println(num)
}

func binarySearch(array []int, target int, lowIndex int, highIndex int) int {

	if highIndex < lowIndex {
		return -1
	}
	mid := int((lowIndex + highIndex) / 2)
	if array[mid] > target {
		return binarySearch(array, target, lowIndex, mid)
	} else if array[mid] < target {
		return binarySearch(array, target, mid+1, highIndex)
	} else {
		return mid
	}
}

func iterBinarySearch(array []int, target int, lowIndex int, highIndex int) int {
	startindex := lowIndex
	endindex := highIndex
	var mid int
	for startindex < endindex {
		mid = int((startindex + endindex) / 2)
		if array[mid] < target {
			startindex = mid
		} else if array[mid] > target {
			endindex = mid
		} else {
			return mid
		}
	}

	return -1
}
