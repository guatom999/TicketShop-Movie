package payment

type (
	MovieBuyReq struct {
		Email         string   `json:"email"`
		CustomerId    string   `json:"customer_id"`
		MovieName     string   `json:"movie_name"`
		MovieId       string   `json:"movie_id"`
		MovieDate     string   `json:"movie_date"`
		MovieShowTime string   `json:"movie_showtime"`
		PosterImage   string   `json:"poster_image"`
		Token         string   `json:"token"`
		SeatNo        []string `json:"seat_no"`
		Price         int64    `json:"price"`
		Date          string   `json:"date"`
		Quantity      int64    `json:"quantity"`
	}

	CheckOutWithCreditCard struct {
		Token string `json:"token"`
		Price int64  `json:"price"`
	}

	ReserveSeatReq struct {
		MovieName string   `json:"movie_name"`
		MovieId   string   `json:"movie_id"`
		SeatNo    []string `json:"seat_no"`
		Error     string   `json:"error"`
	}

	RollBackReservedSeatReq struct {
		MovieId string   `json:"movie_id"`
		SeatNo  []string `json:"seat_no"`
		Error   string   `json:"error"`
	}

	RollBackReserveSeatRes struct {
		MovieId     string   `json:"movie_id"`
		Seat_Number []string `json:"seat_no"`
		Error       string   `json:"error"`
	}

	BuyticketRes struct {
		TransactionId string `json:"transaction_id"`
		Url           string `json:"url"`
	}

	PaymentReserveRes struct {
		Id string `json:"id"`
	}
)
