/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */
package leetcode

// @lc code=start
func coinChange(coins []int, amount int) int {
	l := []int{0}
	for i := 1; i <= amount; i++ {
		var temp []int
		for _, v := range coins {
			if i >= v {
				temp = append(temp, l[i-v]+1)
			}
		}
		l = append(l, MinSlice(temp))
	}
	return l[amount]
}

func MinSlice(nums []int) (rsp int) {
	for k, v := range nums {
		if k == 0 || v < rsp {
			rsp = v
		}
	}
	return
}

// @lc code=end
