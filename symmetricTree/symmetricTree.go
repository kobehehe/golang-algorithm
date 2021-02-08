package main

import "golang-algorithm/binaryTree"

func main() {

}

func isSymmetricTree(node *binaryTree.TreeNode) bool {
	if node == nil {
		return true
	}
	return isSame(node.Left, node.Right)
}

func isSame(l, r *binaryTree.TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil{
		return false
	}
	return l.Val == r.Val && isSame(l.Left,r.Right) && isSame(r.Right,l.Left)
}
