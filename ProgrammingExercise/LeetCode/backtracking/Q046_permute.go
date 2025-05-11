package backtracking

// 46. 全排列 https://leetcode.cn/problems/permutations/?envType=study-plan-v2&envId=top-100-liked
// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
func permute(nums []int) (result [][]int) {
	path := make([]int, 0, len(nums))
	used := make([]bool, len(nums))
	var back func(used []bool)
	back = func(used []bool) {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}
		for i, num := range nums {
			if !used[i] {
				path = append(path, num)
				used[i] = true
				back(used)
				path = path[:len(path)-1]
				used[i] = false
			}
		}
	}
	back(used)
	return
}

func permuteV1(nums []int) (result [][]int) {
	path := make([]int, 0, len(nums))
	var back func()
	back = func() {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}
		for _, num := range nums {
			have := false
			for _, v := range path {
				if v == num {
					have = true
				}
			}
			if !have {
				path = append(path, num)
				back()
				path = path[:len(path)-1]
			}
		}
	}
	back()
	return
}
