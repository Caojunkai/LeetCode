/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */
package leetcode

// @lc code=start
func myAtoi(str string) int {
	s := 0
	flag, sign := true, 1
	max, min := 1<<31-1, -1<<31
	dict := map[rune]int{
		48: 0,
		49: 1,
		50: 2,
		51: 3,
		52: 4,
		53: 5,
		54: 6,
		55: 7,
		56: 8,
		57: 9,
	}
	for _, v := range str {

		if v == 32 && flag {
			continue
		}
		if v == 43 && flag {
			sign = 1
			flag = false
			continue
		}
		if v == 45 && flag {
			sign = -1
			flag = false
			continue
		}

		flag = false

		if val, ok := dict[v]; ok {
			s = s*10 + val

			tmp := s * sign

			if tmp > max {
				return max
			}

			if tmp < min {
				return min
			}
			continue
		}
		break
	}

	return s * sign
}

// @lc code=end
