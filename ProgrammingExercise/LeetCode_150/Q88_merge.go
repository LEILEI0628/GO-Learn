package main

import "sort"

// 88. 合并两个有序数组（https://leetcode.cn/problems/merge-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150）
func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ { // 对nums2进行遍历
		nums1[m+i] = nums2[i]      // 当出现num2中的元素比num1中的每个元素都大时直接插入在num1尾部
		for j := 0; j < m+i; j++ { // 比较nums1的每个元素
			if nums2[i] < nums1[j] { // num2当前元素小于nums1中的元素
				// 插入到该元素之前
				for k := m + i - 1; k >= j; k-- {
					nums1[k+1] = nums1[k]
				}
				nums1[j] = nums2[i]
				break
			}
		}
	}
}

func merge1(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2) // 合并
	sort.Ints(nums1)       // 排序
}
