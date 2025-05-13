package greedy

// 121. 买卖股票的最佳时机 https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/?envType=study-plan-v2&envId=top-interview-150

func maxProfit(prices []int) int { // 双重for循环时间复杂度过高
	max := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			p := prices[j] - prices[i]
			if p > max {
				max = p
			}
		}
	}
	return max
}

func maxProfit1(prices []int) int { // 贪心算法
	if len(prices) == 0 {
		return 0
	}
	leftMin := prices[0]
	result := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < leftMin {
			leftMin = prices[i]
		}
		if prices[i]-leftMin > result {
			result = prices[i] - leftMin
		}
	}
	return result
}
