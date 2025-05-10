package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if root.Left != nil && root.Right != nil {
		temp := root.Left
		root.Left = root.Right
		root.Right = temp
		invertTree(root.Left)
		invertTree(root.Right)
	} else if root.Left == nil {
		root.Left = root.Right
		root.Right = nil
		invertTree(root.Left)
	} else if root.Right == nil {
		root.Right = root.Left
		root.Left = nil
		invertTree(root.Right)
	}

	return root
}
