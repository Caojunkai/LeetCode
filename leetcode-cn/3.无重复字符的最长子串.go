/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 *
 * https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (28.42%)
 * Total Accepted:    87.3K
 * Total Submissions: 307K
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
 *
 * 示例 1:
 *
 * 输入: "abcabcbb"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 *
 *
 * 示例 2:
 *
 * 输入: "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 *
 *
 * 示例 3:
 *
 * 输入: "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 * 1 3 4 5 6 7 3 4 2 1
 *
 */
func lengthOfLongestSubstring(s string) int {
	n, ret, i, j := len(s), 0, 0, 0
	mp := make(map[byte]int)
	for j = 0; j < n; j++ {
		if k, ok := mp[s[j]]; ok {
			i = max(k+1, i)
		}
		ret = max(ret, j-i+1)
		mp[s[j]] = j
	}
	return ret
}

func max(left, right int) int {
	if left > right {
		return left
	}
	return right
}

