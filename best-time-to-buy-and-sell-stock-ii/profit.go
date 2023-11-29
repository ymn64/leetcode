package main

import "fmt"

// Whenever there is a rise in price, we buy yesterday and sell today.
// We can do this on consecutive days because we are allowed to buy and
// sell on the same day.
func maxProfit(prices []int) int {
	profit := 0

	for i := 1; i < len(prices); i++ {
		dailyProfit := prices[i] - prices[i-1]

		if dailyProfit > 0 {
			profit += dailyProfit
		}
	}

	return profit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}), 7)
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}), 4)
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}), 0)
	fmt.Println(maxProfit([]int{7}), 0)
}
