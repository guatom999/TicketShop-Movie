package ticket

type (
	AddTikcetReq struct {
		MovidId    string  `json:"movid_id"`
		CustomerId string  `json:"customer_id"`
		Title      string  `json:"title"`
		Price      float64 `json:"price"`
		Seat       string  `json:"seat"`
		Date       string  `json:"date"`
		Time       string  `json:"time"`
	}
)
