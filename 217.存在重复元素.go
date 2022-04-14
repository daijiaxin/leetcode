package leetcode

/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	m := map[int]int{}
	for _, i := range nums {
		if _, ok := m[i]; ok {
			return true
		}
		m[i] = 1
	}
	return false
}

// @lc code=end
