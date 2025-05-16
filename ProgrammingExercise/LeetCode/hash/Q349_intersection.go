package hash

import "sort"

// 349. 两个数组的交集 https://leetcode.cn/problems/intersection-of-two-arrays/description/
// 给定两个数组 nums1 和 nums2 ，返回 它们的 交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。
func intersection(nums1 []int, nums2 []int) (result []int) {
	numMap := map[int]struct{}{}
	resultMap := map[int]struct{}{}
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	for _, v := range nums1 {
		numMap[v] = struct{}{}
	}
	for _, v := range nums2 {
		if _, ok := numMap[v]; ok {
			resultMap[v] = struct{}{}
		}
	}
	for num, _ := range resultMap {
		result = append(result, num)
	}
	return
}

func intersectionV1(nums1 []int, nums2 []int) (res []int) { // 排序 + 双指针
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		x, y := nums1[i], nums2[j]
		if x == y {
			if res == nil || x > res[len(res)-1] {
				res = append(res, x)
			}
			i++
			j++
		} else if x < y {
			i++
		} else {
			j++
		}
	}
	return
}
