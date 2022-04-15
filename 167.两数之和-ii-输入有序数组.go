package leetcode

/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	var result []int
	for {
		if left >= right {
			break
		}
		sum := numbers[left] + numbers[right]
		if sum == target {
			result = append(result, left+1, right+1)
			break
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	return result
}

// @lc code=end
