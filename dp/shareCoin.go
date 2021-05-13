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
	target := 9
	c := []int{2, 5}

	aa := getNum(target, c)
	fmt.Println(aa)
}

//动态规划来一波

func getNum(target int, c []int) int {
	if target < 0 {
		return -1
	}
	if target == 0 {
		return 0
	}

	dp := make([]int, target+1)

	for i := 0; i < target+1; i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for i := 1; i <= target; i++ {
		for _, coin := range c {
			if i >= coin && dp[i-coin] != math.MaxInt32 {
				dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-coin]+1)))
			}
		}
	}
	res := dp[target]
	if dp[target] == math.MaxInt32 {
		res = -1
	}
	return res
}

func dfs(target int, c []int, memo []int) int {
	if target == 0 {
		return 0
	}
	if memo[target] != -1 {
		return memo[target]
	}
	minCoins := math.MaxInt32
	for i := 0; i < len(c); i++ {
		if target-c[i] < 0 {
			continue
		}
		subMinCoins := dfs(target-c[i], c, memo)
		if subMinCoins == -1 {
			continue
		}
		minCoins = int(math.Min(float64(minCoins), float64(subMinCoins+1)))
	}
	re := minCoins
	if minCoins == math.MaxInt32 {
		re = -1
	}
	memo[target] = re
	return re
}
