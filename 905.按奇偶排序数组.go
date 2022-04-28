package leetcode

/*
 * @lc app=leetcode.cn id=905 lang=golang
 *
 * [905] 按奇偶排序数组
 */

// @lc code=start
func sortArrayByParity(nums []int) []int {
	n1 := 0
	for i, n := range nums {
		if n%2 == 0 {
			t := nums[n1]
			nums[n1] = nums[i]
			nums[i] = t
			n1++
		}
	}
	return nums
}

// @lc code=end
