package main

import (
	"fmt"
)

//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//
//
// 示例 1：
//
//
//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]
//
//
// 示例 2：
//
//
//输入：l1 = [], l2 = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：l1 = [], l2 = [0]
//输出：[0]
//
//
//
//
// 提示：
//
//
// 两个链表的节点数目范围是 [0, 50]
// -100 <= Node.val <= 100
// l1 和 l2 均按 非递减顺序 排列
//
// Related Topics 递归 链表 👍 2524 👎 0

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	tmpList1 := list1
	tmpList2 := list2

	if tmpList1 == nil && tmpList2 == nil {
		return nil
	}
	if tmpList1 == nil {
		return tmpList2
	}

	if tmpList2 == nil {
		return tmpList1
	}
	var head *ListNode = nil
	var tail *ListNode = nil
	var n1, n2 int
	for tmpList1 != nil || tmpList2 != nil {

		if list1 != nil {
			n1 = list1.Val
		}
		if list2 != nil {
			n2 = list2.Val
		}

		if n1 >= n2 {
			for ; n1 >= n2 && list2 != nil; list2 = list2.Next {

				node := &ListNode{
					Val:  n2,
					Next: nil,
				}
				if head == nil {
					head = node
				} else {
					tail.Next = node
				}
				tail = node
			}
		} else {
			for ; n1 < n2 && list1 != nil; list1 = list1.Next {
				n1 = list1.Val
				node := &ListNode{
					Val:  n1,
					Next: nil,
				}
				if head == nil {
					head = node
				} else {
					tail.Next = node
				}
				tail = node
			}
		}
	}
	return head

}
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	tmpList1 := list1
	tmpList2 := list2

	if tmpList1 == nil && tmpList2 == nil {
		return nil
	}
	if tmpList1 == nil {
		return tmpList2
	}

	if tmpList2 == nil {
		return tmpList1
	}

	preHead := &ListNode{
		Val:  -1,
		Next: nil,
	}
	tail := preHead
	for tmpList1 != nil || tmpList2 != nil {

		if tmpList1 == nil {
			tail.Next = tmpList2
			tmpList2 = nil
			break
		}

		if tmpList2 == nil {
			tail.Next = tmpList1
			tmpList1 = nil
			break
		}
		n1 := tmpList1.Val
		n2 := tmpList2.Val
		if n1 >= n2 {
			tail.Next = tmpList2
			tail = tmpList2
			tmpList2 = tmpList2.Next
		} else {
			tail.Next = tmpList1
			tail = tmpList1
			tmpList1 = tmpList1.Next
		}
	}
	return preHead.Next

}

func main() {

	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	mergeTwoLists := mergeTwoLists2(l1, l2)
	for l := mergeTwoLists; l != nil; l = l.Next {
		fmt.Println(l.Val)
	}
}
