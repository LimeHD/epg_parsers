package parsers

import (
	"github.com/LimeHD/parser/base"
	"github.com/LimeHD/parser/utils"
	"github.com/PuerkitoBio/goquery"
	"log"
)

type Digea struct {
	base.Common
}

func (digea *Digea) BaseUrl() string {
	return "https://www.digea.gr/EPG/el"
}

func (digea *Digea) Parse() {
	status, reader := utils.GetHtmlDocumentReader(digea.BaseUrl())
	defer reader.Close()

	if status != 200 {
		log.Fatalf("Не могу получить данные, что-то пошло не так, HTTP код: %d", status)
	}

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
	}

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
			time := ss.Find("span.time").Text()
			title := ss.Find("span.tv-show").Text()

			if prev == "" {
				prev = time
			}

			// интересная особенность, передачи могут заканчиваться в то же время, что и начинаются
			// парадоксально, однако
			if prev != time {
				digea.AppendProgramm(channelName, base.Programm{
					Timestart: prev,
					Timestop:  time,
					Title:     title,
				})
			}

			prev = time
		})
	})
}
