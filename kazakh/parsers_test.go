package main

import (
	"testing"
)

func TestParsers(t *testing.T) {
	t.Run("it should be successfully parse date", func(t *testing.T) {
		dates := map[string]string{
			"01.02.2021 Ok":      "01.02.2021",
			"02-02-2021":         "",
			"OK 01.03.2021":      "01.03.2021",
			"01-02.2021 not ok":  "",
			"ok 02.03.2021 oook": "02.03.2021",
		}

		getDate := buildDateParser()

		for date, expected := range dates {
			d := getDate(date)

			if d != expected {
				t.Fatalf("Failed, give %s expected %s", d, expected)
			}
		}
	})

	t.Run("it should be successfully parse timecodes", func(t *testing.T) {
		timecodes := map[string]bool{
			"10:_25":   true,
			"10:25s":   true,
			"_10:35":   false,
			"11:25":    true,
			"10::25":   true,
			"10:25:02": true,
		}

		getTimecodes := buildTimcodesParser()

		for timecode, state := range timecodes {
			_, _, err := getTimecodes(timecode)

			if err != nil && state {
				t.Fatal("Failed, expected true")
			} else if err == nil && state == false {
				t.Fatal("Failed, expected false")
			}
		}
	})
}
