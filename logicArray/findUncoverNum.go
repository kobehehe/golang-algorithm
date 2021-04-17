package main

import (
	"fmt"
	"math"
)

/**
一个数组里的数组 满足 1 =< nums[i] <= n ,n为数组的长度 要求求出数组中 1-n中 所消失的数字
要求时间复杂度O(N) 空间复杂度O(1)
*/
func main() {
	arr := []int{1, 2, 3, 2, 3, 6, 7, 8} //把这个数组输出4，5就对了
	aa := findDisappearNum(arr)
	fmt.Println(aa)
}

func findDisappearNum(arr []int) []int {
	var res []int
	n := len(arr)
	for i := 0; i < n; i++ {
		index := int(math.Abs(float64(arr[i]))) - 1
		if arr[index] > 0 {
			arr[index] = -arr[index]
		}
	}

	for i := 0; i < n; i++ {
		if arr[i] > 0 {
			res = append(res, i+1)
		}
	}

	return res
}
