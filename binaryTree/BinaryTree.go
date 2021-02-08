//实现查找二叉树

package binaryTree

import (
	"fmt"
	"io"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	Root *TreeNode
}

func (t *Tree) insertTree(data int) *Tree {
	if t.Root == nil {
		t.Root = &TreeNode{Val: data}
	} else {
		t.Root.insertTreeNode(data)
	}
	return t
}

func New(node []int) *Tree {
	lenTree := len(node)
	if lenTree == 0{
		return nil
	}
	tree := &Tree{}

	for _,v := range node{
		tree.insertTree(v)
	}
	return tree
}

func (n *TreeNode) insertTreeNode(data int) {
	if n == nil {
		return
	}
	if data < n.Val {
		if n.Left == nil {
			n.Left = &TreeNode{Val: data}
		} else {
			n.Left.insertTreeNode(data)
		}
	} else {
		if n.Right == nil {
			n.Right = &TreeNode{Val: data}
		} else {
			n.Right.insertTreeNode(data)
		}
	}
}

func PrintTree(w io.Writer, node *TreeNode, ch int, nodeName rune) {

	if node == nil {
		return
	}

	for i := 0; i < ch; i++ {
		fmt.Fprint(w, "  ")
	}
	fmt.Fprintf(w, "%c:%v\n", nodeName, node.Val)

	PrintTree(w, node.Left, ch+1, 'L')
	PrintTree(w, node.Right, ch+1, 'R')

}

func main() {

	treeCurrent := &Tree{}

	treeCurrent.insertTree(1).
		insertTree(10).
		insertTree(11).
		insertTree(21).
		insertTree(31).
		insertTree(41).
		insertTree(15).
		insertTree(16)
	//printTree(os.Stdout,treeCurrent.root,0,'M')
	aa := GetTreeHeight(treeCurrent.Root)
	fmt.Println(aa)
}

func GetTreeHeight(t *TreeNode) int {
	if t == nil {
		return 0
	}
	leftLen := GetTreeHeight(t.Left)
	rightLen := GetTreeHeight(t.Right)
	if leftLen > rightLen {
		return leftLen + 1
	} else {
		return rightLen + 1
	}
}
