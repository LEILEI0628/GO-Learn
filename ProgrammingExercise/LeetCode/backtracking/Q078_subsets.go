package backtracking

// 78. 子集 https://leetcode.cn/problems/subsets/description/?envType=study-plan-v2&envId=top-100-liked
func subsets(nums []int) (result [][]int) {
	path := make([]int, 0, len(nums))
	var back func(start int)
	back = func(start int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		result = append(result, tmp)
		//if start >= len(nums) { // 终止条件可以忽略：start >= len(nums)时都不会进入for循环
		//	return
		//}
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			back(i + 1)
			path = path[:len(path)-1]
		}
	}
	back(0)
	return
}
