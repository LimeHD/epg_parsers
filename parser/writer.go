package base

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

func (d *Day) ToTSV() [][]string {
	var csvArray [][]string

	for _, c := range d.Common.Channels {
		for _, p := range c.Programs {
			csvArray = append(csvArray, []string{
				strings.TrimSpace(p.Timestart),
				strings.TrimSpace(p.Timestop),
				strings.Replace(strings.TrimSpace(c.Name), "\n", "", -1),
				strings.Replace(strings.TrimSpace(p.Title), "\n", "", -1),
				strings.TrimSpace(c.Icon),
				strings.Replace(strings.TrimSpace(p.Description), "\n", "", -1),
			})
		}
	}

	return csvArray
}

func WriteTSV(output string, data map[string]*Day) {
	file, err := os.Create(output)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	writer := csv.NewWriter(file)
	writer.Comma = '\t'
	defer writer.Flush()

	_ = writer.Write([]string{
		"datetime_start",
		"datetime_finish",
		"channel",
		"title",
		"channel_logo_url",
		"description",
	})

	for _, day := range data {
		for _, value := range day.ToTSV() {
			err := writer.Write(value)

			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
