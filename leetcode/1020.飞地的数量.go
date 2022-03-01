/*
 * @lc app=leetcode.cn id=1020 lang=golang
 *
 * [1020] 飞地的数量
 */
package leetcode

// @lc code=start
func numEnclaves(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	for i := 0; i < rows; i++ {
		dfs1020(grid, i, 0)
		dfs1020(grid, i, cols-1)
	}

	for i := 0; i < cols; i++ {
		dfs1020(grid, 0, i)
		dfs1020(grid, rows-1, i)
	}

	var resp int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				resp++
			}
		}
	}
	return resp
}

func dfs1020(grid [][]int, x, y int) {
	rows, cols := len(grid), len(grid[0])
	if x < 0 || y < 0 || x >= rows || y >= cols || grid[x][y] == 0 {
		return
	}

	grid[x][y] = 0

	dfs1020(grid, x+1, y)
	dfs1020(grid, x-1, y)
	dfs1020(grid, x, y+1)
	dfs1020(grid, x, y-1)
}

// @lc code=end
