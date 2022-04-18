package leetcode

/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	m := map[rune]int{}
	for _, a := range s {
		m[a]++
	}
	for _, b := range t {
		m[b]--
		if m[b] == 0 {
			delete(m, b)
		}
	}
	if len(m) == 0 {
		return true
	} else {
		return false
	}
}

// @lc code=end
