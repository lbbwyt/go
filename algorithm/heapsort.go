package main

import (
	"fmt"
)

//大顶堆：arr[i] >= arr[2i+1] && arr[i] >= arr[2i+2]
func main() {
	var arr = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	sort(arr)
	fmt.Printf("%v", arr)
}

func sort(arr []int) {
	//从最后一个非叶子结点从下至上，从右至左调整结构
	//叶结点自然不用调整，最后一个非叶子结点 arr.length/2-1
	//建堆就是调整堆的过程
	for i := len(arr)/2 - 1; i >= 0; i-- {
		adjustHeap(arr, i, len(arr))
	}
	//2.调整堆结构+交换堆顶元素与末尾元素
	for j := len(arr) - 1; j > 0; j-- {
		swap(arr, 0, j)       //将堆顶元素与末尾元素进行交换
		adjustHeap(arr, 0, j) //重新对堆进行调整，堆的大小随j的值减小而减小
	}
}

//调整大顶堆

func adjustHeap(arr []int, i int, length int) {
	temp := arr[i]                              //先取出当前元素i
	for k := i*2 + 1; k < length; k = k*2 + 1 { //从i结点的左子结点开始，也就是2i+1处开始
		if (k+1) < length && arr[k] < arr[k+1] { //如果左子结点小于右子结点，k指向右子结点
			k++
		}
		if arr[k] > arr[i] { //如果子节点大于父节点，将子节点赋值给父节点，（不用进行交换）
			arr[i] = arr[k]
			i = k
		} else {
			break
		}

	}

	arr[i] = temp //将temp值放到最终的位置
}

//交换元素
func swap(arr []int, a int, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}
