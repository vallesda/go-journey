func maxProfit(prices []int) int {
    if len(prices) == 2 { return 0 }
    
    maxProfit := 0
    priceToBuy := 0
    for price := range prices {
        if price < priceToBuy {
            priceToBuy = price 
        } else {
            maxProfit = Max(price, priceToBuy)
        }
    }
    return maxProfit
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}