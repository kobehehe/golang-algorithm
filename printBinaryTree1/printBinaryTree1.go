package main

import (
	"fmt"
	"golang-algorithm/binaryTree"
)
/**
从上到下打印二叉树 同一层节点 从左到右打印
广度优先搜索策略
 */
func main() {
	treeMap := []int{9, 3, 10,15, 7}
	tree := binaryTree.New(treeMap)
	bb := printTreeByWidth(tree.Root)
	fmt.Println(bb)
}

func printTreeByWidth(treeNode *binaryTree.TreeNode) []int  {
	if treeNode == nil {
		return nil
	}
	var queue []*binaryTree.TreeNode
	var res []int
	queue = append(queue,treeNode)
	for len(queue)!=0 {
		cRoot := queue[0]
		res = append(res,cRoot.Val)
		if cRoot.Left != nil {
			queue = append(queue,cRoot.Left)
		}

		if cRoot.Right != nil{
			queue = append(queue,cRoot.Right)
		}
		queue = queue[1:]
	}
	return res
}
