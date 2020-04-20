package parsers

import (
	"epg_parsers/parser"
	"github.com/PuerkitoBio/goquery"
)

type Ept struct {
	base.Common
}

func (digea *Ept) GetLocalTime() string {
	return "Europe/Athens"
}

func (ept *Ept) Parse(doc *goquery.Document, day int) {
	ept.Channels = make(map[string]*base.Channel)
	ept.AppendChannel("EPT1", &base.Channel{
		Name:      "EPT1",
		Icon:      "",
		Programms: nil,
	})

	doc.Find("table.table > tbody > tr > td.table").Not(".time").Each(func(i int, s *goquery.Selection) {

	})
}
