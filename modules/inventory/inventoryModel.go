package inventory

type (
	AddCustomerTicketReq struct {
		CustomerId  string   `json:"customer_id"`
		OrderNumber string   `json:"order_number"`
		Date        string   `json:"date"`
		MovieName   string   `json:"movie_name"`
		MovieId     string   `json:"movie_id"`
		TicketUrl   string   `json:"ticket_url"`
		SeatNo      []string `json:"seat_no"`
		Quantity    int64    `json:"quantity"`
	}

	CustomerTikcetRes struct {
		MovieId      string   `json:"movie_id" `
		MovieName    string   `json:"movie_name" `
		Ticket_Image string   `json:"ticket_image"  `
		Created_At   string   `json:"created_at"  `
		Price        int64    `json:"price" `
		Seat         []string `json:"seat"  `
	}
)
