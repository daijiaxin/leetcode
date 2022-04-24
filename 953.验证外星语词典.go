package leetcode

/*
 * @lc app=leetcode.cn id=953 lang=golang
 *
 * [953] 验证外星语词典
 */

// @lc code=start
func isAlienSorted(words []string, order string) bool {
	m := map[byte]int{}
	for i := 0; i < len(order); i++ {
		m[order[i]] = i
	}
	for i := 0; i < len(words)-1; i++ {
		if isAlienSort(words[i], words[i+1], m) < 0 {
			return false
		}
	}
	return true
}

func isAlienSort(s1 string, s2 string, m map[byte]int) int {
	l1 := len(s1)
	l2 := len(s2)
	i := 0
	for {
		m1 := m[s1[i]]
		m2 := m[s2[i]]
		if m1 == m2 {
			i++
		} else if m1 < m2 {
			return 1
		} else {
			return -1
		}
		if i == l1 && i == l2 {
			return 0
		} else if i == l1 && i != l2 {
			return 1
		} else if i != l1 && i == l2 {
			return -1
		}
	}
}

// @lc code=end
