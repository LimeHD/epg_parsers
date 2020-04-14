package main

import (
	"fmt"
	"github.com/LimeHD/parser/base"
	"github.com/LimeHD/parser/parsers"
)

func main() {
	epg := base.Epg{Days: map[string]*base.Day{}}
	parser := &base.Parser{}

	// 5 - это для теста, сколько дней вперед парсить, включая текущий день
	for i := 0; i <= 5; i++ {
		key := fmt.Sprintf("Day - %d", i)

		if !epg.DayExist(key) {
			epg.AppendDay(key, &base.Day{
				Name:   key,
				Common: &base.Common{},
			})
		}

		digea := parsers.Digea{}
		digea.SetBaseUrl(fmt.Sprintf("https://www.digea.gr/EPG?day=%d", i))

		parser.RunComposeWith(&digea, i)
		epg.Days[key].Common = &digea.Common
	}

	for _, v := range epg.Days["Day - 3"].ToCSV() {
		fmt.Println(v)
	}
}
