package hash

// 1. 两数之和 https://leetcode.cn/problems/two-sum/description/?envType=study-plan-v2&envId=top-100-liked
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
// 你可以按任意顺序返回答案。

// 暴力枚举
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 方法二：哈希表
// 思路及算法
// 注意到方法一的时间复杂度较高的原因是寻找 target - x 的时间复杂度过高。因此，我们需要一种更优秀的方法，能够快速寻找数组中是否存在目标元素。如果存在，我们需要找出它的索引。
// 使用哈希表，可以将寻找 target - x 的时间复杂度降低到从 O(N) 降低到 O(1)。
// 这样我们创建一个哈希表，对于每一个 x，我们首先查询哈希表中是否存在 target - x，然后将 x 插入到哈希表中，即可保证不会让 x 和自己匹配。
func twoSum1(nums []int, target int) []int {
	//var hashTable map[int]int // 仅声明未初始化
	//hashTable = map[int]int{}
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		// 可以直接通过key计算并查找（target-x）
		// 因为是两数之和，所以当答案为相同x时（例：3+3=6）直接返回即可
		// 也不存在key重复需要返回多个答案的问题
		hashTable[x] = i

	}
	return nil
}
