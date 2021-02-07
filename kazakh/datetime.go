package main

import (
	"fmt"
	"time"
)

const (
	DateTimeFormatRfc = time.RFC3339
	DateTimeFormat    = "01.02.2006 15:04:05"
	DateFormat        = "01.02.2006"
)

func Date(date string) (time.Time, error) {
	_date, err := time.Parse(DateTimeFormat, fmt.Sprintf("%s 00:00:00", date))

	if err != nil {
		return time.Time{}, err
	}

	return _date, nil
}

func AddTime(date time.Time, hours, minutes int) time.Time {
	return date.Add(time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute)
}

func FormatDatetime(datetime time.Time) string {
	location, _ := time.LoadLocation("Europe/Moscow")

	return datetime.In(location).Format(DateTimeFormat)
}
