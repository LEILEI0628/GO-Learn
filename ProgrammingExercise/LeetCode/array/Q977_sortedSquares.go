package array

import "sort"

// 977. 有序数组的平方 https://leetcode.cn/problems/squares-of-a-sorted-array/description/
// 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
func sortedSquares(nums []int) []int {
	left, right := 0, len(nums)-1
	result := make([]int, len(nums))
	i := len(nums) - 1
	for left <= right {
		leftSquares := nums[left] * nums[left]
		rightSquares := nums[right] * nums[right]
		if leftSquares > rightSquares {
			result[i] = leftSquares
			left++
		} else {
			result[i] = rightSquares
			right--
		}
		i--
	}
	return result
}

func sortedSquaresV1(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] *= nums[i]
	}
	sort.Ints(nums)
	return nums
}
