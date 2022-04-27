package leetcode

/*
 * @lc app=leetcode.cn id=1760 lang=golang
 *
 * [1760] 袋子里最少数目的球
 */

// @lc code=start
func minimumSize(nums []int, maxOperations int) int {
	start := 1
	end := int(1e9)
	for start < end {
		mid := (start + end) >> 1
		if check(nums, mid, maxOperations) {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return start
}

func check(nums []int, max int, maxOperations int) bool {
	count := 0
	for _, num := range nums {
		if count > maxOperations {
			return false
		}
		count += (num - 1) / max
	}
	if count > maxOperations {
		return false
	} else {
		return true
	}
}

// @lc code=end
