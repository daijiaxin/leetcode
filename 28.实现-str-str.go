package leetcode

/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	lh := len(haystack)
	ln := len(needle)
	if ln == 0 {
		return 0
	}
	if ln > lh {
		return -1
	}
	for i := 0; i <= lh-ln; i++ {
		if haystack[i:i+ln] == needle {
			return i
		}
	}
	return -1
}

// @lc code=end
