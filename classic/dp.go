package classic

import "fmt"

func Knapsack(weight []int, n, max int) {
	states := make([][]bool, n)

	for i := range states {
		states[i] = make([]bool, max+1)
	}

	states[0][0] = true

	if weight[0] <= max {
		states[0][weight[0]] = true
	}

	for i := 1; i < n; i++ {
		// 不放入背包
		// i-1 表示上一层状态拷贝到当前层
		for j := 0; j <= max; j++ {
			if states[i-1][j] {
				states[i][j] = states[i-1][j]
			}
		}

		// 放入背包
		for j := 0; j <= max-weight[i]; j++ {
			if states[i-1][j] {
				states[i][j+weight[i]] = true
			}
		}
	}

	for k, v := range states {
		for j, vv := range v {
			if vv {
				fmt.Printf("states[%d][%d] = %t\n", k, j, vv)
			}
		}
	}

}
