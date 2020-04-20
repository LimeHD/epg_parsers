package base

import (
	"encoding/json"
	"epg_parsers/utils"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"time"
)

// main layout of date time string
const RFC3339local = "2006-01-02T15:04:05Z"

type (
	Parser struct{}

	Programm struct {
		Timestart string `json:"timestart"`
		Timestop  string `json:"timestop"`
		Title     string `json:"title"`
	}

	Channel struct {
		Name      string     `json:"name"`
		Icon      string     `json:"icon"`
		Programms []Programm `json:"programms"`
	}

	Common struct {
		baseUrl       string
		BaseName      string
		Version       string
		localTimeBase string
		LocalTime     Time
		Channels      map[string]*Channel `json:"channels"`
	}

	Time struct {
		Location *time.Location
		Year     int
		Month    time.Month
		Day      int
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

type IParse interface {
	BaseUrl() string
	Parse(doc *goquery.Document, day int)
	Marshal() string
	BootstrapLocalTime(local string) error
	GetLocalTime() string
}

func PrintMeta(name string, version string, output string) {
	fmt.Println(fmt.Sprintf("Run parser: %s", name))
	fmt.Println(fmt.Sprintf("Parser version: %s", version))
	fmt.Println(fmt.Sprintf("Output to: %s", output))
}

func (common *Common) SetLocalTime(local string) *Common {
	common.localTimeBase = local

	return common
}

func (common *Common) GetLocalTime() string {
	if common.localTimeBase == "" {
		panic("Not initialize local time")
	}

	return common.localTimeBase
}

func (common *Common) BootstrapLocalTime(location string) error {
	loc, err := time.LoadLocation(location)

	if err != nil {
		return err
	}

	currentTime := time.Now().In(loc)

	common.LocalTime = Time{
		Year:     currentTime.Year(),
		Month:    currentTime.Month(),
		Day:      currentTime.Day(),
		Location: loc,
	}

	return nil
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

func (common *Common) AppendProgramm(name string, programm Programm) {
	common.Channels[name].Programms = append(common.Channels[name].Programms, programm)
}

func (common *Common) AppendChannel(name string, channel *Channel) {
	common.Channels[name] = channel
}

func (common *Common) ChannelExist(name string) bool {
	_, exist := common.Channels[name]

	return exist
}

func (common *Common) Marshal() string {
	jsonString, err := json.Marshal(common.Channels)

	if err != nil {
		panic(err)
	}

	return string(jsonString)
}

func (common *Common) SetBaseUrl(url string) *Common {
	common.baseUrl = url

	return common
}

func (common *Common) BaseUrl() string {
	return common.baseUrl
}

func (d *Day) ToTSV() []string {
	csv := []string{}
	description := ""

	// надо как-то разрулить этот O(N^2) на более высоком уровне дойдет и до куба...
	for _, c := range d.Common.Channels {
		for _, p := range c.Programms {
			csv = append(csv, fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s",
				strings.TrimSpace(p.Timestart),
				strings.TrimSpace(p.Timestop),
				strings.TrimSpace(c.Name),
				strings.TrimSpace(p.Title),
				strings.TrimSpace(c.Icon),
				description,
			))
		}
	}

	return csv
}
