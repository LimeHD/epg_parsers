package Star

import (
	base "epg_parsers/parser"
	"github.com/PuerkitoBio/goquery"
)

type Star struct {
	base.Common
}

func (s *Star) GetLocalTime() string {
	return "Europe/Athens"
}

func (s *Star) Parse(doc *goquery.Document, day int) {

}
