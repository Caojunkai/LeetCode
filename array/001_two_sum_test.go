package main

import (
	"testing"
)

var (
	nums = []int{2, 7, 11, 15, 233, 32, 2323, 2323, 2323, 2321, 32332}
	target = 32334
)

func BenchmarkTwoSum1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwoSum1(nums, target)
	}
}

func BenchmarkTwoSum2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwoSum2(nums, target)
	}
}
