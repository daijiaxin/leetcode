package leetcode

/*
 * @lc app=leetcode.cn id=1588 lang=golang
 *
 * [1588] 所有奇数长度子数组的和
 */

// @lc code=start
func sumOddLengthSubarrays(arr []int) int {
	sum := 0
	n := len(arr)
	prefixSums := make([]int, n+1)
	for i, v := range arr {
		prefixSums[i+1] = prefixSums[i] + v
	}
	for start := range arr {
		for length := 1; start+length <= n; length += 2 {
			end := start + length - 1
			sum += prefixSums[end+1] - prefixSums[start]
		}
	}
	return sum
}

// @lc code=end
