package main

import (
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"regexp"
	"strconv"
	"strings"
)

func buildDateParser() func(string) string {
	re := regexp.MustCompile(`\d{2}\.\d{2}\.\d{4}`)

	return func(s string) string {
		return re.FindString(strings.TrimSpace(s))
	}
}

func buildTimcodesParser() func(s string) (int, int, error) {
	return func(s string) (int, int, error) {
		times := strings.Split(strings.TrimSpace(s), ":")

		if len(times) < 2 {
			return 0, 0, errors.New("Failed, expect two items")
		}

		hours, err := strconv.Atoi(times[0])

		if err != nil {
			return 0, 0, err
		}

		minutes, err := strconv.Atoi(times[1])

		if err != nil {
			return 0, 0, nil
		}

		return hours, minutes, nil
	}
}

func buildSheetReaders(f *excelize.File, sheet1, sheet2 string) func(i int) (Row, error) {
	// пиздец блять, ебаная либа не могет в 24 формат времени, я просто в ахуе
	// далее творится неописуемый беспредел
	parseTime := func(sheet, c string) (string, error) {
		if err := f.SetCellStyle(sheet, c, c, 0); err != nil {
			return "", err
		}

		value, err := f.GetCellValue(sheet, c)

		if err != nil {
			return "", err
		}

		// это нормальное явление, т.к. ячейка может быстой - это значит что достигли очередного дна, простите хедера (дня)
		if value == "" {
			return "", nil
		}

		f, err := strconv.ParseFloat(value, 64)

		if err != nil {
			return "", err
		}

		t, err := excelize.ExcelDateToTime(f, false)

		if err != nil {
			return "", err
		}

		return fmt.Sprintf("%02d:%02d", t.Hour(), t.Minute()), nil
	}

	return func(i int) (Row, error) {
		A, err := parseTime(sheet1, fmt.Sprintf("A%d", i))

		if err != nil {
			return Row{}, err
		}

		E, err := f.GetCellValue(sheet1, fmt.Sprintf("E%d", i))

		if err != nil {
			return Row{}, err
		}

		AA, err := parseTime(sheet2, fmt.Sprintf("A%d", i))

		if err != nil {
			return Row{}, err
		}

		D, err := f.GetCellValue(sheet2, fmt.Sprintf("D%d", i))

		if err != nil {
			return Row{}, err
		}

		EE, err := parseTime(sheet2, fmt.Sprintf("E%d", i))

		if err != nil {
			return Row{}, err
		}

		return Row{
			SheetOne: Values{
				Time:     A,
				Title:    E,
				Duration: "",
			},
			SheetTwo: Values{
				Time:     AA,
				Title:    D,
				Duration: EE,
			},
		}, nil
	}
}

func buildMultilangProgramParser() func(s string) map[string]string {
	const (
		RU = "RU"
		KZ = "KZ"
		EN = "EN"
	)

	return func(s string) map[string]string {
		items := strings.Split(strings.TrimSpace(s), " / ")
		langs := map[string]string{
			KZ: "", RU: "", EN: "",
		}

		switch len(items) {
		case 2:
			langs[RU] = items[0]
			langs[EN] = items[1]
		case 3:
			langs[KZ] = items[0]
			langs[RU] = items[1]
			langs[EN] = items[2]
		}

		return langs
	}
}
