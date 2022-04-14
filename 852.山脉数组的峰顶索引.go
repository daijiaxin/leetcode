package leetcode

/*
 * @lc app=leetcode.cn id=852 lang=golang
 *
 * [852] 山脉数组的峰顶索引
 */

// @lc code=start
func peakIndexInMountainArray(arr []int) int {
	start := 0
	end := len(arr) - 1
	for {
		i := (start + end) / 2
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			return i
		} else if arr[i] > arr[i+1] {
			end = i
		} else {
			start = i
		}
	}
}

// @lc code=end
