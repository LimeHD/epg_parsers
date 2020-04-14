package main

import (
	"fmt"
	"github.com/LimeHD/parser/base"
	"github.com/LimeHD/parser/parsers"
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
		},
	}

	app.Action = func(c *cli.Context) error {
		format := c.String("format")

		epg := base.Epg{Days: map[string]*base.Day{}}
		parser := &base.Parser{}

		// 5 - это для теста, сколько дней вперед парсить, включая текущий день
		for i := 0; i <= 5; i++ {
			key := time.Now().AddDate(0, 0, i).Format("20060102")

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

		// todo пока что временно так, потом заревакторю
		if format == "csv" {
			for _, v := range epg.Days["20200417"].ToCSV() {
				fmt.Println(v)
			}
		} else {
			fmt.Println(epg.Days["20200417"].Common.Marshal())
		}

		return nil
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
