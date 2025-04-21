package main

// 11. 盛最多水的容器 https://leetcode.cn/problems/container-with-most-water/description/?envType=study-plan-v2&envId=top-100-liked
// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 返回容器可以储存的最大水量。
// 说明：你不能倾斜容器。

// 以下为双指针做法，还可以暴力解决（计算所有两条边围成的区域大小取最大值）
func maxArea(height []int) int {
	max := 0
	left, right := 0, len(height)-1 // 双指针思想，设定左右指针边界值（right-left为底）
	for left < right {
		area := 0
		if height[left] > height[right] { // 短板效应
			area = height[right] * (right - left)
		} else {
			area = height[left] * (right - left)
		}
		if area > max {
			max = area
		}
		// 动态判断最大值的高度
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}

// 双指针正确性证明：https://leetcode.cn/problems/container-with-most-water/solutions/207215/sheng-zui-duo-shui-de-rong-qi-by-leetcode-solution/?envType=study-plan-v2&envId=top-100-liked
