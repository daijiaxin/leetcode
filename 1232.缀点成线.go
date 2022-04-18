package leetcode

/*
 * @lc app=leetcode.cn id=1232 lang=golang
 *
 * [1232] 缀点成线
 */

// @lc code=start
func checkStraightLine(coordinates [][]int) bool {
	deltaX, deltaY := coordinates[0][0], coordinates[0][1]
	for _, p := range coordinates {
		p[0] -= deltaX
		p[1] -= deltaY
	}
	A, B := coordinates[1][1], -coordinates[1][0]
	for _, p := range coordinates[2:] {
		x, y := p[0], p[1]
		if A*x+B*y != 0 {
			return false
		}
	}
	return true
}

// @lc code=end
