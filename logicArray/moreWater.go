package main

import "fmt"

/**
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器。

*/
func main() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	num := getMoreWater(arr)

	fmt.Println(num)
}

func getMoreWater(arr []int) int {
	left := 0
	right := len(arr) - 1
	square := 0
	y := 0
	for right >= left {

		x := right - left
		if arr[right] < arr[left] {
			y = arr[right]
		} else {
			y = arr[left]
		}
		area := x * y
		if area > square {
			square = area
		}

		if arr[left] <= arr[right] {
			left++
		} else {
			right--
		}
	}

	return square
}
