package leetcode

/*
 * @lc app=leetcode.cn id=875 lang=golang
 *
 * [875] 爱吃香蕉的珂珂
 */

// @lc code=start
func minEatingSpeed(piles []int, h int) int {
	min := 1
	max := int(1e9)
	for min <= max {
		mid := (min + max) >> 1
		if check(piles, h, mid) {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return min
}

func check(piles []int, h int, k int) bool {
	sum := 0
	for _, v := range piles {
		sum += (v-1)/k + 1
	}
	return sum <= h
}

// @lc code=end
