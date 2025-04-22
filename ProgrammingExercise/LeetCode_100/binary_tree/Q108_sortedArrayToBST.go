package main

import (
	"math/rand"
	"time"
)

// 108. 将有序数组转换为二叉搜索树 https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/description/?envType=study-plan-v2&envId=top-100-liked
// 给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 平衡 二叉搜索树。（平衡二叉树：平衡二叉树 是指该树所有节点的左右子树的高度相差不超过 1。）

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	rootIndex := len(nums) / 2
	return &TreeNode{
		Val:   nums[rootIndex],
		Left:  sortedArrayToBST(nums[0:rootIndex]),
		Right: sortedArrayToBST(nums[rootIndex+1:]),
	}
}

// 方法一：中序遍历，总是选择中间位置左边的数字作为根节点
func sortedArrayToBST1(nums []int) *TreeNode {
	return helper1(nums, 0, len(nums)-1)
}

func helper1(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = helper1(nums, left, mid-1)
	root.Right = helper1(nums, mid+1, right)
	return root
}

// 方法二：中序遍历，总是选择中间位置右边的数字作为根节点
func sortedArrayToBST2(nums []int) *TreeNode {
	return helper2(nums, 0, len(nums)-1)
}

func helper2(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	// 总是选择中间位置右边的数字作为根节点
	mid := (left + right + 1) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = helper2(nums, left, mid-1)
	root.Right = helper2(nums, mid+1, right)
	return root
}

// 方法三：中序遍历，选择任意一个中间位置数字作为根节点
func sortedArrayToBST3(nums []int) *TreeNode {
	rand.Seed(time.Now().UnixNano())
	return helper3(nums, 0, len(nums)-1)
}

func helper3(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	// 选择任意一个中间位置数字作为根节点
	mid := (left + right + rand.Intn(2)) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = helper3(nums, left, mid-1)
	root.Right = helper3(nums, mid+1, right)
	return root
}
