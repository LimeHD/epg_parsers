package main

import (
	"epg_parsers/mega_parser/Mega"
	base "epg_parsers/parser"
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
				Value: "mega.csv",
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
		base.PrintMeta("Mega", "0.0.1", output)
		location, _ := time.LoadLocation("Europe/Athens")
		currentTime := time.Now().In(location)

		for i := 0; i <= 5; i++ {
			key := currentTime.AddDate(0, 0, i).Format("2/1/2006")

			if !epg.DayExist(key) {
				epg.AppendDay(key, &base.Day{
					Name:   key,
					Common: &base.Common{},
				})
			}

			mega := Mega.Mega{}
			mega.SetBaseUrl(fmt.Sprintf("https://www.megatv.com/incl/v5megatvprogram_35895.asp?pageid=956&catid=17496&catidlocal=17496&date=%s&subidlocal=1&ajaxid=35895&ajaxgroup=PROGRAM", key))
			err := parser.RunComposeWith(&mega, i)

			if err != nil {
				_ = bug.Notify(err, bugsnag.MetaData{
					"Parser": {
						"Name": "Mega",
					},
				})

				fmt.Println("Something went wrong... Information sent to bugsnag")
				log.Fatal(err)
			}

			epg.Days[key].Common = &mega.Common
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
