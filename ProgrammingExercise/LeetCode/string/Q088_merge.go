package string

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

func mergeA1(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2) // 合并
	sort.Ints(nums1)       // 排序
}

func mergeA2(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for {
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}

		if p2 == n {
			sorted = append(sorted, nums1[p1:]...)
			break
		}

		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}
	copy(nums1, sorted)
}

func mergeA3(nums1 []int, m int, nums2 []int, n int) {
	for p1, p2, tail := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0; tail-- {
		var cur int
		if p1 == -1 {
			cur = nums2[p2]
			p2--
		} else if p2 == -1 {
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}
		nums1[tail] = cur
	}
}
