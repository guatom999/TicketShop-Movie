package payment

type (
	MovieBuyReq struct {
		Email      string   `json:"email"`
		CustomerId string   `json:"customer_id"`
		MovieName  string   `json:"movie_name"`
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
		MovieName string   `json:"movie_name"`
		MovieId   string   `json:"movie_id"`
		SeatNo    []string `json:"seat_no"`
	}

	AddCustomerTicket struct {
		CustomerId string   `json:"customer_id"`
		Date       string   `json:"date"`
		MovieName  string   `json:"movie_name"`
		MovieId    string   `json:"movie_id"`
		TicketUrl  string   `json:"ticket_url"`
		SeatNo     []string `json:"seat_no"`
		Quantity   int64    `json:"quantity"`
	}

	BuyticketRes struct {
		Url string `json:"url"`
	}

	PaymentReserveRes struct {
		Id string `json:"id"`
	}
)
