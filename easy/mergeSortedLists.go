package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func newNode(v int) *ListNode {
	return &ListNode{
		Val:  v,
		Next: nil,
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		head = newNode(list1.Val)
		list1 = list1.Next
	} else {
		head = newNode(list2.Val)
		list2 = list2.Next
	}
	returnVal := head

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			head.Next = newNode(list1.Val)
			list1 = list1.Next
		} else {
			head.Next = newNode(list2.Val)
			list2 = list2.Next
		}
		head = head.Next
	}

	if list1 != nil {
		for list1 != nil {
			head.Next = newNode(list1.Val)
			list1 = list1.Next
			head = head.Next
		}
	} else if list2 != nil {
		for list2 != nil {
			head.Next = newNode(list2.Val)
			list2 = list2.Next
			head = head.Next
		}
	}

	return returnVal
}
