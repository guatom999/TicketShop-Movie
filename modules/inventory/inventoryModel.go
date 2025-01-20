package inventory

type (
	AddCustomerTicketReq struct {
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

	CustomerTicketDetail struct {
		TicketId     string   `json:"ticket_id"`
		OrderNumber  string   `json:"order_number"`
		MovieId      string   `json:"movie_id" `
		MovieName    string   `json:"movie_name" `
		MovieImage   string   `json:"movie_image"`
		Ticket_Image string   `json:"ticket_image"`
		Created_At   string   `json:"created_at"  `
		Price        int64    `json:"price" `
		Seat         []string `json:"seat"  `
	}

	CustomerTikcetRes struct {
		TicketId      string   `json:"ticket_id"`
		OrderNumber   string   `json:"order_number"`
		MovieId       string   `json:"movie_id" `
		MovieName     string   `json:"movie_name" `
		MovieImage    string   `json:"movie_image"`
		MovieDate     string   `json:"movie_date"`
		MovieShowTime string   `json:"movie_show_time"`
		Ticket_Image  string   `json:"ticket_image"  `
		Created_At    string   `json:"created_at"  `
		Price         int64    `json:"price" `
		Seat          []string `json:"seat"  `
	}
)
