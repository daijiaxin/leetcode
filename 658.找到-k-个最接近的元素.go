package leetcode

/*
 * @lc app=leetcode.cn id=658 lang=golang
 *
 * [658] 找到 K 个最接近的元素
 */

// @lc code=start
func findClosestElements(arr []int, k int, x int) []int {
	l := len(arr)
	if k == l {
		return arr
	}
	i := searchIndex(arr, x)
	start := getIndex(arr, i-1, i, x)
	end := start
	for end-start+1 < k {
		r := getIndex(arr, start-1, end+1, x)
		if r < start {
			start = r
		} else {
			end = r
		}
	}
	return arr[start : end+1]
}

func getIndex(arr []int, l int, r int, x int) int {
	if l < 0 {
		return r
	}
	if r >= len(arr) {
		return l
	}
	da := arr[r] - x
	db := x - arr[l]
	if da >= db {
		return l
	}
	return r
}

func searchIndex(arr []int, x int) int {
	start := 0
	end := len(arr) - 1
	for start <= end {
		i := (start + end) >> 1
		if arr[i] == x {
			return i
		} else if arr[i] < x {
			start = i + 1
		} else {
			end = i - 1
		}
	}
	return start
}

// @lc code=end
