package code

import "math"

func maxProfit(prices []int) int {
	var minPrice, maxprofit = math.MaxInt32, 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxprofit {
			maxprofit = prices[i] - minPrice
		}
	}
	return maxprofit
}
