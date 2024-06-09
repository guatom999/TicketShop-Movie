package inventory

type (
	AddCustomerTicketReq struct {
		CustomerId string   `json:"customer_id"`
		Date       string   `json:"date"`
		MovieId    string   `json:"movie_id"`
		SeatNo     []string `json:"seat_no"`
		Quantity   int64    `json:"quantity"`
	}
)
