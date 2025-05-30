package search

// 34. 在排序数组中查找元素的第一个和最后一个位置 https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/description/?envType=study-plan-v2&envId=top-100-liked
// 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
// 如果数组中不存在目标值 target，返回 [-1, -1]。
// 你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。

func searchRange(nums []int, target int) []int { // 双指针算法
	if len(nums) == 0 || target < nums[0] || target > nums[len(nums)-1] {
		return []int{-1, -1}
	}
	i, j := 0, len(nums)-1
	for i <= j {
		if nums[i] == target && nums[j] == target {
			return []int{i, j}
		}
		if nums[j] != target {
			j--
		}
		if nums[i] != target {
			i++
		}
	}
	return []int{-1, -1}
}
