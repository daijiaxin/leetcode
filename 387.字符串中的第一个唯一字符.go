package leetcode

/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 */

// @lc code=start
func firstUniqChar(s string) int {
	m := map[rune]int{}
	for i, ss := range s {
		if _, ok := m[ss]; ok {
			m[ss] = -1
		} else {
			m[ss] = i
		}
	}
	for _, ss := range s {
		if m[ss] != -1 {
			return m[ss]
		}
	}
	return -1
}

// @lc code=end
