package Mega

import (
	base "epg_parsers/parser"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type Mega struct {
	base.Common
}

func (m *Mega) GetLocalTime() string {
	return "Europe/Athens"
}

func (m *Mega) Parse(doc *goquery.Document, day int) {
	m.Channels = make(map[string]*base.Channel)
	m.AppendChannel("Mega", &base.Channel{
		Name:      "Mega",
		Icon:      "https://www.tvcreta.gr/wp-content/uploads/2017/02/logo.png",
		Programms: nil,
	})

	prev := ""
	tmp := map[int]base.Programm{}

	doc.Find("body .program__item ").Each(func(i int, selection *goquery.Selection) {
		time := selection.Find("div.program__item-time").Text()
		caption := selection.Find("div.program__item-caption")

		title := caption.Find("div.program__item-caption-text-wrapper span.program__item-caption-title").Text()
		subtitle := caption.Find("div.program__item-caption-text-wrapper span.program__item-caption-subtitle").Text()

		time = strings.TrimSpace(time)

		tmp[i] = base.Programm{
			Timestart:   m.LocalTime.RFC3339local(time, day),
			Timestop:    prev,
			Title:       title,
			Description: subtitle,
		}

		prev = time
	})

	for k, v := range tmp {
		m.AppendProgramm("Mega", base.Programm{
			Timestart:   v.Timestart,
			Timestop:    tmp[k+1].Timestart,
			Title:       v.Title,
			Description: v.Description,
		})
	}
}
