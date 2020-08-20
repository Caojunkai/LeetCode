package main

import (
	"sync"

	"github.com/fitzix/leetcode/classic"
)

func main() {
	wight := []int{2, 2, 4, 6, 3}
	classic.Knapsack(wight, len(wight), 9)
}

type Set struct {
	lock sync.RWMutex
	m    map[int64]struct{}
}

func (s *Set) Add(data int64) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, ok := s.m[data]; !ok {
		s.m[data] = struct{}{}
	}
}
