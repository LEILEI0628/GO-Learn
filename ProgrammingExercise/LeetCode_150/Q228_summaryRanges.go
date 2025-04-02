package main

import "fmt"

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
