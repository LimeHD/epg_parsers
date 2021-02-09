package main

import (
	"testing"
)

func TestDatetime(t *testing.T) {
	t.Run("it should be give correct date", func(t *testing.T) {
		date, err := Date("01.02.2021")

		if err != nil {
			t.Fatal(err)
		}

		t.Run("it should be successfully convert date format", func(t *testing.T) {
			if FormatDatetime(date) != "01.02.2021 01:00:00" {
				t.Fatal("Failed, expect datetime: `01.02.2021 01:00:00`")
			}
		})

		_, err = Date("01.02.2021 01:00:00")

		if err == nil {
			t.Fatal("Failed, expect error")
		}

		_, err = Date("01-02-2021")

		if err == nil {
			t.Fatal("Failed, expect error")
		}
	})

	t.Run("it should be successfully add times", func(t *testing.T) {
		date, err := Date("07.02.2021")

		if err != nil {
			t.Fatal(err)
		}

		datetime := AddTime(date, 2, 5)

		t.Run("it should be correct give UTC time zone", func(t *testing.T) {
			if datetime.Format(DateTimeFormat) != "07.02.2021 02:05:00" {
				t.Fatalf("Failed, give %s expect `07.02.2021 05:05:00`", datetime.Format(DateTimeFormat))
			}
		})

		t.Run("it should be successfully convert date format", func(t *testing.T) {
			if FormatDatetime(datetime) != "07.02.2021 03:05:00" {
				t.Fatal("Failed, expect datetime: `07.02.2021 03:05:00`")
			}
		})
	})
}
