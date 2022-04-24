package leetcode

/*
 * @lc app=leetcode.cn id=633 lang=golang
 *
 * [633] 平方数之和
 */

// @lc code=start
func judgeSquareSum(c int) bool {
	start := 0
	end := judgeSquare(c)
	if end*end == c {
		return true
	}
	for start <= end {
		eec := c - end*end
		start = judgeSquare(eec)
		if start*start == eec {
			return true
		} else {
			end--
		}
	}
	return false
}

func judgeSquare(c int) int {
	start := 0
	end := c
	for start <= end {
		i := (start + end) / 2
		ii := i * i
		if ii > c {
			end = i - 1
		} else if ii < c {
			start = i + 1
		} else {
			return i
		}
	}
	return end
}

// @lc code=end
