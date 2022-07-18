package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历：递归实现
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}
	leftVal := inorderTraversal(root.Left)
	rightVal := inorderTraversal(root.Right)

	leftVal = append(leftVal, root.Val)
	leftVal = append(leftVal, rightVal...)
	return leftVal
}

// 中序遍历：循环实现
func inorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}
	p := root
	stack := make([]*TreeNode, 0)
	result := make([]int, 0)
	for p != nil {
		if p.Left != nil {
			stack = append(stack, p)
			p = p.Left
		} else {
			// 本节点
			result = append(result, p.Val)
			// 右节点
			if p.Right != nil {
				p = p.Right
			} else {
				var parent *TreeNode = nil
				for parent = stack[len(stack)-1]; parent.Right == nil; parent = stack[len(stack)-1] {
					result = append(result, parent.Val)

				}
				result = append(result, parent.Val)
				stack = stack[:len(stack)-1]
				p = parent.Right
			}
		}
	}
	return result
}

func main() {
	root := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	traversal := inorderTraversal2(root)

	// 3 4 5 6 7 8 9
	for _, e := range traversal {
		fmt.Printf("%d ", e)
	}

}
