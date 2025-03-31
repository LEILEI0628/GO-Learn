package main

// 26. 删除有序数组中的重复项 https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150

func removeDuplicates(nums []int) int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i] == nums[j] {
				copy(nums[j:], nums[j+1:])
				j--
				length--
			}
		}
	}
	return length
}

func removeDuplicates1(nums []int) int {
	result := make([]int, 0, len(nums))
	flag := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(result); j++ {
			if nums[i] == result[j] {
				flag = 1
				break
			}
		}
		if flag == 0 {
			result = append(result, nums[i])
		} else {
			flag = 0
		}
	}
	copy(nums, result)
	return len(result)
}

func removeDuplicatesA1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
