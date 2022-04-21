package leetcode

/*
 * @lc app=leetcode.cn id=994 lang=golang
 *
 * [994] 腐烂的橘子
 */

// @lc code=start
func orangesRotting(grid [][]int) int {
	ml := len(grid)
	nl := len(grid[0])
	l := [][]int{}
	t := [][]int{}
	g := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	c := 0
	r := 0
	for m, gr := range grid {
		for n, g := range gr {
			if g == 1 {
				c++
			}
			if g == 2 {
				t = append(t, []int{m, n})
			}
		}
	}
	if c == 0 {
		return 0
	}
	for len(t) > 0 {
		if c == 0 {
			return r
		}
		r++
		l = append(l, t...)
		t = [][]int{}
		for len(l) > 0 {
			m := l[0][0]
			n := l[0][1]
			l = l[1:]
			for i := 0; i < 4; i++ {
				mm := m + g[i][0]
				nn := n + g[i][1]
				if mm >= 0 && mm < ml && nn >= 0 && nn < nl && grid[mm][nn] == 1 {
					c--
					grid[mm][nn] = 2
					t = append(t, []int{mm, nn})
				}
			}
		}
	}
	return -1
}

// @lc code=end
