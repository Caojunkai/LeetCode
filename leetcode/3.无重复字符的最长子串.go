/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */
package leetcode

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	entity := make(map[byte]int)
	var left, max int

	for i := 0; i < len(s); i++ {
		if k, ok := entity[s[i]]; ok {
			left = Max(left, k+1)
		}
		entity[s[i]] = i
		max = Max(max, i-left+1)
	}
	return max
}

func Max(left, right int) int {
	if left <= right {
		return right
	}
	return left
}

// @lc code=end
