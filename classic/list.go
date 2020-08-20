package classic

type Element struct {
	prev, next *Element
	Value      interface{}
}

type List struct {
	len  int
	root *Element
}

func (l *List) Delete(el *Element) {
	if l == nil || el == nil {
		return
	}

	el.prev.next = el.next
	el.next.prev = el.prev

	el.prev = nil
	el.next = nil

	l.len--

}
