/*
 * @lc app=leetcode.cn id=1905 lang=golang
 *
 * [1905] 统计子岛屿
 */
package leetcode

// @lc code=start
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	rows, cols := len(grid1), len(grid1[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid2[i][j] == 1 && grid1[i][j] == 0 {
				dfs1905(grid2, i, j)
			}
		}
	}

	var resp int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid2[i][j] == 1 {
				resp++
				dfs1905(grid2, i, j)
			}
		}
	}

	return resp
}

func dfs1905(grid [][]int, x, y int) {
	rows, cols := len(grid), len(grid[0])

	if x < 0 || y < 0 || x >= rows || y >= cols {
		return
	}

	if grid[x][y] == 0 {
		return
	}

	grid[x][y] = 0

	dfs1905(grid, x+1, y)
	dfs1905(grid, x-1, y)
	dfs1905(grid, x, y+1)
	dfs1905(grid, x, y-1)
}

// @lc code=end
