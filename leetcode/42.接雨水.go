/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */
package leetcode

// @lc code=start
func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}

	var sum int

	maxLeftHeights, maxRightHeights := make([]int, len(height)), make([]int, len(height))

	for i := 1; i < len(height); i++ {
		maxLeftHeights[i] = maxInt(height[i-1], maxLeftHeights[i-1])
	}

	for i := len(height) - 2; i >= 0; i-- {
		maxRightHeights[i] = maxInt(height[i+1], maxRightHeights[i+1])
	}

	for i := 0; i < len(height); i++ {
		min := minInt(maxLeftHeights[i], maxRightHeights[i])
		if min > height[i] {
			sum += min - height[i]
		}
	}

	return sum
}

func maxInt(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
