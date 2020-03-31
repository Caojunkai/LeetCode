package classic

import "testing"

var (
	nums = []int{1, 3, 22, 12, 43, 6, 2, 12, 345, 2, -10}
)

func TestBubbleSort(t *testing.T) {
	BubbleSort(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			t.Fail()
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(nums)
	}
}

func TestSelectSort(t *testing.T) {
	SelectSort(nums)
	t.Logf("%v", nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			t.Fail()
		}
	}
}

func BenchmarkSelectSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectSort(nums)
	}
}
