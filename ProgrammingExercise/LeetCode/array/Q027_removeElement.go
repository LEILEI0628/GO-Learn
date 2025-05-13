package array

// 27. 移除元素 https://leetcode.cn/problems/remove-element/?envType=study-plan-v2&envId=top-interview-150
// 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素。元素的顺序可能发生改变。然后返回 nums 中与 val 不同的元素的数量。
// 假设 nums 中不等于 val 的元素数量为 k，要通过此题，您需要执行以下操作：
// 更改 nums 数组，使 nums 的前 k 个元素包含不等于 val 的元素。nums 的其余元素和 nums 的大小并不重要。
// 返回 k。

// 双指针做法
func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for ; fast < len(nums); fast++ {
		if nums[fast] != val {
			if slow != fast {
				nums[slow] = nums[fast] // 剪枝
			}
			slow++
		}
	}
	return slow
}

func removeElement1(nums []int, val int) int {
	i := 0
	for j, v := range nums {
		if nums[j] != val {
			//if i == j {
			//	i++
			//	continue
			//}
			nums[i] = v
			i++
		}
	}
	return i
}

func removeElementA1(nums []int, val int) int {
	left := 0
	for _, v := range nums { // v 即 nums[right]
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}

func removeElementA2(nums []int, val int) int {
	left, right := 0, len(nums)
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return left
}
