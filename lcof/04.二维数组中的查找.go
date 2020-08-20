/**
给定一个二维数组，其每一行从左到右递增排序，从上到下也是递增排序。给定一个数，判断这个数是否在该二维数组中。
Consider the following matrix:
[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]

Given target = 5, return true.
Given target = 20, return false.
*/

package lcof

func findNumberIn2DArray(matrix [][]int, target int) bool {
	row := len(matrix)

	if row == 0 || len(matrix[0]) == 0 {
		return false
	}

	col := len(matrix[0])

	i, j, temp := 0, col-1, 0

	for i < row && j >= 0 {
		temp = matrix[i][j]
		if temp == target {
			return true
		} else if temp > target {
			j--
		} else {
			i++
		}
	}
	return false
}
