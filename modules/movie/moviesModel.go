package movie

type (
	AddMovieReq struct {
		Title     string  `json:"title"`
		Price     float64 `json:"price"`
		ImageUrl  string  `json:"image_url"`
		Avaliable int     `json:"avaliable"`
	}
)
