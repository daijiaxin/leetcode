package leetcode

/*
 * @lc app=leetcode.cn id=1672 lang=golang
 *
 * [1672] 最富有客户的资产总量
 */

// @lc code=start
func maximumWealth(accounts [][]int) int {
	max := 0
	for _, acc := range accounts {
		sum := 0
		for _, a := range acc {
			sum = sum + a
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

// @lc code=end
