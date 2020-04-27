package Star

import (
	base "epg_parsers/parser"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type Star struct {
	base.Common
}

func (s *Star) GetLocalTime() string {
	return "Europe/Athens"
}

func (s *Star) Parse(doc *goquery.Document, day int) {
	s.Channels = make(map[string]*base.Channel)
	s.AppendChannel("Star", &base.Channel{
		Name:      "Star",
		Icon:      "https://www.star.gr/tv/faviconnew.ico",
		Programms: nil,
	})

	prev := ""
	tmp := map[int]base.Programm{}

	doc.Find("div.list > div.row").Each(func(i int, selection *goquery.Selection) {
		// mark now as "ΤΩΡΑ" label
		time := strings.TrimSpace(selection.Find(".time").Children().First().Text())
		item := selection.Find(".item")

		title := strings.TrimSpace(item.Find("h3 > a").Text())
		desc := strings.TrimSpace(item.Find("h5").Text())

		tmp[i] = base.Programm{
			Timestart:   s.LocalTime.RFC3339local(time, day),
			Timestop:    prev,
			Title:       title,
			Description: desc,
		}

		prev = time
	})

	for k, v := range tmp {
		s.AppendProgramm("Star", base.Programm{
			Timestart:   v.Timestart,
			Timestop:    tmp[k+1].Timestart,
			Title:       v.Title,
			Description: v.Description,
		})
	}
}
