package _range

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	} else if len(nums) == 1 {
		return []string{fmt.Sprintf("%d", nums[0])}
	}
	result := make([]string, 0, len(nums))
	i := 0
	j := 1
	for ; j < len(nums); j++ {
		if nums[j]-nums[j-1] != 1 {
			if i == j-1 {
				result = append(result, fmt.Sprintf("%d", nums[i]))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", nums[i], nums[j-1]))
			}
			i = j
		}
	}
	if i == j-1 {
		result = append(result, fmt.Sprintf("%d", nums[i]))
	} else {
		result = append(result, fmt.Sprintf("%d->%d", nums[i], nums[j-1]))
	}
	return result
}

func summaryRanges1(nums []int) (ans []string) {
	if len(nums) == 0 {
		return []string{}
	}
	start := 0
	end := 0
	for end < len(nums) {
		end++
		for (end < len(nums)) && (nums[end]-nums[end-1] == 1) {
			end++
		}
		if start == end-1 {
			ans = append(ans, fmt.Sprintf("%d", nums[start]))
		} else {
			ans = append(ans, fmt.Sprintf("%d->%d", nums[start], nums[end-1]))
		}
		start = end

	}
	return
}

func summaryRangesA1(nums []int) (ans []string) {
	for i, n := 0, len(nums); i < n; {
		left := i
		for i++; i < n && nums[i-1]+1 == nums[i]; i++ {
		}
		s := strconv.Itoa(nums[left])
		if left < i-1 {
			s += "->" + strconv.Itoa(nums[i-1])
		}
		ans = append(ans, s)
	}
	return
}
