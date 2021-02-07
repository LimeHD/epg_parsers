package main

import "testing"

func TesDatetime(t *testing.T) {
	t.Run("it should be give correct date", func(t *testing.T) {
		date, err := Date("01.02.2021")

		if err != nil {
			t.Fatal(err)
		}

		t.Run("it should be successfully convert date format", func(t *testing.T) {
			if FormatDatetime(date) != "01.02.2021 00:00:00" {
				t.Fatal("Failed, expect datetime: `01.02.2021 00:00:00`")
			}
		})

		_, err = Date("01.02.2021 01:00:00")

		if err != nil {
			t.Fatal(err)
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

		datetime := AddTime(date, 2, 5).UTC()

		t.Run("it should be successfully convert date format", func(t *testing.T) {
			if FormatDatetime(datetime) != "07.02.2021 02:05:00" {
				t.Fatal("Failed, expect datetime: `07.02.2021 02:05:00`")
			}
		})
	})
}
