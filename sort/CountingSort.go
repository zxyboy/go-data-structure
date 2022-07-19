package main

import (
	"fmt"
	"sort"
)

// 计数排序
func countingSort(arr []int) {

	countMap := make(map[int]int, 11)

	for _, v := range arr {
		countMap[v] = countMap[v] + 1
	}
	// 对map中的key进行排序
	var keys []int
	for k, _ := range countMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	idx := 0
	for _, k := range keys {
		count := countMap[k]
		for count > 0 {
			arr[idx] = k
			idx++
			count--
		}
	}
}

func main() {
	arr := []int{3, 5, 7, 1, 2, 10, 6}
	//i := partition(arr, 0, len(arr)-1)
	//fmt.Printf("left point = %d \n", i)
	countingSort(arr)
	for _, e := range arr {
		fmt.Printf("%d ", e)
	}
}
