package array

// 704. 二分查找 https://leetcode.cn/problems/binary-search/description/
// 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
func search(nums []int, target int) int { // 左闭右闭写法
	left, right := 0, len(nums)-1 // 左闭右闭的左右边界为数组下标最小/最大值
	for left <= right {           // 左闭右闭写法时左右边界可以相等
		mid := left + (right-left)/2 // 避免越界的写法
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1 // 已经判断出nums[mid]不是目标值，需要跳过（右同理）
		} else {
			right = mid - 1
		}
	}
	return -1
}

func searchV1(nums []int, target int) int { // 左闭右开写法
	left, right := 0, len(nums) // 右开所以右边界为最大下标+1
	for left < right {          // 左闭右开写法左右边界不能相等（左边界闭能取到右边界开取不到，矛盾）
		mid := left + (right-left)/2 // 避免越界的写法
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1 // 左闭同上
		} else {
			right = mid // 右开不包含 nums[mid]这个元素，下次查找需要放进来
		}
	}
	return -1
}
