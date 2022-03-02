/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */
package leetcode

// @lc code=start
func minWindow(s string, t string) string {
	need, window := make(map[byte]int), make(map[byte]int)

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left, right, match := 0, 0, 0
	start, minLength := 0, 1<<63-1

	for right < len(s) {
		v := s[right]
		right++
		if _, ok := need[v]; ok {
			window[v]++
			if window[v] == need[v] {
				match++
			}
		}
		// 收缩
		for match == len(need) {
			if right-left < minLength {
				start = left
				minLength = right - left
			}
			leftItem := s[left]
			if _, ok := need[leftItem]; ok {
				if window[leftItem] == need[leftItem] {
					match--
				}
				window[leftItem]--
			}
			left++
		}
	}

	if minLength > len(s) {
		return ""
	}

	return s[start : start+minLength]
}

// @lc code=end
