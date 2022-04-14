package leetcode

/*
 * @lc app=leetcode.cn id=278 lang=golang
 *
 * [278] 第一个错误的版本
 */

func isBadVersion(version int) bool

// @lc code=start
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	start_number := 1
	end_number := n
	if isBadVersion(start_number) {
		return start_number
	}
	for {
		i := (start_number + end_number) / 2
		if isBadVersion(i) {
			end_number = i
		} else {
			start_number = i
		}
		if end_number-1 <= start_number {
			break
		}
	}
	return end_number
}

// @lc code=end
