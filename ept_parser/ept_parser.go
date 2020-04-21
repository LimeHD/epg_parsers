package main

import (
	parsers "epg_parsers/ept_parser/Ept"
	base "epg_parsers/parser"
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
				Value: "ept.csv",
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
		//format := c.String("format")
		output := c.String("output")
		bugsnagApiKey := c.String("bugsnag_api_key")

		bug := bugsnag.New(bugsnag.Configuration{
			APIKey:          bugsnagApiKey,
			AppVersion:      "0.0.1",
			ProjectPackages: []string{"main", "github.com/LimeHD/epg_parsers"},
		})

		epg := base.Epg{Days: map[string]*base.Day{}}
		parser := &base.Parser{}
		base.PrintMeta("Ept", "0.0.1", output)

		// загружаем локальное время
		location, _ := time.LoadLocation("Europe/Athens")
		currentTime := time.Now().In(location)

		// 5 - это для теста, сколько дней вперед парсить, включая текущий день
		for i := 0; i <= 0; i++ {
			key := currentTime.AddDate(0, 0, i).Format("02/01/2006")

			if !epg.DayExist(key) {
				epg.AppendDay(key, &base.Day{
					Name:   key,
					Common: &base.Common{},
				})
			}

			ept := parsers.Ept{}
			ept.SetBaseUrl(fmt.Sprintf("https://program.ert.gr/Ert1/index.asp?id=9&pdate=%s", key))
			err := parser.RunComposeWith(&ept, i)

			if err != nil {
				_ = bug.Notify(err, bugsnag.MetaData{
					"Parser": {
						"Name": "EPT1",
					},
				})

				fmt.Println("Something went wrong... Information sent to bugsnag")
				log.Fatal(err)
			}

			epg.Days[key].Common = &ept.Common
		}

		for _, day := range epg.Days {
			utils.WriteTSV(output, day.ToTSV())
		}

		fmt.Println("Finished for parse & export")

		return nil
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
