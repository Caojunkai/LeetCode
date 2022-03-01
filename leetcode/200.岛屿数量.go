/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */
package leetcode

// @lc code=start
func numIslands(grid [][]byte) int {
	var resp int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				resp++
				dfs200(grid, i, j)
			}
		}
	}
	return resp
}

func dfs200(grid [][]byte, x, y int) {
	rows := len(grid)
	cols := len(grid[0])

	if x < 0 || y < 0 || x >= rows || y >= cols || grid[x][y] == '0' {
		return
	}

	grid[x][y] = '0'

	dfs200(grid, x, y-1)
	dfs200(grid, x, y+1)
	dfs200(grid, x-1, y)
	dfs200(grid, x+1, y)
}

// @lc code=end
