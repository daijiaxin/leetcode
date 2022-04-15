package leetcode

/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	m1 := map[int]int{}
	for _, num := range nums1 {
		m1[num]++
	}
	res := []int{}
	for _, num := range nums2 {
		if m1[num] > 0 {
			res = append(res, num)
			m1[num]--
		}
	}
	return res
}

// @lc code=end
