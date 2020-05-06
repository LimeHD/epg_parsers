package parsers

import (
	"epg_parsers/parser"
	"fmt"
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
		icon = fmt.Sprintf("https://www.digea.gr%s", icon)

		if !digea.ChannelExist(channelName) {
			digea.AppendChannel(channelName, &base.Channel{
				Name:      channelName,
				Icon:      icon,
				Programms: []base.Programm{},
			})

		}

		// У каналов есть только начало телепередачи
		timestart := ""
		s.Find("ul.epg-list > li.list-group-item").Each(func(ii int, ss *goquery.Selection) {
			timestop := ss.Find("span.time").Text()
			title := ss.Find("span.tv-show").Text()

			_safeTimestop := timestop
			timestop = digea.LocalTime.RFC3339local(strings.TrimSpace(timestop), day)

			if timestart == "" {
				timestart = timestop
			}

			if timestop < timestart {
				// tomorrow
				timestop = digea.LocalTime.RFC3339local(strings.TrimSpace(_safeTimestop), day+1)
			}

			// интересная особенность, передачи могут заканчиваться в то же время, что и начинаются
			// парадоксально, однако
			if timestart != timestop {
				if timestop < timestart {
					digea.HandleWrongTime(true, channelName, title, timestart, timestop)
				}

				digea.AppendProgramm(channelName, base.Programm{
					Timestart: timestart,
					Timestop:  timestop,
					Title:     title,
				})
			}

			timestart = timestop
		})
	})
}
