package main

import (
	"errors"
	"fmt"
	"time"
)

type DateRange struct {
	start time.Time
	end   time.Time
}

func (d DateRange) Hours() float64 {
	return d.end.Sub(d.start).Hours()
}

func NewDateRange(startDate, endDate time.Time) (DateRange, error) {
	if time.Time.IsZero(startDate) || time.Time.IsZero(endDate) {
		return DateRange{}, errors.New("date cannot be empty.")
	}

	if time.Time.Before(endDate, startDate) {
		return DateRange{}, errors.New("enddate cannot be earlier than start date.")
	}

	return DateRange{
		start: startDate,
		end:   endDate,
	}, nil
}

func main() {
	lifetime, _ := NewDateRange(time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC), time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC))

	fmt.Println(lifetime.Hours())

	travelInTime := DateRange{
		start: time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC),
		end:   time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC),
	}

	fmt.Println(travelInTime.Hours())
}
