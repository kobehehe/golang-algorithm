package main

import (
	"golang-algorithm/binaryTree"
)
//是否是对称二叉树
func main() {

}

func isSymmetricTree(node *binaryTree.TreeNode) bool {
	if node == nil{
		return true
	}
	aa := isSame(node.Left,node.Right)
	return aa
}

func isSame(l, r *binaryTree.TreeNode) bool {
	if l == nil && r == nil {
		return true
	}

	if l == nil || r == nil{
		return false
	}
	return l.Val == r.Val && isSame(l.Left,r.Right) && isSame(l.Right,r.Left)
}
