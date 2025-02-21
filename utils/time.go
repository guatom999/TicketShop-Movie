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

func ConvertStringDateToTime(stringDate string) time.Time {

	layout := "02-01-2006"

	t, err := time.Parse(layout, stringDate)

	if err != nil {
		log.Printf("Error parsing string: %s", err.Error())
		return time.Now()
	}

	return t
}

func SetSpecificTime(year int, month time.Month, day, hour, minute, second int) time.Time {

	specificTime := time.Date(year, month, day, hour, minute, second, 0, time.UTC)
	// log.Println("Specific Time is:", specificTime)

	return specificTime

}

func GetStringTime(showTime time.Time) string {

	formattedTime := showTime.Format("2006-01-02:15:04")

	return formattedTime
}
