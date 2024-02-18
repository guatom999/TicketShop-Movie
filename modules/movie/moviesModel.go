package movie

type (
	AddMovieReq struct {
		Title     string  `json:"title"`
		Price     float64 `json:"price"`
		ImageUrl  string  `json:"image_url"`
		Avaliable int     `json:"avaliable"`
	}

	MovieData struct {
		MovieId   string  `json:"movie_id"`
		Title     string  `json:"title"`
		Price     float64 `json:"price" bson:"price"`
		ImageUrl  string  `json:"image_url" bson:"image_url"`
		Avaliable int     `json:"valiable" bson:"avaliable"`
	}

	MovieShowCase struct {
		Title     string  `json:"title"`
		Price     float64 `json:"price" bson:"price"`
		ImageUrl  string  `json:"image_url" bson:"image_url"`
		Avaliable int     `json:"valiable" bson:"avaliable"`
	}

	MovieSearchReq struct {
		Title     string  `json:"title"`
		Price     float64 `json:"price" `
		Avaliable int     `json:"avaliable" `
	}

	MovieShowTimeRes struct {
		Title         string           `json:"title"`
		ShowTime      string           `json:"show_time"`
		SeatAvailable []*SeatAvailable `json:"seat_available"`
	}
)
