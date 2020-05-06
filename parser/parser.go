package base

import (
	"epg_parsers/utils"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"time"
)

// main layout of date time string
const RFC3339local = "2006-01-02T15:04:05Z"

type (
	Parser struct{}

	Programm struct {
		Timestart   string `json:"timestart"`
		Timestop    string `json:"timestop"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	Channel struct {
		Name      string     `json:"name"`
		Icon      string     `json:"icon"`
		Programms []Programm `json:"programms"`
	}

	Day struct {
		Name   string
		Common *Common
	}

	Epg struct {
		Days map[string]*Day
	}
)

func (p *Parser) RunComposeWith(parser IParse, day int) error {
	fmt.Println(fmt.Sprintf("Get HTML document from %s", parser.BaseUrl()))

	status, reader := utils.GetHtmlDocumentReader(parser.BaseUrl())
	defer reader.Close()

	if status != 200 {
		return errors.New(fmt.Sprintf("Не могу получить данные, что-то пошло не так, HTTP код: %d", status))
	}

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return err
	}

	err = parser.BootstrapLocalTime(parser.GetLocalTime())
	if err != nil {
		return err
	}

	parser.Parse(doc, day)

	return nil
}

func (epg *Epg) AppendDay(name string, day *Day) {
	epg.Days[name] = day
}

func (epg *Epg) DayExist(day string) bool {
	_, exist := epg.Days[day]

	return exist
}

func PrintMeta(name string, version string, output string) {
	fmt.Println(fmt.Sprintf("Run parser: %s", name))
	fmt.Println(fmt.Sprintf("Parser version: %s", version))
	fmt.Println(fmt.Sprintf("Output to: %s", output))
}

func (t *Time) RFC3339local(times string, day int) string {
	layout := fmt.Sprintf("%d-%02d-%02dT%s:00Z",
		t.Year,
		t.Month,
		t.Day,
		times,
	)

	tt, _ := time.Parse(RFC3339local, layout)
	timeFormatted := tt.AddDate(0, 0, day).In(t.Location).Format(time.RFC3339)

	return timeFormatted
}
