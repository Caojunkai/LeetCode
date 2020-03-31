package lcof

type CircularList struct {
	Value int
	Next  *CircularList
}

// f(n) = (f(n-1) + k) % n
// nolint
func lastRemaining(n int, m int) int {
	// if n < 1 {
	// 	return n
	// }
	// var head, p, q *CircularList
	// for i := 1; i <= 3; i++ {
	// 	p = &CircularList{
	// 		Value: i,
	// 	}
	// 	if i == 1 {
	// 		head = p
	// 	} else {
	// 		q.Next = p
	// 	}
	// 	q = p
	// }
	// q.Next = head
	var rsp int
	for i := 2; i <= n; i++ {
		rsp = (rsp + m) % i
	}
	return rsp
}
