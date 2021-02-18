package main

import (
	"fmt"
	"golang-algorithm/binaryTree"
)

func main() {
	treeMap := []int{9, 3, 10, 15, 7}
	tree := binaryTree.New(treeMap)
	flipTree := flipBinaryTree(tree.Root)
	fmt.Println(flipTree)
}

func flipBinaryTree(root *binaryTree.TreeNode) *binaryTree.TreeNode {
	if root == nil{
		return nil
	}
	root.Left,root.Right = root.Right,root.Left
	flipBinaryTree(root.Left)
	flipBinaryTree(root.Right)
	return root
}
