package leetcode

/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */

// @lc code=start
func countBits(n int) []int {
	r := []int{0}
	for i := 1; i <= n; i++ {
		a := i / 2
		b := i % 2
		r = append(r, r[a]+b)
	}
	return r
}

// @lc code=end
