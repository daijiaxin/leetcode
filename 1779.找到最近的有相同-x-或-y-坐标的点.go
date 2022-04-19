package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=1779 lang=golang
 *
 * [1779] 找到最近的有相同 X 或 Y 坐标的点
 */

// @lc code=start
func nearestValidPoint(x int, y int, points [][]int) int {
	size := math.MaxFloat64
	index := -1
	for i, point := range points {
		if point[0] != x && point[1] != y {
			continue
		}
		count := math.Abs(float64(point[0]-x)) + math.Abs(float64(point[1]-y))
		if count < size {
			size = count
			index = i
		}
	}
	return index
}

// @lc code=end
