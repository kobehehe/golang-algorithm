package main

import (
	"fmt"
	"sort"
)

//给定两个数组 一个数组代表饼干的大小 另一个数组代表小孩的胃口大小 怎么把饼干分给小孩 能够满足最多的小孩

//解题思路 利用贪心算法 每次都分配合适的 最后的结果就是合适的

func main() {

	g1 := []int{1, 2, 3} //胃口
	c1 := []int{2, 1}    //饼干

	num := shareCookie(g1, c1)
	fmt.Println(num)
}

func shareCookie(g1 []int, c1 []int) int {

	sort.Ints(g1)
	sort.Ints(c1)

	fmt.Println(g1, c1)

	i := 0
	j := 0

	for i < len(g1) && j < len(c1) {
		if c1[j] >= g1[i] {
			j++
		}
		i++
	}
	return j
}
