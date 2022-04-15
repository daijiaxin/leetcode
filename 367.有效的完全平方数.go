package leetcode

/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	i := 0
	j := num + 1
	for {
		if i > j {
			return false
		}
		m := (i + j) / 2
		mm := m * m
		if mm == num {
			return true
		}
		if mm > num {
			j = m - 1
		} else {
			i = m + 1
		}
	}
}

// @lc code=end
