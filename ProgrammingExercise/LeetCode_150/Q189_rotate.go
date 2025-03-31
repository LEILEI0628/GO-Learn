package main

// 189. 轮转数组 https://leetcode.cn/problems/rotate-array/description/?envType=study-plan-v2&envId=top-interview-150

func rotate(nums []int, k int) {
	length := len(nums)
	if k > length { // 防止k>length导致的数组越界
		k = k % length
	}
	result := make([]int, length)
	copy(result[0:], nums[length-k:])
	copy(result[k:], nums[0:length-k])
	copy(nums, result)
}

func rotate1(nums []int, k int) {
	length := len(nums)
	if k > length {
		k = k % length
	}
	result := make([]int, 0, length)
	for i := 0; i < length; i++ {
		result = append(result, nums[(length-k+i)%length])
	}
	copy(nums, result)
}

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums1, 9)
}
