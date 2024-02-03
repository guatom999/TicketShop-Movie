package book

import "time"

type (
	BookingHistory struct {
		CustomerId string    `bson:"customer_id"`
		MovieName  string    `bson:"movie_name"`
		Quantity   int       `bson:"quantity"`
		Price      int       `bson:"price"`
		BookingAt  time.Time `bson:"booking_at"`
		Seat       []Seat    `bson:"seat"`
		ShowTime   string    `bson:"show_time"`
		CreateAt   time.Time `bson:"create_at"`
	}
)
