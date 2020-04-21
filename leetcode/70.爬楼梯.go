/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */
package leetcode

// f(n) = f(n - 1) + f(n - 2)
// @lc code=start
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	d := make([]int, n+1)
	d[1] = 1
	d[2] = 2
	for i := 3; i <= n; i++ {
		d[i] = d[i-1] + d[i-2]
	}
	return d[n]
}

// @lc code=end
