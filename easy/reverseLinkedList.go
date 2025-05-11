package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func newNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	list := []int{}

	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}

	newList := newNode(list[len(list)-1])
	returnVal := newList
	for i := len(list) - 2; i >= 0; i-- {
		newList.Next = newNode(list[i])
		newList = newList.Next
	}

	return returnVal
}
