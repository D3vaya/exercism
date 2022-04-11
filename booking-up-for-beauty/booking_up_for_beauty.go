//package booking
package main

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	const (
		layout = "2019-07-25 13:45:00 +0000 UTC"
	)
	t, _ := time.Parse(layout, date)
	fmt.Println(t)

	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	panic("Please implement the HasPassed function")
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	panic("Please implement the IsAfternoonAppointment function")
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	panic("Please implement the Description function")
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	panic("Please implement the AnniversaryDate function")
}

func main() {
	Schedule("7/25/2019 13:45:00")
}
