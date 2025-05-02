package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func newNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

func halfAdd(x, y int) (value, carry int) {
	sum := x + y
	if sum > 9 {
		return (x + y) % 10, 1
	} else {
		return x + y, 0
	}
}

func fullAdd(x, y, z int) (value, carry int) {
	sum := x + y + z
	if sum > 9 {
		carry := int(sum / 10)
		value := sum % 10

		return value, carry
	} else {
		return sum, 0
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	curr1 := l1
	curr2 := l2
	val, carry := halfAdd(curr1.Val, curr2.Val)

	head := newNode(val)
	returnVal := head

	curr1 = curr1.Next
	curr2 = curr2.Next
	for curr1 != nil && curr2 != nil {
		val, carry = fullAdd(curr1.Val, curr2.Val, carry)
		fmt.Println(val)
		head.Next = newNode(val)
		head = head.Next
		curr1 = curr1.Next
		curr2 = curr2.Next
	}

	if curr1 != nil {
		for curr1 != nil {
			val, carry = fullAdd(curr1.Val, 0, carry)
			head.Next = newNode(val)
			head = head.Next
			curr1 = curr1.Next
		}
	} else if curr2 != nil {
		for curr2 != nil {
			val, carry = fullAdd(curr2.Val, 0, carry)
			head.Next = newNode(val)
			head = head.Next
			curr2 = curr2.Next
		}
	}

	if carry != 0 {
		head.Next = newNode(carry)
	}
	return returnVal.Next
}
