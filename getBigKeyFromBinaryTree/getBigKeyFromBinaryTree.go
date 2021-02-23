package main

import (
	"fmt"
	"golang-algorithm/binaryTree"
)
//给定一棵二叉搜索树，请找出其中第k大的节点。
func main() {
	treeMap := []int{9, 3, 10,15, 7}
	binaryTree := binaryTree.New(treeMap)
	var res []int
	dfsNode(binaryTree.Root,&res)
	k := 2
	fmt.Println(res,res[len(res)-k])
}

func dfsNode(n *binaryTree.TreeNode, res *[]int)  {
	if n !=nil {
		dfsNode(n.Left,res)
		*res = append(*res,n.Val)
		dfsNode(n.Right,res)
	}
}


