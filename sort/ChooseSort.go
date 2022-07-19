package main

import "fmt"

// 选择排序
func selectSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		// 假设： 切片中第一个元素为最大
		maxIndex := 0
		max := arr[maxIndex]
		for j := 1; j < len(arr)-i; j++ {
			// 找到比max还大的元素
			if max <= arr[j] {
				max = arr[j]
				maxIndex = j
			}
			// 到达最后一轮结尾
			if j == len(arr)-1-i {
				tmp := arr[j]
				arr[j] = max
				arr[maxIndex] = tmp
			}
		}
	}
}

// 选择排序
func selectSort2(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		// 假设： 切片中第一个元素为最小
		minIndex := i
		min := arr[minIndex]

		for j := i + 1; j < len(arr)-1; j++ {
			// 找到比max还大的元素
			if min > arr[j] {
				min = arr[j]
				minIndex = j
			}
		}
		if i != minIndex {
			// 替换元素
			tmp := arr[i]
			arr[i] = min
			arr[minIndex] = tmp
		}
	}
}

func main() {
	arr := []int{3, 5, 7, 1, 2, 10}
	//selectSort(arr)
	selectSort2(arr)
	for _, e := range arr {
		fmt.Printf("%d ", e)
	}
}
