package leetcode

/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] 判断子序列
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	si := 0
	ti := 0
	sl := len(s)
	tl := len(t)
	for si < sl && ti < tl {
		if s[si] == t[ti] {
			si++
			ti++
		} else {
			ti++
		}
	}
	if si == sl {
		return true
	}
	return false
}

// @lc code=end
