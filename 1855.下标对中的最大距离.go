package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=1855 lang=golang
 *
 * [1855] 下标对中的最大距离
 */

// @lc code=start
func maxDistance(nums1 []int, nums2 []int) int {
	res := 0
	l2 := len(nums2)
	for i, num := range nums1 {
		n := sort.Search(l2, func(l int) bool { return nums2[l] < num })
		d := n - i - 1
		if d > res {
			res = d
		}
	}
	return res
}

// @lc code=end
