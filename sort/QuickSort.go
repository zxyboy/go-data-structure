package main

import "fmt"

// QuickSort 快速排序
func QuickSort(arr []int, left, right int) {
	if right <= left {
		return
	}
	pivot := partition(arr, left, right)
	// 递归基准值左边
	QuickSort(arr, left, pivot-1)

	// 递归基准值右边
	QuickSort(arr, pivot+1, right)

}

func partition(arr []int, left, right int) int {

	l := left
	// 取最右边的值为基准值
	pivot := right
	r := pivot - 1
	// 3 4 7
	// 左指针小于右指针则循环，即：左指针大于等于有指针，则退出
	for l <= r {
		// 左边：找到大于基准值的下标
		for arr[l] < arr[pivot] {
			l++
		}
		// 右边： 找到小于基准值的下标
		for r >= 0 && arr[r] > arr[pivot] {
			r--
		}

		// 如果右边的下标大于左边下标，则交换值
		if l < r {
			swap(arr, l, r)
		}
	}
	swap(arr, l, pivot)
	return l

}

// Swap 交换切片中left和right下标位置的值
func swap(arr []int, left, right int) {
	tmp := arr[left]
	arr[left] = arr[right]
	arr[right] = tmp
}

func main() {
	arr := []int{3, 5, 7, 1, 2, 10, 6}
	//i := partition(arr, 0, len(arr)-1)
	//fmt.Printf("left point = %d \n", i)
	QuickSort(arr, 0, len(arr)-1)
	for _, e := range arr {
		fmt.Printf("%d ", e)
	}
}
