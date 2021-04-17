package main

import "fmt"

/**
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
*/
func main() {
	nums := []int{0, 1, 0, 3, 12}
	aa := removeZero(nums)
	fmt.Println(aa)
}

func removeZero(nums []int) []int {
	n := len(nums)
	if n < 1 {
		return nums
	}

	for i, j := 0, 0; i < n; i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return nums
}
