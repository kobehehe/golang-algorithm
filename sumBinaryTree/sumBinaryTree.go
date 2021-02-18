package main

import (
	"fmt"
	"golang-algorithm/binaryTree"
)

//输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。
//
// 
//
//示例:
//给定如下二叉树，以及目标和 sum = 22，

//5
/// \
//4   8
///   / \
//11  13  4
///  \    / \
//7    2  5   1
//返回:
//
//[
//[5,4,11,2],
//[5,8,4,5]
//]

func main() {
	treeMap := []int{9, 3, 10, 15, 7}
	tree := binaryTree.New(treeMap)
	var res [][]int
	getPath(tree.Root,19,[]int{},&res)
	fmt.Println(res)
}

func getPath(root *binaryTree.TreeNode,num int,arr []int,res *[][]int){
	if root == nil{
		return
	}

	arr = append(arr,root.Val)

	if root.Val == num && root.Left==nil && root.Right == nil {
		temp := make([]int,len(arr))
		copy(temp,arr)
		*res = append(*res,temp)
	}
	getPath(root.Left,num-root.Val,arr,res)
	getPath(root.Right,num-root.Val,arr,res)

}










