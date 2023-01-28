package leetcode

/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 */

// @lc code=start
func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	res := make([]int, length)
	stack := []int{}
	for i := 0; i < length; i++ {
		temp := temperatures[i]
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}
	return res
}

// @lc code=end
