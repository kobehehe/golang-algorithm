package main

import "fmt"

func main() {
	arr1 := []int{1, 4, 2, 3, 6}
	target := 7
	num := len(arr1)
	arr1 = arr1[0 : len(arr1)-1]
	fmt.Println(arr1)
	return
	getArrayTarget(arr1, target, num)
}

func getArrayTarget(arr []int, target int, num int) {
	if num <= 2 {
		return
	}

}

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	var dfs func(start int, temp []int, sum int)

	dfs = func(start int, temp []int, sum int) {
		if sum >= target {
			if sum == target {
				newTmp := make([]int, len(temp))
				copy(newTmp, temp)
				res = append(res, newTmp)
			}
			return
		}
		for i := start; i < len(candidates); i++ {
			temp = append(temp, candidates[i])
			dfs(i, temp, sum+candidates[i])
			temp = temp[:len(temp)-1]
		}
	}
	dfs(0, []int{}, 0)
	return res
}

//func combinationSum(candidates []int, target int) (ans [][]int) {
//	comb := []int{}
//	var dfs func(target, idx int)
//	dfs = func(target, idx int) {
//		if idx == len(candidates) {
//			return
//		}
//		if target == 0 {
//			ans = append(ans, append([]int(nil), comb...))
//			return
//		}
//		// 直接跳过
//		dfs(target, idx+1)
//		// 选择当前数
//		if target-candidates[idx] >= 0 {
//			comb = append(comb, candidates[idx])
//			dfs(target-candidates[idx], idx)
//			comb = comb[:len(comb)-1]
//		}
//	}
//	dfs(target, 0)
//	return
//}
