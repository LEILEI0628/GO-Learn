package string

// 27. 移除元素 https://leetcode.cn/problems/remove-element/?envType=study-plan-v2&envId=top-interview-150

func removeElement(nums []int, val int) int {
	diff := make([]int, 0, len(nums))
	for _, v := range nums {
		if v != val {
			diff = append(diff, v)
		}
	}
	copy(nums[0:], diff[0:])
	return len(diff)
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
