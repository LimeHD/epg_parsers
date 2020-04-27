package main

import (
	base "epg_parsers/parser"
	"epg_parsers/star_parser/Star"
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
				Value: "star.csv",
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
		output := c.String("output")
		bugsnagApiKey := c.String("bugsnag_api_key")

		bug := bugsnag.New(bugsnag.Configuration{
			APIKey:          bugsnagApiKey,
			AppVersion:      "0.0.1",
			ProjectPackages: []string{"main", "github.com/LimeHD/epg_parsers"},
		})

		epg := base.Epg{Days: map[string]*base.Day{}}
		parser := &base.Parser{}
		base.PrintMeta("Star", "0.0.1", output)
		location, _ := time.LoadLocation("Europe/Athens")
		currentTime := time.Now().In(location)

		for i := 0; i <= 5; i++ {
			key := currentTime.AddDate(0, 0, i).Format("2006-01-02")

			if !epg.DayExist(key) {
				epg.AppendDay(key, &base.Day{
					Name:   key,
					Common: &base.Common{},
				})
			}

			ept := Star.Star{}
			ept.SetBaseUrl(fmt.Sprintf("https://www.star.gr/tv/ajax/Atcom.Sites.StarTV.Components.Schedule.Detail?dateRequest=%sT00:00:00&la=1&ajax=true&view=AjaxDetail", key))
			err := parser.RunComposeWith(&ept, i)

			if err != nil {
				_ = bug.Notify(err, bugsnag.MetaData{
					"Parser": {
						"Name": "Star",
					},
				})

				fmt.Println("Something went wrong... Information sent to bugsnag")
				log.Fatal(err)
			}

			epg.Days[key].Common = &ept.Common
		}

		base.WriteTSV(output, epg.Days)

		fmt.Println("Finished for parse & export")
		return nil
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
