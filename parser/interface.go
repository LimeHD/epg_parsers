package base

import "github.com/PuerkitoBio/goquery"

type IParse interface {
	BaseUrl() string
	Parse(doc *goquery.Document, day int)
	Marshal() string
	BootstrapLocalTime(local string) error
	GetLocalTime() string
}
