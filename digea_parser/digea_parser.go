package main

import (
	"epg_parsers/digea_parser/Digea"
	"epg_parsers/parser"
	"epg_parsers/utils"
	"fmt"
	"github.com/blacked/go-zabbix"
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

			&cli.StringFlag{
				Name:  "servrice_host",
				Value: "",
				Usage: "",
			},
			&cli.IntFlag{
				Name:  "zabbix_port",
				Value: 10051,
				Usage: "",
			},
			&cli.StringFlag{
				Name:  "zabbix_server",
				Value: "",
				Usage: "",
			},
			&cli.StringFlag{
				Name:  "zabbix_metric",
				Value: "",
				Usage: "",
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		format := c.String("format")
		output := c.String("output")
		bugsnagApiKey := c.String("bugsnag_api_key")
		sHost := c.String("service_host")
		zPort := c.Int("zabbix_port")
		zServer := c.String("zabbix_server")
		zMetric := c.String("zabbix_metric")

		bug := bugsnag.New(bugsnag.Configuration{
			APIKey:          bugsnagApiKey,
			AppVersion:      "0.0.1",
			ProjectPackages: []string{"main", "github.com/LimeHD/epg_parsers"},
		})

		metrics := []*zabbix.Metric{}

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
				utils.WriteCSV(output, day.ToCSV())
			}

			fmt.Println("Finished for parse & export")
		}

		packet := zabbix.NewPacket(metrics)
		metrics = append(metrics, zabbix.NewMetric(zServer, "service", sHost))
		metrics = append(metrics, zabbix.NewMetric(zServer, zMetric, "OK"))
		z := zabbix.NewSender(zServer, zPort)
		_, err := z.Send(packet)

		if err != nil {
			_ = bug.Notify(err, bugsnag.MetaData{
				"Parser": {
					"Name": "Digea",
				},
			})
		}

		return nil
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
