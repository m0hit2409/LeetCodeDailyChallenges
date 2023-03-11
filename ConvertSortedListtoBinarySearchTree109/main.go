package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	ans := sortedListToBST(&ListNode{Val: -10, Next: &ListNode{Val: -3, Next: &ListNode{Val: 0, Next: &ListNode{Val: 5, Next: &ListNode{Val: 9}}}}})
	fmt.Println(ans)
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	return sortedToBST(head, nil)
}

func sortedToBST(head, tail *ListNode) *TreeNode {
	if head == tail {
		return nil
	}

	slow, fast := head, head

	for fast != tail && fast.Next != tail {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return &TreeNode{Val: slow.Val, Left: sortedToBST(head, slow), Right: sortedToBST(slow.Next, tail)}
}
