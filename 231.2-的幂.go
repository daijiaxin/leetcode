package leetcode

/*
 * @lc app=leetcode.cn id=231 lang=golang
 *
 * [231] 2 的幂
 */

// @lc code=start
func isPowerOfTwo(n int) bool {
	i := 1
	for i <= n {
		if i == n {
			return true
		}
		i = i << 1
	}
	return false
}

// @lc code=end
