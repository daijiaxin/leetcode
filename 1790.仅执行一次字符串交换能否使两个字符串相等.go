package leetcode

/*
 * @lc app=leetcode.cn id=1790 lang=golang
 *
 * [1790] 仅执行一次字符串交换能否使两个字符串相等
 */

// @lc code=start
func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	f := []int{}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			f = append(f, i)
		}
	}
	if len(f) != 2 {
		return false
	}
	if s1[f[0]] == s2[f[1]] && s1[f[1]] == s2[f[0]] {
		return true
	}
	return false
}

// @lc code=end
