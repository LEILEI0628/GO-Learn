package array

// 238. 除自身以外数组的乘积 https://leetcode.cn/problems/product-of-array-except-self/description/
// 给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
// 题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
// 请 不要使用除法，且在 O(n) 时间复杂度内完成此题。
func productExceptSelf(nums []int) []int {
	length := len(nums)
	result := make([]int, length)
	result[0] = 1
	for i := 1; i < length; i++ {
		result[i] = result[i-1] * nums[i-1]
	}
	rightProduct := 1
	for i := length - 1; i >= 0; i-- {
		result[i] = result[i] * rightProduct
		rightProduct *= nums[i]
	}
	return result
}

func productExceptSelfV1(nums []int) []int {
	length := len(nums)
	result := make([]int, length)
	leftProduct := make([]int, length)
	rigthtProduct := make([]int, length)
	leftProduct[0] = 1
	for i := 1; i < length; i++ {
		leftProduct[i] = leftProduct[i-1] * nums[i-1]
	}
	rigthtProduct[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		rigthtProduct[i] = rigthtProduct[i+1] * nums[i+1]
	}
	for i := 0; i < length; i++ {
		result[i] = leftProduct[i] * rigthtProduct[i]
	}
	return result
}

func productExceptSelfV2(nums []int) (result []int) {
	for i := 0; i < len(nums); i++ {
		product := 1
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			product *= nums[j]
		}
		result = append(result, product)
	}
	return
}
