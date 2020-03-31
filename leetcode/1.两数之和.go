/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */
package leetcode

// nolint
// @lc code=start
func twoSum(nums []int, target int) []int {

	// solution-1
	// or i, v := range nums {
	// 	for j := i + 1; j < length; j++ {
	// 		if nums[j]+v == target {
	// 			return []int{
	// 				i,
	// 				j,
	// 			}
	// 		}
	// 	}
	// }
	// return []int{}

	// solution-2
	hash := make(map[int]int)
	for i, v := range nums {
		if j, ok := hash[target-v]; ok {
			return []int{
				i,
				j,
			}
		}
		hash[v] = i
	}
	return []int{}
}

// @lc code=end
