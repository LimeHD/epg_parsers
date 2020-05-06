package base

import (
	"epg_parsers/utils"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

type Parser struct{}

func (p *Parser) RunComposeWith(parser IParse, day int) error {
	fmt.Println(fmt.Sprintf("Get HTML document from %s", parser.BaseUrl()))

	status, reader := utils.GetHtmlDocumentReader(parser.BaseUrl())
	defer reader.Close()

	if status != 200 {
		return errors.New(fmt.Sprintf("Не могу получить данные, что-то пошло не так, HTTP код: %d", status))
	}

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return err
	}

	err = parser.BootstrapLocalTime(parser.GetLocalTime())
	if err != nil {
		return err
	}

	parser.Parse(doc, day)

	return nil
}
