package leetcode

/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	m1 := map[int]int{}
	m2 := map[int]int{}
	for _, num := range nums1 {
		m1[num] = 1
	}
	for _, num := range nums2 {
		if m1[num] == 1 {
			m2[num] = 1
		}
	}
	keys := make([]int, 0, len(m2))
	for k := range m2 {
		keys = append(keys, k)
	}
	return keys
}

// @lc code=end
