package leetcode

/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
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
func deleteDuplicates(head *ListNode) *ListNode {
	c := head
	var last *ListNode
	for head != nil {
		if last != nil && head.Val == last.Val {
			last.Next = head.Next
		} else {
			last = head
		}
		head = head.Next
	}
	return c
}

// @lc code=end
