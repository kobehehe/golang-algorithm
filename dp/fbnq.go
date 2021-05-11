package main

import "fmt"

/**
非波那契数列
*/
func main() {
	n := getNums3(8)
	fmt.Println(n)
}

//递归
//func getNum(n int) int {
//	if n ==0 {
//		return 0
//	}
//	if n ==1 {
//		return 1
//	}
//	lMap := make(map[int]int)
//	if v,ok := lMap[n]; ok {
//		return v
//	}
//	k := getNum(n-1) + getNum(n-2)
//	lMap[n] = k
//	return k
//}

//迭代
func getNum2(n int) int {
	if n <= 1 {
		return n
	}

	arr := make([]int, n+1)
	arr[0] = 0
	arr[1] = 1
	for i := 2; i < n+1; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n]
}

//空间复杂度为O(1)的迭代

func getNums3(n int) int {
	if n <= 1 {
		return n
	}
	prev := 0
	curr := 1

	for i := 2; i <= n; i++ {
		sum := prev + curr
		prev = curr
		curr = sum
	}
	return curr
}
