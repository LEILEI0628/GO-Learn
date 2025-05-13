package main

// 209. 长度最小的子数组 https://leetcode.cn/problems/minimum-size-subarray-sum/description/
// 给定一个含有 n 个正整数的数组和一个正整数 target 。
// 找出该数组中满足其总和大于等于 target 的长度最小的 子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
func minSubArrayLen(target int, nums []int) int {
	sum, result, i := 0, 0, 0        // i是头部（subArr）
	for j := 0; j < len(nums); j++ { // j是尾部
		sum += nums[j]
		for sum >= target { // 当和大于目标时
			if result == 0 || result > j-i+1 {
				result = j - i + 1
			}
			sum -= nums[i] // 子数组头部向后移动试图找到最小长度的子串
			i++
		}
	}
	return result
}
