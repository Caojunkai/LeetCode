package lcof

type Stack []int

func (s *Stack) Push(data int) {
	*s = append(*s, data)
}

func (s *Stack) Pop() int {
	length := len(*s)
	if length == 0 {
		return -1
	}

	item := (*s)[length-1]
	*s = (*s)[:length-1]
	return item
}

type CQueue struct {
	in  *Stack
	out *Stack
}

func (q *CQueue) Pop() int {
	if len(*q.out) > 0 {
		return q.out.Pop()
	}

	for len(*q.in) > 0 {
		q.out.Push(q.in.Pop())
	}

	return q.out.Pop()
}

func (q *CQueue) Push(v int) {
	q.in.Push(v)
}
