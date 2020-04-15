package main

import (
	"fmt"
	"github.com/LimeHD/epg_parsers/base"
	"github.com/LimeHD/epg_parsers/parsers"
	"github.com/LimeHD/epg_parsers/utils"
	"github.com/urfave/cli"
	"log"
	"os"
	"time"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "format",
				Value: "csv",
				Usage: "Export data format",
			},
			&cli.StringFlag{
				Name:  "output",
				Value: "./output",
				Usage: "Export data directory",
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		format := c.String("format")
		output := c.String("output")

		epg := base.Epg{Days: map[string]*base.Day{}}
		parser := &base.Parser{}

		// 5 - это для теста, сколько дней вперед парсить, включая текущий день
		for i := 0; i <= 5; i++ {
			key := time.Now().AddDate(0, 0, i).Format(time.RFC1123Z)

			if !epg.DayExist(key) {
				epg.AppendDay(key, &base.Day{
					Name:   key,
					Common: &base.Common{},
				})
			}

			digea := parsers.Digea{}
			digea.SetBaseUrl(fmt.Sprintf("https://www.digea.gr/EPG?day=%d", i))

			parser.RunComposeWith(&digea, i)
			epg.Days[key].Common = &digea.Common
		}

		// todo пока что временно так, потом зарефакторю
		if format == "csv" {
			// todo run with sync & goroutines
			for _, day := range epg.Days {
				utils.WriteCSV(output, "digea", day.ToCSV())
			}

			fmt.Println("Finished for parse & export")
		}

		return nil
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
