package leetcode

/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	for {
		if m == 0 && n == 0 {
			break
		} else if m == 0 && n != 0 {
			nums1[n-1] = nums2[n-1]
			n--
		} else if m != 0 && n != 0 {
			if nums1[m-1] > nums2[n-1] {
				nums1[m+n-1] = nums1[m-1]
				m--
			} else {
				nums1[m+n-1] = nums2[n-1]
				n--
			}
		} else {
			m--
		}
	}
}

// @lc code=end
