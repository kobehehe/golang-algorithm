package main

type linkNode struct {
	Val int
	Next *linkNode
}
/**
删除链表指定的节点
 */
func main() {
	
}

func deleteNode(head *linkNode,num int) *linkNode {
	if head == nil || head.Next == nil {
		return nil
	}
	if head.Val == num {
		return head.Next
	}
	if head.Next.Val == num {
		head.Next = head.Next.Next
	}
	deleteNode(head.Next,num)
	return head
}
