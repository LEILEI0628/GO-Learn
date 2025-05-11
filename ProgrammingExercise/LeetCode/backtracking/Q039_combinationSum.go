package backtracking

import "sort"

// 39. 组合总和 https://leetcode.cn/problems/combination-sum/description/
// 给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
// candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
// 对于给定的输入，保证和为 target 的不同组合数少于 150 个。
func combinationSum(candidates []int, target int) (result [][]int) {
	sort.Ints(candidates)
	path := make([]int, 0)
	var backTracking func(start int, sum int)
	backTracking = func(start int, sum int) {
		if sum > target {
			return
		}
		if sum == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
		}
		for i := start; i < len(candidates); i++ {
			if sum+candidates[i] > target { // 排序后才能剪枝
				return
			}
			path = append(path, candidates[i])
			backTracking(i, sum+candidates[i])
			path = path[:len(path)-1]
		}
	}
	backTracking(0, 0)
	return
}
