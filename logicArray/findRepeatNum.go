package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	aa := findRepeatNumber(nums)
	fmt.Println(aa)
}

func findRepeatNumber(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		index := int(math.Abs(float64(nums[i])))
		if index == n {
			index -= n
		}
		if nums[index] > 0 {
			nums[index] = -nums[index]
		} else if nums[index] == 0 {
			nums[index] = -n
		} else {
			return int(math.Abs(float64(nums[i])))
		}
	}
	return -1
}
