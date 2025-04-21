package main

import "sort"

// 128. 最长连续序列 https://leetcode.cn/problems/longest-consecutive-sequence/description/?envType=study-plan-v2&envId=top-100-liked
//给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

// 排序
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums) // 排序
	result := 1
	i := 1
	count := 1 // 连续长度
	for i < len(nums) {
		if nums[i] == nums[i-1] {
			// 数组中元素重复时跳过
		} else if nums[i]-nums[i-1] == 1 {
			count++ // 连续时
		} else {
			count = 1 // 不连续，连续长度清零
		}

		if count > result { // 判断最长连续长度
			result = count
		}
		i++
	}
	return result
}
