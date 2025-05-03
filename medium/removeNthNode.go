package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// we are gonna make the difference between our 2 pointers n
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	index := 1
	returnVal := head
	slow := head
	head = head.Next
	for head != nil {
		if head.Next == nil {
			if index == n {
				slow.Next = slow.Next.Next
				break
			} else if slow.Next == head {
				slow.Next = nil
			} else {
				return returnVal.Next
			}
		}

		if index == n {
			slow = slow.Next
		} else {
			index++
		}
		head = head.Next
	}

	return returnVal
}
