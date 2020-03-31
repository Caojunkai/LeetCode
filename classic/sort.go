package classic

func swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func BubbleSort(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				swap(nums, j, j+1)
			}
		}
	}
}

func SelectSort(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		var maxIndex int
		for j := 1; j <= i; j++ {
			if nums[maxIndex] < nums[j] {
				maxIndex = j
			}
		}
		swap(nums, maxIndex, i)
	}
}
