package leetcode

/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */

// @lc code=start
func numDecodings(s string) int {
	n := len(s)
	// a = f[i-2], b = f[i-1], c = f[i]
	a, b, c := 0, 1, 0
	for i := 1; i <= n; i++ {
		c = 0
		if s[i-1] != '0' {
			c += b
		}
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0') <= 26) {
			c += a
		}
		a, b = b, c
	}
	return c
}

// @lc code=end
