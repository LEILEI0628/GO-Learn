package hash

// 454. 四数相加 II https://leetcode.cn/problems/4sum-ii/description/
// 给你四个整数数组 nums1、nums2、nums3 和 nums4 ，数组长度都是 n ，请你计算有多少个元组 (i, j, k, l) 能满足：
// 0 <= i, j, k, l < n
// nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) (count int) { // 分组 + hash
	hash := map[int]int{}
	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			hash[num1+num2]++
		}
	}
	for _, num3 := range nums3 {
		for _, num4 := range nums4 {
			if c, ok := hash[0-(num3+num4)]; ok {
				count += c
			}
		}
	}
	return
}

func fourSumCountV1(nums1 []int, nums2 []int, nums3 []int, nums4 []int) (count int) { // 分组 + hash
	hash1 := map[int]int{}
	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			hash1[num1+num2]++
		}
	}
	hash2 := map[int]int{} // 可以直接判断而不用再次进行此操作
	for _, num3 := range nums3 {
		for _, num4 := range nums4 {
			hash2[num3+num4]++
		}
	}
	for num, c1 := range hash1 {
		if c2, ok := hash2[0-num]; ok {
			count += c1 * c2
		}
	}
	return
}
