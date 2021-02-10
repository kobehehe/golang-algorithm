package main

import (
	"fmt"
	"golang-algorithm/binaryTree"
)

func main() {
	treeMap := []int{9, 3, 10,15, 7}
	binaryTree := binaryTree.New(treeMap)
	var res []int
	dfsNode(binaryTree.Root,&res)
	fmt.Println(res,res[len(res)-2])
}

func dfsNode(n *binaryTree.TreeNode, res *[]int)  {
	if n !=nil {
		dfsNode(n.Left,res)
		*res = append(*res,n.Val)
		dfsNode(n.Right,res)
	}
}


