package base

import (
	"encoding/json"
	"fmt"
	"strings"
)

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
		baseUrl  string
		Channels map[string]*Channel `json:"channels"`
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
	BaseUrl()
	Parse()
	Marshal() string
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
