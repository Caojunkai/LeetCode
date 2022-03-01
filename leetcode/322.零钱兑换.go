/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */
package leetcode

// @lc code=start
func coinChange(coins []int, amount int) int {
	f := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		f[i] = -1
		for _, coin := range coins {
			if i >= coin && f[i-coin] != -1 {
				if f[i] != -1 {
					f[i] = min322(f[i], f[i-coin])
				} else {
					f[i] = f[i-coin]
				}
			}
		}

		if f[i] != -1 {
			f[i] += 1
		}
	}

	return f[amount]
}

func min322(x, y int) int {
	if x < y {
		return x
	}

	return y
}

// @lc code=end
