package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	// ------------------------------------------------------
	// DATE & TIME HANDLING IN GO (time package)
	// ------------------------------------------------------

	// time.Now() returns the current local date and time
	presentTime := time.Now()
	fmt.Println(presentTime)

	// Format() is used to format date/time in Go
	// IMPORTANT:
	// Go does NOT use YYYY-MM-DD formatting
	// It uses a reference date:
	// "01-02-2006 15:04:05 Monday"
	//
	// Meaning of reference values:
	// 01 → Month
	// 02 → Day
	// 2006 → Year
	// 15 → Hour (24-hour)
	// 04 → Minute
	// 05 → Second
	// Monday → Weekday
	fmt.Println(presentTime.Format("01-02-2022 12:03:23 Monday"))

	// time.Date() is used to create a custom date and time
	// Parameters:
	// Year, Month, Day, Hour, Minute, Second, Nanosecond, Location
	createdDate := time.Date(
		2020,        // Year
		time.August, // Month
		10,          // Day
		23,          // Hour
		23,          // Minute
		0,           // Second
		0,           // Nanosecond
		time.UTC,    // Time zone
	)

	fmt.Println(createdDate)

	// Format the custom date using the Go reference layout
	fmt.Println(createdDate.Format("01-02-2023 Monday"))
}
