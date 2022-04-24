package leetcode

/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	m := map[string]bool{}
	return wBreak(s, wordDict, &m)
}

func wBreak(s string, wordDict []string, m *map[string]bool) bool {
	if r, ok := (*m)[s]; ok {
		return r
	}
	l := len(s)
	for _, word := range wordDict {
		w := len(word)
		if l == w {
			if s == word {
				(*m)[s] = true
				return true
			} else {
				continue
			}
		} else if l < w {
			continue
		} else {
			if s[:w] == word && wBreak(s[w:], wordDict, m) {
				(*m)[s] = true
				return true
			}
		}
	}
	(*m)[s] = false
	return false
}

// @lc code=end
