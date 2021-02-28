package main

func main() {
	
}

//输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。
//
//例如，一个链表有 6 个节点，从头节点开始，它们的值依次是 1、2、3、4、5、6。这个链表的倒数第 3 个节点是值为 4 的节点。
//
// 
//
//示例：
//
//给定一个链表: 1->2->3->4->5, 和 k = 2.
//
//返回链表 4->5.
//

func getKthFromEnd1(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	node,_ := getKthFromEndre(head,k)
	return node
}

func getKthFromEndre(head *ListNode, k int)(*ListNode,int){
	if head == nil {
		return nil,0
	}
	//递
	node,res := getKthFromEndre(head.Next,k)
	//归
	if res == k {
		return node,k
	}
	res++
	return head,k
}



func getKthFromEnd2(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	var res []*ListNode
	for head != nil {
		res = append (res,head.Next)
		head = head.Next
	}
	l := len(res)
	if l >= k{
		return res[l-k]
	}
	return nil
}



