package payment

type (
	MovieBuyReq struct {
		Email      string   `json:"email"`
		CustomerId string   `json:"customer_id"`
		MovieId    string   `json:"movie_id"`
		Token      string   `json:"token"`
		SeatNo     []string `json:"seat_no"`
		Price      int64    `json:"price"`
		Date       string   `json:"date"`
		Quantity   int64    `json:"quantity"`
	}

	CheckOutWithCreditCard struct {
		Token string `json:"token"`
		Price int64  `json:"price"`
	}

	ReserveSeatReq struct {
		MovieId string   `json:"movie_id"`
		SeatNo  []string `json:"seat_no"`
	}

	AddCustomerTicket struct {
		CustomerId string   `json:"customer_id"`
		Date       string   `json:"date"`
		MovieId    string   `json:"movie_id"`
		SeatNo     []string `json:"seat_no"`
		Quantity   int64    `json:"quantity"`
	}

	PaymentReserveRes struct {
		Id string `json:"id"`
	}
)
