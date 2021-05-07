package main

import (
	"fmt"
	"math"
	"sort"
)

//给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
//
//你可以认为每种硬币的数量是无限的。
//输入：coins = [1, 2, 5], amount = 11
//输出：3
//解释：11 = 5 + 5 + 1

func main() {
	coins := []int{3, 5}
	amount := 12
	//num := getMinCount(coins,amount) //贪心实现
	path := []int{}
	//res := [][]int{}

	maxNum := math.MaxInt32
	getAllRes(amount, coins, path, &maxNum)
	fmt.Println(maxNum)

}

func getNum(res [][]int) int {
	if len(res) == 0 {
		return -1
	}

	min := 0

	for i := 0; i < len(res); i++ {
		if len(res[i]) < len(res[min]) {
			min = i
		}
	}
	return len(res[min])
}

//回溯解决一下
func getAllRes(target int, coins []int, path []int, maxNum *int) {

	if target == 0 {
		if len(path) < *maxNum {
			*maxNum = len(path)
		}
		return
	}
	for i := 0; i < len(coins); i++ {
		if target-coins[i] < 0 {
			continue
		}
		path = append(path, coins[i])
		getAllRes(target-coins[i], coins, path, maxNum)
		path = path[:len(path)-1]
	}
}

//** 贪心实现有瑕疵 不能够回溯了
func getMinCount(coins []int, amount int) int {

	sort.Ints(coins)

	rest := amount
	num := 0
	for i := len(coins) - 1; i >= 0; i-- {
		cNum := rest / coins[i]
		rest = rest - cNum*coins[i]
		num += cNum
		if rest == 0 {
			return num
		}
	}

	return -1
}
