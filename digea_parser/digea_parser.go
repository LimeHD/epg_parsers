package main

import (
	"epg_parsers/digea_parser/Digea"
	"epg_parsers/parser"
	"epg_parsers/utils"
	"fmt"
	"github.com/bugsnag/bugsnag-go"
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
				Value: "digea.csv",
				Usage: "Export data directory",
			},

			&cli.StringFlag{
				Name:  "bugsnag_api_key",
				Value: "",
				Usage: "Export data directory",
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		format := c.String("format")
		output := c.String("output")
		bugsnagApiKey := c.String("bugsnag_api_key")

		bug := bugsnag.New(bugsnag.Configuration{
			APIKey:          bugsnagApiKey,
			AppVersion:      "0.0.1",
			ProjectPackages: []string{"main", "github.com/LimeHD/epg_parsers"},
		})

		epg := base.Epg{Days: map[string]*base.Day{}}
		parser := &base.Parser{}
		base.PrintMeta("Digea", "0.0.1", output)

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

			err := parser.RunComposeWith(&digea, i)

			if err != nil {
				_ = bug.Notify(err, bugsnag.MetaData{
					"Parser": {
						"Name": "Digea",
					},
				})

				fmt.Println("Something went wrong... Information sent to bugsnag")
				log.Fatal(err)
			}

			epg.Days[key].Common = &digea.Common
		}

		// todo пока что временно так, потом зарефакторю
		if format == "csv" {
			// todo run with sync & goroutines
			for _, day := range epg.Days {
				utils.WriteTSV(output, day.ToTSV())
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
