/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 *
 * https://leetcode-cn.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (36.44%)
 * Total Accepted:    48.5K
 * Total Submissions: 133K
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
 *
 * 有效字符串需满足：
 *
 *
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 *
 *
 * 注意空字符串可被认为是有效字符串。
 *
 * 示例 1:
 *
 * 输入: "()"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: "()[]{}"
 * 输出: true
 *
 *
 * 示例 3:
 *
 * 输入: "(]"
 * 输出: false
 *
 *
 * 示例 4:
 *
 * 输入: "([)]"
 * 输出: false
 *
 *
 * 示例 5:
 *
 * 输入: "{[]}"
 * 输出: true
 *
 */
type Stack []byte

func (s Stack) Pop() (Stack, byte) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s Stack) Push(v byte) Stack {
	return append(s, v)
}

func (s Stack) Last() byte {
	return s[len(s)-1]
}

func isValid(s string) bool {
	mappings := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := make(Stack, 0)
	for i := range s {
		if temp, ok := mappings[s[i]]; ok && len(stack) != 0 && temp == stack.Last() {
			stack, _ = stack.Pop()
		} else {
			stack = stack.Push(s[i])
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
