package leetcode

/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1 := len(nums1)
	n2 := len(nums2)
	l := n1 + n2
	m1 := l / 2
	m2 := -1
	if l%2 == 0 {
		m2 = m1 - 1
	}
	i := 0
	i1 := 0
	i2 := 0
	for i <= m1 {
		var a int
		if i1 >= n1 {
			a = nums2[i2]
			i2++
		} else if i2 >= n2 {
			a = nums1[i1]
			i1++
		} else {
			if nums1[i1] < nums2[i2] {
				a = nums1[i1]
				i1++
			} else {
				a = nums2[i2]
				i2++
			}
		}
		if i == m1 {
			m1 = a
			break
		}
		if i == m2 {
			m2 = a
		}
		i++
	}
	if m2 == -1 {
		return float64(m1)
	} else {
		return (float64(m1) + float64(m2)) / 2
	}
}

// @lc code=end
