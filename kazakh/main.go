package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"log"
	"time"
)

const SHEET_NATIVE_AND_RUS = "Сетка на каз,рус,анг языке"
const SHEET_ENG = "Сетка на анг языке"

func main() {
	f, err := excelize.OpenFile("./examples/01.02.2021-07.02.2021.xlsx")

	if err != nil {
		log.Fatal(err)
	}

	rowCount := 0
	rows, err := f.GetRows(SHEET_ENG)

	for _, _ = range rows {
		rowCount++
	}

	var date time.Time
	var date2 time.Time

	getDate := buildDateParser()
	getTimecodes := buildTimcodesParser()
	getValues := buildSheetReaders(f, SHEET_ENG, SHEET_NATIVE_AND_RUS)
	getMultilangs := buildMultilangProgramParser()

	// все, что нужно начинается с 8-ой строки
	for i := 8; i <= rowCount; i++ {
		A, E, AA, DD, err := getValues(i)

		if err != nil {
			log.Fatal(err)
		}

		// если у нас нет временных меток в виде минут и секунд - значит это хидер
		if A == "" {
			//
			// HEADER
			//
			if date, err = Date(getDate(E)); err != nil {
				log.Fatal(err)
			}

			if date2, err = Date(getDate(DD)); err != nil {
				log.Fatal(err)
			}

			fmt.Println("DATE",
				date.String(),
				date2.String(),
			)

			// после хидера с датой идет бесполезная строка со следующим содержимым: 00:00-00:00
			// пропускаем его
			i++
		} else {
			//
			// PROGRAMS
			//
			hours, minutes, err := getTimecodes(A)

			if err != nil {
				log.Fatal(err)
			}

			hours2, minutes2, err := getTimecodes(AA)

			if err != nil {
				log.Fatal(err)
			}

			datetime := AddTime(date, hours, minutes)
			datetime2 := AddTime(date2, hours2, minutes2)

			fmt.Println(
				datetime.String(), FormatDatetime(datetime), E,
			)

			fmt.Println(
				datetime2.String(), FormatDatetime(datetime), getMultilangs(DD),
			)

			fmt.Println("======//")
		}
	}
}
