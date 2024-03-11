package payment

type (
	MovieBuyReq struct {
		MovieId string `json:"movie_id"`
		Seat    string `json:"seat"`
	}
)
