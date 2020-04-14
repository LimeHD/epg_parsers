package base

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// main layout of date time string
const RFC3339local = "2006-01-02T15:04:05Z"

type (
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
		baseUrl   string
		LocalTime Time
		Channels  map[string]*Channel `json:"channels"`
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

func (epg *Epg) AppendDay(name string, day *Day) {
	epg.Days[name] = day
}

func (epg *Epg) DayExist(day string) bool {
	_, exist := epg.Days[day]

	return exist
}

type IParse interface {
	BaseUrl() string
	Parse(day int)
	Marshal() string
}

func (common *Common) BootstrapLocalTime(location string) {
	loc, err := time.LoadLocation(location)

	if err != nil {
		panic(err)
	}

	currentTime := time.Now().In(loc)

	common.LocalTime = Time{
		Year:     currentTime.Year(),
		Month:    currentTime.Month(),
		Day:      currentTime.Day(),
		Location: loc,
	}
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

func (common *Common) SetBaseUrl(url string) {
	common.baseUrl = url
}

func (common *Common) BaseUrl() string {
	return common.baseUrl
}

func (d *Day) ToCSV() []string {
	csv := []string{}

	// надо как-то разрулить этот O(N^2) на более высоком уровне дойдет и до куба...
	for _, c := range d.Common.Channels {
		for _, p := range c.Programms {
			csv = append(csv, fmt.Sprintf("%s;%s;%s;%s;%s;%s",
				strings.TrimSpace(d.Name),
				strings.TrimSpace(c.Name),
				strings.TrimSpace(c.Icon),
				strings.TrimSpace(p.Title),
				strings.TrimSpace(p.Timestart),
				strings.TrimSpace(p.Timestop),
			))
		}
	}

	return csv
}
