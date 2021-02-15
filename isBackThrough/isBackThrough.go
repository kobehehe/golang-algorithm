package main

import (
	"fmt"
)

/**
数组里的数字是否是二叉树后序遍历的结果
 */
func main() {
	treeMap := []int{1,6,3,2,5}
	//twoBinaryTree := binaryTree.New(treeMap)
	result := isTwoBinaryTree(treeMap)
	fmt.Println(result)
}

func isTwoBinaryTree(treeMap []int) bool {
	if len(treeMap) <= 1 {
		return true
	}

	root := treeMap[len(treeMap)-1]

	//找到左右子树分界点
	index := 0
	for i,v := range treeMap{
		if v > root {
			index = i
			break
		}
	}

	leftArr := treeMap[:index]
	rightArr := treeMap[index:len(treeMap)-1]

	for _,v := range leftArr{
		if v > root{
			return false
		}
	}

	for _,v := range rightArr{
		if v < root{
			return false
		}
	}


	return isTwoBinaryTree(leftArr) && isTwoBinaryTree(rightArr)
}
