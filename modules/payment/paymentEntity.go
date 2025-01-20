package payment

type (
	AddCustomerTicket struct {
		CustomerId    string   `json:"customer_id"`
		OrderNumber   string   `json:"order_number"`
		Date          string   `json:"date"`
		MovieName     string   `json:"movie_name"`
		MovieId       string   `json:"movie_id"`
		MovieDate     string   `json:"movie_date"`
		MovieShowTime string   `json:"movie_show_time"`
		PosterImage   string   `json:"poster_image"`
		TicketUrl     string   `json:"ticket_url"`
		SeatNo        []string `json:"seat_no"`
		Quantity      int64    `json:"quantity"`
		Price         int64    `json:"price"`
	}
)
