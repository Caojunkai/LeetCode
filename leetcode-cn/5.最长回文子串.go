/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 *
 * https://leetcode-cn.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (24.72%)
 * Total Accepted:    41.9K
 * Total Submissions: 169.1K
 * Testcase Example:  '"babad"'
 *
 * 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
 *
 * 示例 1：
 *
 * 输入: "babad"
 * 输出: "bab"
 * 注意: "aba" 也是一个有效答案。
 *
 *
 * 示例 2：
 *
 * 输入: "cbbd"
 * 输出: "bb"
 *
 *
 */
func longestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return s
	}
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	// 计算d[i][j]是 保证d[i+1][j-1]以计算完成
	var left, right int
	for i := n - 2; i >= 0; i-- {
		dp[i][i] = true
		for j := i + 1; j < n; j++ {
			// 两端相等 在3个以下字符肯定是回文
			dp[i][j] = s[i] == s[j] && (j-i < 3 || dp[i+1][j-1])
			if dp[i][j] && right-left < j-i {
				left = i
				right = j
			}
		}
	}
	return s[left : right+1]
}

