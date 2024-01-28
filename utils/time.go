package utils

import (
	"log"
	"time"
)

func GetLocaltime() time.Time {
	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Printf("Error: Load Localtime Failed")
		panic(err)
	}
	return time.Now().In(loc)
}
