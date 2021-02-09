package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
	"time"
)

func (d *Day) ToTSV() [][]string {
	var csvArray [][]string

	for _, p := range d.Programs {
		csvArray = append(csvArray, []string{
			strings.TrimSpace(p.Timestart),
			strings.TrimSpace(p.Timestop),
			strings.Replace(strings.TrimSpace("Kazakh"), "\n", "", -1),
			strings.Replace(strings.TrimSpace(p.Title), "\n", "", -1),
			strings.TrimSpace("kazakh-icon"),
			strings.Replace(strings.TrimSpace(p.Description), "\n", "", -1),
		})
	}

	return csvArray
}

func WriteTSV(output string, data map[time.Time]*Day) error {
	file, err := os.Create(output)

	defer func() {
		_ = file.Close()
	}()

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
				return err
			}
		}
	}

	return nil
}
