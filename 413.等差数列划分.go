package leetcode

/*
 * @lc app=leetcode.cn id=413 lang=golang
 *
 * [413] 等差数列划分
 */

// @lc code=start
func numberOfArithmeticSlices(nums []int) int {
	l := len(nums)
	if l <= 2 {
		return 0
	}
	dps := 0
	dpl := 0
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dpl++
			dps += dpl
		} else {
			dpl = 0
		}
	}
	return dps
}

// @lc code=end
