/*
 * @lc app=leetcode.cn id=1254 lang=golang
 *
 * [1254] 统计封闭岛屿的数目
 */
package leetcode

// @lc code=start
func closedIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])

	for i := 0; i < rows; i++ {
		dfs1254(grid, i, 0)
		dfs1254(grid, i, cols-1)
	}

	for i := 0; i < cols; i++ {
		dfs1254(grid, 0, i)
		dfs1254(grid, rows-1, i)
	}

	var resp int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 {
				resp++
				dfs1254(grid, i, j)
			}
		}
	}

	return resp
}

func dfs1254(grid [][]int, x, y int) {
	rows, cols := len(grid), len(grid[0])

	if x < 0 || y < 0 || x >= rows || y >= cols || grid[x][y] == 1 {
		return
	}

	grid[x][y] = 1

	dfs1254(grid, x+1, y)
	dfs1254(grid, x-1, y)
	dfs1254(grid, x, y+1)
	dfs1254(grid, x, y-1)
}

// @lc code=end
