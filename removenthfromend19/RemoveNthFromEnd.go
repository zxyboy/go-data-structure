package main

import "fmt"

//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//
//
// 示例 2：
//
//
//输入：head = [1], n = 1
//输出：[]
//
//
// 示例 3：
//
//
//输入：head = [1,2], n = 1
//输出：[1]
//
//
//
//
// 提示：
//
//
// 链表中结点的数目为 sz
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz
//
//
//
//
// 进阶：你能尝试使用一趟扫描实现吗？
// Related Topics 链表 双指针 👍 2122 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) (result *ListNode) {
	listLen, prdDeleteNodeIdx := 0, 0
	var preDeleteNode *ListNode = nil
	result = head
	for head != nil {
		listLen++
		prdDeleteNodeIdx = listLen - n
		if prdDeleteNodeIdx >= 1 {
			if preDeleteNode == nil {
				preDeleteNode = result
			} else {
				preDeleteNode = preDeleteNode.Next
			}
		}
		head = head.Next
	}

	if prdDeleteNodeIdx > 0 {
		next := preDeleteNode.Next
		preDeleteNode.Next = next.Next
		next.Next = nil
	} else {
		next := result.Next
		result.Next = nil
		result = next
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	head = &ListNode{
		Val:  1,
		Next: nil,
	}
	n := 1
	nthFromEnd := removeNthFromEnd(head, n)
	for l := nthFromEnd; l != nil; l = l.Next {
		fmt.Println(l.Val)
	}
}
