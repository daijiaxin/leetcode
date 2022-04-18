package leetcode

/*
 * @lc app=leetcode.cn id=1539 lang=golang
 *
 * [1539] 第 k 个缺失的正整数
 */

// @lc code=start
func findKthPositive(arr []int, k int) int {
	res := 0
	i := 0
	n := 1
	for {
		if search(arr, n) == -1 {
			res = n
			i++
		}
		n++
		if i == k {
			break
		}
	}
	return res
}

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for {
		if start > end {
			break
		}
		i := (start + end) / 2
		if nums[i] == target {
			return i
		} else if nums[i] > target {
			end = i - 1
		} else {
			start = i + 1
		}
	}
	return -1
}

// @lc code=end
