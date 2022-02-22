/*
 * @lc app=leetcode.cn id=838 lang=golang
 *
 * [838] 推多米诺
 */
package leetcode

import "strings"

// @lc code=start
func pushDominoes(dominoes string) string {
	left := -1

	arr := make([]byte, len(dominoes))

	for k, v := range dominoes {
		if v == '.' {
			continue
		}

		if v == 'L' {
			// 左侧没有力
			if left == -1 {
				for i := 0; i < k; i++ {
					arr[i] = 'L'
				}
			} else {
				if arrow := dominoes[left]; arrow == 'L' {
					for i := left; i < k; i++ {
						arr[i] = 'L'
					}
				} else {
					mid := (k + left) / 2

					for i := left; i < mid; i++ {
						arr[i] = 'R'
					}

					if (k+left)%2 != 0 {
						arr[mid] = 'R'
					}

					for i := mid + 1; i < k; i++ {
						arr[i] = 'L'
					}
				}
			}
			left = k
			arr[left] = 'L'
		}

		if v == 'R' {
			// 左侧没有力
			if left == -1 {
				// do nothing
			} else {
				if arrow := dominoes[left]; arrow == 'R' {
					for i := left; i < k; i++ {
						arr[i] = 'R'
					}
				}
			}

			left = k
			arr[left] = 'R'
		}
	}

	if left != -1 && dominoes[left] == 'R' && dominoes[len(dominoes)-1] == '.' {
		for i := left; i < len(dominoes); i++ {
			arr[i] = 'R'
		}
	}

	var resp strings.Builder
	for _, v := range arr {
		temp := v
		if v == 0 {
			temp = '.'
		}
		resp.WriteByte(temp)
	}
	return resp.String()
}

// @lc code=end
