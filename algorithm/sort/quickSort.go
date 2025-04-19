package main

import (
	"fmt"
	"math/rand"
)

func partition(arr []int, low, high int) (int, int) {
	// 随机选择基准
	index := rand.Intn(high-low+1) + low
	arr[index], arr[low] = arr[low], arr[index]
	pivot := arr[low]

	// 三路快排
	i, j, k := low-1, high+1, low
	for k < j {
		if arr[k] < pivot {
			i++
			arr[i], arr[k] = arr[k], arr[i]
		} else if arr[k] > pivot {
			j--
			arr[j], arr[k] = arr[k], arr[j]
		} else {
			k++
		}
	}
	return i, j
}

func QuickSort(arr []int, low, high int) {
	// 如果，low<high则继续分治
	if low < high {
		// 获取上次递归基准索引
		l, r := partition(arr, low, high)
		// 左右分区，继续快速排序
		QuickSort(arr, low, l)
		QuickSort(arr, r, high)
	}
}

func main() {
	arr := []int{5, 2, 3, 1}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
