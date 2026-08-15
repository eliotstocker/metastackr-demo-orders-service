package handlers

type OrderItem struct {
	ID       string  `json:"id"`
	Quantity int     `json:"qty"`
	Price    float64 `json:"price"`
}

func CalculateSubtotal(items []OrderItem) float64 {
	var total float64
	for _, it := range items {
		total += it.Price * float64(it.Quantity)
	}
	return total
}
