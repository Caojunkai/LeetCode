/*
 * @lc app=leetcode.cn id=1380 lang=golang
 *
 * [1380] 矩阵中的幸运数
 */
package leetcode

// @lc code=start
func luckyNumbers(matrix [][]int) []int {
	resp := make([]int, 0)
	for _, v := range matrix {
		rowMinValue, colIndex := getMinValueFromRow(v)
		if isMaxInCol(matrix, colIndex, rowMinValue) {
			resp = append(resp, rowMinValue)
		}
	}

	return resp
}

func getMinValueFromRow(row []int) (int, int) {
	if len(row) == 0 {
		return 0, 0
	}

	minValue, index := row[0], 0

	for k, v := range row {
		if v < minValue {
			minValue = v
			index = k
		}
	}

	return minValue, index
}

func isMaxInCol(matrix [][]int, colIndex, targetValue int) bool {

	var maxValue int
	for k := range matrix {
		if temp := matrix[k][colIndex]; temp > maxValue {
			maxValue = temp
		}
	}

	return maxValue == targetValue
}

// @lc code=end
