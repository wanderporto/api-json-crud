package utils

import (
	"fmt"
	"time"
)

func GetDateTime() string {
	today := time.Now()
	day := today.Day()
	month := today.Month()
	year := today.Year()
	hour := today.Hour()
	minutes := today.Minute()
	seconds := today.Second()

	return fmt.Sprintf("%d-%d-%d T%d:%d:%d", day, month, year, hour, minutes, seconds)
}
