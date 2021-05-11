package main

import (
	"fmt"
	"math"
)

//给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
//
//你可以认为每种硬币的数量是无限的。
//输入：coins = [1, 2, 5], amount = 11
//输出：3
//解释：11 = 5 + 5 + 1
func main() {
	target := 11
	c := []int{3, 5}
	aa := dfs(target, c)
	fmt.Println(aa)
}

func dfs(target int, c []int) int {
	if target == 0 {
		return 0
	}
	minCoins := math.MaxInt32
	for i := 0; i < len(c); i++ {
		if target-c[i] < 0 {
			continue
		}
		subMincoins := dfs(target-c[i], c)
		if subMincoins == -1 {
			continue
		}
		minCoins = int(math.Min(float64(minCoins), float64(subMincoins+1)))
	}
	re := minCoins

	if minCoins == math.MaxInt32 {
		re = -1
	}
	return re

}
