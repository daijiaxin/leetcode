package leetcode

/*
 * @lc app=leetcode.cn id=733 lang=golang
 *
 * [733] 图像渲染
 */

// @lc code=start
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]
	if oldColor == newColor {
		return image
	}
	m := len(image)
	n := len(image[0])
	image[sr][sc] = newColor
	if sr-1 >= 0 && image[sr-1][sc] == oldColor {
		image = floodFill(image, sr-1, sc, newColor)
	}
	if sr+1 < m && image[sr+1][sc] == oldColor {
		image = floodFill(image, sr+1, sc, newColor)
	}
	if sc-1 >= 0 && image[sr][sc-1] == oldColor {
		image = floodFill(image, sr, sc-1, newColor)
	}
	if sc+1 < n && image[sr][sc+1] == oldColor {
		image = floodFill(image, sr, sc+1, newColor)
	}
	return image
}

// @lc code=end
