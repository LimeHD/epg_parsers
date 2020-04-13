package parsers

import (
	"github.com/LimeHD/parser/base"
	"github.com/LimeHD/parser/utils"
	"github.com/PuerkitoBio/goquery"
	"log"
)

type Ept struct {
	base.Common
}

func (digea *Ept) BaseUrl() string {
	return "https://www.digea.gr/EPG/el"
}

func (ept *Ept) Parse() {
	status, reader := utils.GetHtmlDocumentReader("https://program.ert.gr/Ert1/")
	defer reader.Close()

	if status != 200 {
		log.Fatalf("Не могу получить данные, что-то пошло не так, HTTP код: %d", status)
	}

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
	}

	ept.Channels = make(map[string]*base.Channel)
	ept.AppendChannel("EPT1", &base.Channel{
		Name:      "EPT1",
		Icon:      "",
		Programms: nil,
	})

	doc.Find("table.table > tbody > tr > td.table").Not(".time").Each(func(i int, s *goquery.Selection) {

	})
}
