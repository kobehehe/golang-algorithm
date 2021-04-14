package main

import "fmt"

type listNode struct {
	val  int
	next *listNode
}

func main() {
	l1 := &listNode{
		val:  2,
		next: nil,
	}

	l2 := &listNode{
		val:  2,
		next: nil,
	}

	twoNumberSum := twoNumberSum(l1, l2)
	fmt.Println(twoNumberSum)

}

func twoNumberSum(l1, l2 *listNode) *listNode {
	headList := new(listNode)
	head := headList
	num := 0

	for l1 != nil || l2 != nil || num > 0 {
		headList.next = new(listNode)
		headList = headList.next
		if l1 != nil {
			num = num + l1.val
			l1 = l1.next
		}

		if l2 != nil {
			num = num + l2.val
			l2 = l2.next
		}
		headList.val = num % 10
		num = num / 10
		fmt.Println(num)
	}
	return head.next
}
