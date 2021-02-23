package main

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	aa := &ListNode{
		Val: 1,
		Next: nil,
	}
	reverseList1(aa)
}

//迭代法
func reverseList1(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head

	for curr != nil{
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}


//递归法
/**
运用递归的三个条件：
1可以拆分为若干个子问题
2子问题的解决方式和大问题解决方案一样
3最小的子问题存在边界
 */
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//首先递
	p := reverseList2(head.Next)

	head.Next.Next = head
	head.Next = nil

	//然后归
	return p
}