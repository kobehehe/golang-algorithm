package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	aa := removeDuplicates3(nums)
	fmt.Println(aa)
}

/**
这个升级难度 每个元素最多出现两次
*/

func removeDuplicates3(nums []int) int {
	n := len(nums)
	slow := 1
	for fast := 2; fast < n; fast++ {
		if nums[fast] != nums[slow-1] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}

/**
直接赋值的算法 i是快指针 j是慢指针
*/
func removeDuplicates1(nums []int) int {
	n := len(nums)
	j := 0
	for i := 1; i < n; i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}

//自己想的 删除效率不高。 不如直接赋值效率高
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
