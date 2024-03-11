package ticket

type (
	AddTikcetReq struct {
		MovieId    string `json:"movie_id"`
		CustomerId string `json:"customer_id"`
		Seat       string `json:"seat"`
	}
)
