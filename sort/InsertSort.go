package main

import "fmt"

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

func main() {
	arr := []int{3, 5, 7, 1, 2, 10}
	insertSort2(arr)
	for _, e := range arr {
		fmt.Printf("%d ", e)
	}
}
