package main

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"log"
	"testing"
	"time"
)

func TestMainParser(t *testing.T) {
	f, err := excelize.OpenFile("./examples/01.02.2021-07.02.2021.xlsx")

	if err != nil {
		log.Fatal(err)
	}

	rowCount := 0

	t.Run("it should give 344 rows", func(t *testing.T) {
		rows, err := f.GetRows(SHEET_ENG)

		if err != nil {
			t.Fatal(err)
		}

		for _, _ = range rows {
			rowCount++
		}

		if rowCount != 344 {
			t.Fatalf("Failed, give %d, expected 344", rowCount)
		}
	})

	location, _ := time.LoadLocation("Europe/Moscow")

	var date time.Time

	getDate := buildDateParser()
	getTimecodes := buildTimcodesParser()
	getValues := buildSheetReaders(f, SHEET_ENG, SHEET_NATIVE_AND_RUS)

	for i := 8; i <= rowCount; i++ {
		A, _, AA, DD, err := getValues(i)

		if err != nil {
			t.Fatal(err)
		}

		if A == "" {
			if date, err = Date(getDate(DD)); err != nil {
				t.Fatal(err)
			}

			i++
		} else {
			hours, minutes, err := getTimecodes(AA)

			if err != nil {
				t.Fatal(err)
			}

			datetimeUTC := AddTime(date, hours, minutes)
			diff := time.Duration(3) * time.Hour

			if datetimeUTC.Add(diff).Sub(datetimeUTC.In(location)) != diff {
				t.Fatal("Failed UTC converting")
			}

			if datetimeUTC.In(location).After(datetimeUTC) {
				t.Fatal("Failed UTC converting")
			}
		}
	}
}
