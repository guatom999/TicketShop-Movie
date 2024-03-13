package ticket

type (
	AddTikcetReq struct {
		MovieId    string  `json:"movie_id"`
		CustomerId string  `json:"customer_id"`
		MovieName  string  `json:"movie_name"`
		Seat       string  `json:"seat"`
		Date       string  `json:"date"`
		Time       string  `json:"time"`
		Price      float64 `json:"price"`
	}

	TicketShowCase struct {
		Title string  `json:"title"`
		Seat  string  `json:"seat"`
		Date  string  `json:"date"`
		Time  string  `json:"time"`
		Price float64 `json:"price"`
	}
)
