package leetcode

/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 */

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := []int{}
	m := map[int]int{}
	for _, num := range nums2 {
		for {
			if len(stack) == 0 {
				break
			}
			last := stack[len(stack)-1]
			if num > last {
				m[last] = num
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		if _, ok := m[num]; ok {
			continue
		}
		m[num] = -1
		stack = append(stack, num)
	}
	res := []int{}
	for _, num := range nums1 {
		if n, ok := m[num]; ok {
			res = append(res, n)
		} else {
			res = append(res, -1)
		}
	}
	return res
}

// @lc code=end
