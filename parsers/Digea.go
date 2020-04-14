package parsers

import (
	"github.com/LimeHD/parser/base"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type Digea struct {
	base.Common
}

func (digea *Digea) GetLocalTime() string {
	return "Europe/Athens"
}

func (digea *Digea) Parse(doc *goquery.Document, day int) {
	// required
	digea.Channels = make(map[string]*base.Channel)

	doc.Find(".epg-table-row > .col-lg-4").Each(func(i int, s *goquery.Selection) {
		channelName := s.Find(".panel strong").Text()
		icon, _ := s.Find(".panel .epg-icon").Children().Attr("src")

		if !digea.ChannelExist(channelName) {
			digea.AppendChannel(channelName, &base.Channel{
				Name:      channelName,
				Icon:      icon,
				Programms: []base.Programm{},
			})

		}

		// У каналов есть только начало телепередачи
		prev := ""
		s.Find("ul.epg-list > li.list-group-item").Each(func(ii int, ss *goquery.Selection) {
			times := ss.Find("span.time").Text()
			title := ss.Find("span.tv-show").Text()

			times = digea.LocalTime.RFC3339local(strings.TrimSpace(times), day)

			if prev == "" {
				prev = times
			}

			// интересная особенность, передачи могут заканчиваться в то же время, что и начинаются
			// парадоксально, однако
			if prev != times {
				digea.AppendProgramm(channelName, base.Programm{
					Timestart: prev,
					Timestop:  times,
					Title:     title,
				})
			}

			prev = times
		})
	})
}
