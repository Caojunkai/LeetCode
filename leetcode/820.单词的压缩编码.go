/*
 * @lc app=leetcode.cn id=820 lang=golang
 *
 * [820] 单词的压缩编码
 */
package leetcode

// nolint
// @lc code=start
func minimumLengthEncoding(words []string) int {
	hash := make(map[string]int, len(words))
	for _, v := range words {
		hash[v] = 0
	}
	for k := range hash {
		for i := 1; i < len(k); i++ {
			if _, ok := hash[k[i:]]; ok {
				delete(hash, k[i:])
			}
		}
	}
	var rsp int
	for k := range hash {
		// #号长度也要计算
		rsp += len(k) + 1
	}
	return rsp
}

// @lc code=end
