package leetcode

/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	l := len(nums) / 2
	m := map[int]int{}
	for _, num := range nums {
		m[num]++
		if m[num] > l {
			return num
		}
	}
	return 0
}

// @lc code=end
