package main

import "fmt"

// 快速排序
func QuickSort(arr []int, left, right int) {

}

func main() {
	arr := []int{3, 5, 7, 1, 2, 10}
	insertSort2(arr)
	for _, e := range arr {
		fmt.Printf("%d ", e)
	}
}
