package parsers

import (
	"epg_parsers/parser"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"io/ioutil"
	"strings"
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
		Icon:      "https://program.ert.gr/images/ChannelLogo-ERT1.png",
		Programms: nil,
	})

	a := map[int]map[string]string{}

	doc.Find("table.table > tbody > tr > td.table").Not(".time").Each(func(i int, s *goquery.Selection) {
		sr := strings.NewReader(s.Text())
		tr := transform.NewReader(sr, charmap.Windows1253.NewDecoder())
		buf, err := ioutil.ReadAll(tr)

		href, _ := s.Find("a.black").Attr("href")

		if err != err {
			// handle error
			// default skip
		}

		prepareString := strings.TrimSpace(string(buf))
		prepareString = strings.ReplaceAll(prepareString, "\n", "")

		// dirty items
		a[i] = map[string]string{
			"data": prepareString,
			"meta": href,
		}
	})

	for i := 0; i <= len(a)-1; i++ {
		time := ept.LocalTime.RFC3339local(a[i]["data"], day)
		timestop := ept.LocalTime.RFC3339local(a[i+3]["data"], day) // it is next programm start
		title := a[i+1]["data"]

		// todo get meta from url
		desc := a[i+1]["meta"]

		ept.AppendProgramm("EPT1", base.Programm{
			Timestart:   time,
			Timestop:    timestop,
			Title:       title,
			Description: desc,
		})

		// increnment to next program
		i = i + 2
	}
}
