package main

import "fmt"

// 冒泡排序
func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		// 标志：一轮冒泡以后，是否出现交换位置的情况， 因为一轮冒泡后，如果没有进行数据交换，则说明数据已经有序
		flag := true
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] >= arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
				flag = false
			}
		}
		// 因为一轮冒泡后，如果没有进行数据交换，则说明数据已经有序，则不需要再进行冒泡处理
		if flag {
			break
		}
	}
}

func main() {
	arr := []int{3, 5, 7, 1, 2, 10}
	bubbleSort(arr)
	for _, e := range arr {
		fmt.Printf("%d ", e)
	}
}
