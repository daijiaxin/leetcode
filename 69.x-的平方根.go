package leetcode

/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	i := 0
	j := x + 1
	for {
		if i > j {
			return i - 1
		}
		m := (i + j) / 2
		mm := m * m
		if mm == x {
			return m
		}
		if mm > x {
			j = m - 1
		} else {
			i = m + 1
		}
	}
}

// @lc code=end
