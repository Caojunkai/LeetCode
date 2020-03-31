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

func InsertSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		j := i
		for j > 0 && nums[j] < nums[j-1] {
			swap(nums, j, j-1)
			j--
		}
	}
}

func MergeSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) / 2

	MergeSort(nums, left, mid)
	MergeSort(nums, mid+1, right)

	// 创建临时数组
	temp := make([]int, right-left+1)
	i := left
	j := mid + 1
	cur := 0

	// 由于两侧都是有序数组 从index=0比较两侧数组
	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			temp[cur] = nums[i]
			i++
		} else {
			temp[cur] = nums[j]
			j++
		}
		cur++
	}

	// 左侧i <= mid 说明左侧值比较大 直接将剩余数组添加到临时数组后边
	for ; i <= mid; i++ {
		temp[cur] = nums[i]
		cur++
	}

	// 同上
	for ; j <= right; j++ {
		temp[cur] = nums[j]
		cur++
	}

	// 将排序后的临时数组写入原始数字
	for i := 0; i < right-left+1; i++ {
		nums[left+i] = temp[i]
	}
}
