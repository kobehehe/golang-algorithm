package main

import (
	"fmt"
	"golang-algorithm/binaryTree"
	"os"
)

func main() {


	treeMap := []int{9, 3, 10,15, 7}
	aa := binaryTree.New(treeMap)
	bb := isBalanced(aa)
	fmt.Println(bb)
	binaryTree.PrintTree(os.Stdout, aa.Root, 0, 'M')
}

func isBalanced(tree *binaryTree.Tree) bool {
	if tree == nil {
		return false
	}
	lenRight := binaryTree.GetTreeHeight(tree.Root.Right)
	lenLeft := binaryTree.GetTreeHeight(tree.Root.Left)
	a := lenRight - lenLeft
	if WithBranch(a) <= 1 {
		return true
	}
	return false
}

func WithBranch(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
