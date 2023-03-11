package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	time := []int{2}
	ans := sortedArrayToBST(time)
	fmt.Println(ans)
}

func sortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBSTUtil(nums, 0, len(nums)-1)
}

func sortedArrayToBSTUtil(nums []int, l, r int) *TreeNode {
	if l <= r {
		mid := (l + r) / 2
		node := makeNode(nums[mid])
		node.Left = sortedArrayToBSTUtil(nums, l, mid-1)
		node.Right = sortedArrayToBSTUtil(nums, mid+1, r)
		return node
	}
	return nil
}

func makeNode(num int) *TreeNode {
	return &TreeNode{Val: num}
}
