package leetcode

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	r := &ListNode{}
	head.Next = r
	for list1 != nil || list2 != nil {
		var tmp *ListNode
		if list1 == nil {
			tmp = list2
			list2 = list2.Next
		} else if list2 == nil {
			tmp = list1
			list1 = list1.Next
		} else {
			if list1.Val < list2.Val {
				tmp = list1
				list1 = list1.Next
			} else {
				tmp = list2
				list2 = list2.Next
			}
		}
		r.Next = tmp
		r = r.Next
	}
	return head.Next.Next
}

// @lc code=end
