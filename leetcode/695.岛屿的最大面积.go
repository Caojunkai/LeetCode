/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] 岛屿的最大面积
 */
package leetcode

// @lc code=start
func maxAreaOfIsland(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	var resp int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				resp = max695(resp, dfs695(grid, i, j))
			}
		}
	}

	return resp
}

func dfs695(grid [][]int, x, y int) int {
	rows, cols := len(grid), len(grid[0])
	// 越界
	if x < 0 || y < 0 || x >= rows || y >= cols {
		return 0
	}
	// 海水
	if grid[x][y] == 0 {
		return 0
	}

	grid[x][y] = 0

	return dfs695(grid, x-1, y) +
		dfs695(grid, x+1, y) +
		dfs695(grid, x, y-1) +
		dfs695(grid, x, y+1) + 1

}

func max695(x, y int) int {
	if x > y {
		return x
	}

	return y
}

// @lc code=end
