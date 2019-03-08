/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 *
 * https://leetcode-cn.com/problems/length-of-last-word/description/
 *
 * algorithms
 * Easy (28.80%)
 * Total Accepted:    18.3K
 * Total Submissions: 63.4K
 * Testcase Example:  '"Hello World"'
 *
 * 给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。
 *
 * 如果不存在最后一个单词，请返回 0 。
 *
 * 说明：一个单词是指由字母组成，但不包含任何空格的字符串。
 *
 * 示例:
 *
 * 输入: "Hello World"
 * 输出: 5
 *
 *
 */
func lengthOfLastWord(s string) int {
	ret := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if ret != 0 {
				return ret
			}
			continue
		}
		ret++
	}
	return ret
}

