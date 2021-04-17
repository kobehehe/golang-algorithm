package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	aa := removeDuplicates(nums)
	fmt.Println(aa)
}

func removeDuplicates(nums []int) []int {
	n := len(nums)
	for i, j := 1, 0; i < n; i++ {
		if nums[i] == nums[j] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		} else {
			j++
		}
		n = len(nums)
	}

	return nums
}
