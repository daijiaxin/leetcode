package leetcode

/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	index := 0
	for _, num := range nums {
		if num != 0 {
			nums[index] = num
			index++
		}
	}
	for i := index; i < len(nums); i++ {
		nums[i] = 0
	}
}

// @lc code=end
