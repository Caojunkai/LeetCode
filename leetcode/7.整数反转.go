/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */
package leetcode

// @lc code=start
func reverse(x int) int {

	maxInt32 := 1<<31 - 1
	minInt32 := -1 << 31

	var resp int
	for x != 0 {
		resp = resp*10 + x%10
		if resp > maxInt32 || resp < minInt32 {
			return 0
		}
		x /= 10
	}

	return resp
}

// @lc code=end
