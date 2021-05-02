package main

import (
	"fmt"
	"sort"
)

/**
输出一个数组的全排列。抽象为N叉树  深度递归 回溯思想来解决
*/
func main() {
	var path []int
	var res [][]int
	nums := []int{1, 2, 3}
	sort.Ints(nums)
	used := make([]bool, len(nums))
	dfs(nums, path, &res, used)
	fmt.Println(res)
}

func dfs(nums []int, path []int, res *[][]int, used []bool) {
	if len(path) == len(nums) {
		newp := make([]int, len(path))
		copy(newp, path)
		*res = append(*res, newp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		path = append(path, nums[i])
		used[i] = true
		dfs(nums, path, res, used)
		path = path[:len(path)-1]
		used[i] = false
	}
}

func isHaveValue(arr []int, v int) bool {

	for _, v1 := range arr {
		if v1 == v {
			return true
		}
	}
	return false
}

//func dfs(nums []int,path []int, res *[][]int) {
//	 if len(path) == len(nums){
//	 	fmt.Println(path)
//		 *res = append(*res,path)
//		 fmt.Println(res)
//		 return
//	 }
//	 //回溯完删除
//
//	 for i:=0;i<len(nums);i++{
//		 path  = append(path,nums[i])
//		 dfs(nums,path,res)
//		 path =  path[:len(path)-1]
//	 }
//}
