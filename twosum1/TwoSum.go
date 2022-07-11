package main

import "fmt"

// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那 两个 整数，并返回它们的数组下标。
//
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//
// 你可以按任意顺序返回答案。
//
// 示例 1：
//
// 输入：nums = [2,7,11,15], target = 9
// 输出：[0,1]
// 解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
// 示例 2：
//
// 输入：nums = [3,2,4], target = 6
// 输出：[1,2]
// 示例 3：
//
// 输入：nums = [3,3], target = 6
// 输出：[0,1]
// 提示：
//
// 2 <= nums.length <= 104
// -109 <= nums[i] <= 109
// -109 <= target <= 109
// 只会存在一个有效答案
// 进阶：你可以想出一个时间复杂度小于 O(n2) 的算法吗？
//
// Related Topics
// 数组
// 哈希表
//
// 👍 14770
// 👎 0
func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	ints := twoSum1(nums, target)
	fmt.Println(ints)
	sum2 := twoSum2(nums, target)
	fmt.Println(sum2)

}

// 两数之和
func twoSum1(nums []int, target int) []int {

	m := make(map[int]int, len(nums))

	for i, num := range nums {
		diff := target - num
		if v, ok := m[diff]; ok {
			return []int{v, i}
		} else {
			m[num] = i
		}
	}
	return []int{0}
}

func twoSum2(nums []int, target int) []int {

	for i := 0; i < len(nums)-1; i++ {

		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{0}
}
