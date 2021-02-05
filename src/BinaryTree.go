package main

import (
	"fmt"
	"io"
)

type treeNode struct {
	val int
	left *treeNode
	right *treeNode
}

type tree struct {
	root *treeNode
}

func (t *tree) insertTree(data int) *tree {
	if t.root == nil{
		t.root = &treeNode{val: data}
	}else{
		t.root.insertTreeNode(data)
	}
	return t
}

func (n *treeNode) insertTreeNode(data int)  {
	if n == nil{
		return
	}
	if data < n.val{
		if n.left == nil{
			n.left = &treeNode{val: data}
		}else{
			n.left.insertTreeNode(data)
		}
	}else{
		if n.right == nil{
			n.right = &treeNode{val: data}
		}else{
			n.right.insertTreeNode(data)
		}
	}
}

func printTree(w io.Writer,node *treeNode, ch int, nodeName rune)  {

	if node == nil{
		return
	}

	for i:=0;i<ch;i++{
		fmt.Fprint(w,"  ")
	}
	fmt.Fprintf(w,"%c:%v\n",nodeName,node.val)

	printTree(w,node.left,ch+1,'L')
	printTree(w,node.right,ch+1,'R')

}


func main() {

	treeCurrent := &tree{}
	treeCurrent.insertTree(1).
				insertTree(10).
				insertTree(11).
				insertTree(21).
				insertTree(31).
				insertTree(41).
				insertTree(15).
				insertTree(16)
	//printTree(os.Stdout,treeCurrent.root,0,'M')
	aa := getTreeHeight(treeCurrent.root)
	fmt.Println(aa)
}

func getTreeHeight(t *treeNode) int{
	if t == nil{
		return 0
	}
	leftLen := getTreeHeight(t.left)
	rightLen := getTreeHeight(t.right)
	if leftLen > rightLen{
		return leftLen+1
	}else{
		return rightLen+1
	}
}



