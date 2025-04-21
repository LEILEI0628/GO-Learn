package main

// 283. 移动零 https://leetcode.cn/problems/move-zeroes/description/?envType=study-plan-v2&envId=top-100-liked
// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。

func moveZeroes(nums []int) {
	i, j := 0, 0 // i指向前端最后非零元素的后一位（最先出现的零元素），j指向后端最前非零元素
	for j < len(nums) {
		if nums[i] != 0 {
			i++ // 遇到非零元素指针后移
			j++
		} else {
			j++
		}
		if i != j && j < len(nums) && nums[j] != 0 {
			nums[i] = nums[j]
			nums[j] = 0
		}
	}
}

func moveZeroes1(nums []int) {
	i, j := 0, 0 // i指向前端最后非零元素的后一位（最先出现的零元素），j指向后端最前非零元素
	for j < len(nums) {
		if nums[j] != 0 {
			if i != j { // 前几位元素不为0时不移动（快慢指针同时移动）
				nums[i], nums[j] = nums[j], 0 // 快指针非零则将非零元素移动到慢指针的位置
			}
			i++ // 慢指针++
		}
		j++ // 快指针++
	}
}

// 方法一：双指针
// 思路及解法
// 使用双指针，左指针指向当前已经处理好的序列的尾部，右指针指向待处理序列的头部。
// 右指针不断向右移动，每次右指针指向非零数，则将左右指针对应的数交换，同时左指针右移。
// 注意到以下性质：
// 左指针左边均为非零数；
// 右指针左边直到左指针处均为零。
// 因此每次交换，都是将左指针的零与右指针的非零数交换，且非零数的相对顺序并未改变。
func moveZeroes2(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
