package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=1552 lang=golang
 *
 * [1552] 两球之间的磁力
 */

// @lc code=start
func maxDistance(position []int, m int) int {
	sort.Ints(position)
	start := 0
	end := int(1e9)
	for start <= end {
		mid := (start + end) >> 1
		if check(position, m, mid) {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return end
}

func check(position []int, m int, k int) bool {
	l := len(position)
	start := 0
	count := 1
	for i := 1; i < l; i++ {
		if position[i]-position[start] >= k {
			count++
			start = i
		}
	}
	return count >= m
}

// @lc code=end
