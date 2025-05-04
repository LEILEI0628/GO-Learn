package main

import "sort"

// 15. 三数之和 https://leetcode.cn/problems/3sum/description/?envType=study-plan-v2&envId=top-100-liked
// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。
// 注意：答案中不可以包含重复的三元组。

func threeSum(nums []int) (ans [][]int) {
	n := len(nums)
	if n < 3 {
		return nil
	}
	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] { // 确保和上一次枚举数据不同
			continue
		}
		if nums[i]+nums[i+1]+nums[i+2] > 0 { // 排序后某三个元素相加大于零，后面都大于零（剪枝）
			break
		}
		if nums[i]+nums[n-2]+nums[n-1] < 0 { // 排序后某三个元素相加小于零，前面都小于零（剪枝）
			continue
		}
		j, k := i+1, n-1 // 左右指针分别指向当前元素后一位，数组最后
		for j < k {      // 左指针需小于右指针
			if nums[i]+nums[j]+nums[k] < 0 { // 和小于零，左指针需右移
				j++
			} else if nums[i]+nums[j]+nums[k] > 0 { // 和大于零，右指针需左移
				k--
			} else {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				// 找到答案后将左右指针向中间移动至没有重复元素为止（去重操作）
				for j++; j < k && nums[j] == nums[j-1]; j++ {
				}
				for k--; j < k && nums[k] == nums[k+1]; k-- {
				}
			}
		}
	}
	return
}

func threeSum1(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b b对应的指针指向a指针后的元素
		for second := first + 1; second < n; second++ { // 此时结果小了，右移左指针
			// 需要和上一次枚举的数不相同
			if second != first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target { // 此时结果大了，左移右指针
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

// 三重循环，时间复杂度很高
func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k := j + 1; k < len(nums); k++ {
				if k != j+1 && nums[k] == nums[k-1] {
					continue
				}
				if nums[i]+nums[j]+nums[k] == 0 {
					result = append(result, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}

	return result
}

// 三重循环，时间复杂度很高
func threeSum3(nums []int) [][]int {
	sort.Ints(nums)
	resultMap := make(map[[3]int]bool)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					resultMap[[3]int{nums[i], nums[j], nums[k]}] = true
				}
			}
		}
	}
	result := make([][]int, 0, len(resultMap))
	for k, _ := range resultMap {
		result = append(result, []int{k[0], k[1], k[2]})
	}
	return result
}
