package main

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

const SHEET_NATIVE_AND_RUS = "Сетка на каз,рус,анг языке"
const SHEET_ENG = "Сетка на анг языке"

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "input",
				Usage:    "Filepath to .xlsx file, ./examples/01.02.2021-07.02.2021.xlsx",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "output",
				Usage: "Output CSV file",
				Value: "./tmp/output.tsv",
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		f, err := excelize.OpenFile(c.String("input"))

		if err != nil {
			return err
		}

		rowCount := 0
		rows, err := f.GetRows(SHEET_NATIVE_AND_RUS)

		for _, _ = range rows {
			rowCount++
		}

		var date time.Time

		getDate := buildDateParser()
		getTimecodes := buildTimcodesParser()
		getValues := buildSheetReaders(f, SHEET_ENG, SHEET_NATIVE_AND_RUS)
		getMultilangs := buildMultilangProgramParser()

		epg := Epg{
			Channel: "Kazakh",
			Days:    map[time.Time]*Day{},
		}

		// все, что нужно начинается с 8-ой строки
		for i := 8; i <= rowCount; i++ {
			row, err := getValues(i)

			if err != nil {
				return err
			}

			// если у нас нет временных меток в виде минут и секунд - значит это хидер
			if row.SheetTwo.Time == "" {
				if date, err = Date(getDate(row.SheetTwo.Title)); err != nil {
					return err
				}

				epg.Days[date] = &Day{
					Programs: []*Program{},
				}

				// после хидера с датой идет бесполезная строка со следующим содержимым: 00:00-00:00
				// пропускаем его
				i++
			} else {
				hours, minutes, err := getTimecodes(row.SheetTwo.Time)

				if err != nil {
					return err
				}

				hours2, minutes2, err := getTimecodes(row.SheetTwo.Duration)

				if err != nil {
					return err
				}

				timestart := AddTime(date, hours, minutes)
				timestop := AddTime(timestart, hours2, minutes2)

				langs := getMultilangs(row.SheetTwo.Title)

				epg.Days[date].Programs = append(epg.Days[date].Programs, &Program{
					Timestart:   FormatDatetime(timestart),
					Timestop:    FormatDatetime(timestop),
					Title:       priority(langs),
					Description: "",
				})
			}
		}

		return WriteTSV(c.String("output"), epg.Days)
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func priority(l map[string]string) string {
	const (
		RU = "RU"
		KZ = "KZ"
		EN = "EN"
	)

	if l[EN] != "" {
		return l[EN]
	}

	if l[RU] != "" {
		return l[RU]
	}

	return l[KZ]
}
