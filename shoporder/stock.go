package shoporder

type Stock struct {
	Total int
}

// NewStock initiates stock struct to hold total stock
func NewStock(stock int) Stock {
	return Stock{
		Total: stock,
	}
}

// Order decreases stock by one.
func (s Stock) Order() (int, string, string) {
	status := "success"
	message := "order successfully created, please proceed to payment!"
	if s.Total <= 0 {
		status = "error"
		message = "stock is not found!"
	} else {
		s.Total--
	}

	return s.Total, status, message
}
