package search

// 35. 搜索插入位置 https://leetcode.cn/problems/search-insert-position/description/?envType=study-plan-v2&envId=top-100-liked
// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
// 请必须使用时间复杂度为 O(log n) 的算法。
func searchInsert(nums []int, target int) int { // 二分查找法
	index, start, end := 0, 0, len(nums)
	if target <= nums[0] {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	for start < end {
		index = (start + end) / 2
		if nums[index] == target {
			return index
		} else if nums[index] > target {
			end = index
		} else {
			start = index + 1
		}
	}
	return start
}
