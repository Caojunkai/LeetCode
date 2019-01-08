package main

// 暴力双循环遍历
func TwoSum1(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n-1; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func TwoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for i, value := range nums {
		if j, ok := m[target - value]; ok && i != j {
			return []int{i, j}
		}
		m[value] = i
	}
	return nil
}


