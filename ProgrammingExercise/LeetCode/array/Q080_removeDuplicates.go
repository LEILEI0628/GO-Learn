package array

// 80. 删除有序数组中的重复项 II https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/description/?envType=study-plan-v2&envId=top-interview-150

func Q080removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	slow := 2
	for fast := 2; fast < len(nums); fast++ {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func Q080removeDuplicatesA1(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	slow, fast := 2, 2
	for fast < n {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
