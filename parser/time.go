package base

import (
	"fmt"
	"time"
)

// main layout of date time string
const RFC3339local = "2006-01-02T15:04:05Z"

type Time struct {
	Location *time.Location
	Year     int
	Month    time.Month
	Day      int
}

// method create local date time as RFC3339 format through parser location
func (t *Time) RFC3339local(times string, day int) string {
	layout := fmt.Sprintf("%d-%02d-%02dT%s:00Z",
		t.Year,
		t.Month,
		t.Day,
		times,
	)

	tt, _ := time.Parse(RFC3339local, layout)
	timeFormatted := tt.AddDate(0, 0, day).In(t.Location).Format(time.RFC3339)

	return timeFormatted
}
