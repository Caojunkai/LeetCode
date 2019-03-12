package main

func main() {
	println(dp(11))
}

func dp(n int) int {
	n++
	V := []int{1, 3, 5}
	var m = make([]int, n)
	for i := 1; i < n; i++ {
		m[i] = m[i-1] + 1
		for _, v := range V {
			if i < v {
				break
			}
			if m[i-v] < m[i-1] {
				m[i] = m[i-v] + 1
			}
		}
	}
	return m[n-1]
}
