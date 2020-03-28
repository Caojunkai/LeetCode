/*
 * @lc app=leetcode.cn id=914 lang=golang
 *
 * [914] 卡牌分组
 */
package leetcodecn

//nolint
// @lc code=start
func hasGroupsSizeX(deck []int) bool {
	length := len(deck)
	if length < 2 {
		return false
	}
	count := make(map[int]int)
	for _, v := range deck {
		count[v]++
	}
	var g int
	for _, v := range count {
		if g == 0 {
			g = v
			continue
		}
		g = gcd(g, v)
	}
	return g >= 2
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// @lc code=end
