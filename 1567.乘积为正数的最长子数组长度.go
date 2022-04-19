package leetcode

/*
 * @lc app=leetcode.cn id=1567 lang=golang
 *
 * [1567] 乘积为正数的最长子数组长度
 */

// @lc code=start
func getMaxLen(nums []int) int {
	ans := 0
	pos, neg := 0, 0
	for _, num := range nums {
		if num > 0 {
			pos++
			if neg > 0 {
				neg++
			}
		} else if num == 0 {
			pos, neg = 0, 0
		} else {
			if neg > 0 {
				pos, neg = neg+1, pos+1
			} else {
				pos, neg = 0, pos+1
			}
		}
		ans = max(ans, pos)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
