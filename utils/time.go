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

func SetSpecificTime(year int, month time.Month, day, hour, minute, second int) time.Time {

	specificTime := time.Date(year, month, day, hour, minute, second, 0, time.UTC)
	log.Println("Specific Time is:", specificTime)

	return specificTime

}

func GetStringTime(showTime time.Time) string {

	// fmt.Println("movieShowTime is :", showTime)

	formattedTime := showTime.Format("2006-01-02:15:04")

	return formattedTime
}
