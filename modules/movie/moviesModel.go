package movie

type (
	AddMovieReq struct {
		Title           string  `json:"title"`
		Description     string  `json:"description" `
		RunningTime     string  `json:"running_time"`
		Price           float64 `json:"price"`
		ImageUrl        string  `json:"image_url"`
		Avaliable       int     `json:"avaliable"`
		ReleaseAt       string  `json:"release_at"`
		OutOfTheatersAt string  `json:"out_of_theaters_at"`
	}

	AddMovieShowtime struct {
	}

	MovieData struct {
		MovieId string `json:"movie_id"`
		Title   string `json:"title"`
		// Description string  `json:"description" `
		// RunningTime string  `json:"running_time"`
		Price     float64 `json:"price" bson:"price"`
		ImageUrl  string  `json:"image_url" bson:"image_url"`
		Avaliable int     `json:"valiable" bson:"avaliable"`
	}

	MovieShowCase struct {
		Title       string  `json:"title"`
		Description string  `json:"description" `
		RunningTime string  `json:"running_time"`
		Price       float64 `json:"price" bson:"price"`
		ImageUrl    string  `json:"image_url" bson:"image_url"`
		Avaliable   int     `json:"valiable" bson:"avaliable"`
	}

	MovieSearchReq struct {
		Title     string  `json:"title"`
		Price     float64 `json:"price" `
		Avaliable int     `json:"avaliable" `
	}

	MovieShowTimeRes struct {
		Movie_id      string          `json:"movie_id"`
		Title         string          `json:"title"`
		ShowTime      string          `json:"show_time"`
		SeatAvailable []SeatAvailable `json:"seat_available"`
	}

	ReserveDetailReq struct {
		MovieId string   `json:"movie_id"`
		SeatNo  []string `json:"seat_no"`
	}

	ReserveSeatReqTest struct {
		MovieId     string   `json:"movie_id"`
		Seat_Number []string `json:"seat_no"`
	}
)
