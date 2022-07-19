package main

import (
	"fmt"
)

// 插入排序
func insertSort(arr []int) {

	for i := 1; i < len(arr); i++ {

		for j := 0; j < i; j++ {
			// 找到位置
			if arr[i] < arr[j] {
				tmp := arr[i]
				// 移动后面的位置
				for x := i; x > j; x-- {
					arr[x] = arr[x-1]
				}
				// 插入数据
				arr[j] = tmp
				break
			}
		}

	}
}

// 插入排序
func insertSort2(arr []int) {

	for i := 1; i < len(arr); i++ {
		position := -1
		for j := 0; j < i; j++ {
			// 找到位置
			if arr[i] < arr[j] {
				position = j
				break
			}
		}
		if position != -1 {
			tmp := arr[i]
			// 移动后面的位置
			for x := i; x > position; x-- {
				arr[x] = arr[x-1]
			}
			// 插入数据
			arr[position] = tmp
		}

	}
}

// 插入排序
func insertSort3(arr []int) {
	for i := 1; i < len(arr); i++ {
		x := arr[i]
		// 默认插入位置为下标：0
		insertPosition := 0
		// for循环找出插入位置
		for j := i - 1; j >= 0; j-- {
			// 找到插入位置
			if x > arr[j] {
				insertPosition = j + 1
				break
			} else {
				arr[j+1] = arr[j]
			}
		}
		// 插入数据
		arr[insertPosition] = x
	}
}

// Swap 交换切片中left和right下标位置的值
func Swap(arr []int, left, right int) {
	tmp := arr[left]
	arr[left] = arr[right]
	arr[right] = tmp
}

func main() {
	arr := []int{3, 5, 7, 1, 2, 10}
	//insertSort2(arr)
	insertSort3(arr)
	for _, e := range arr {
		fmt.Printf("%d ", e)
	}
}
