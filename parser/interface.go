package base

import "github.com/PuerkitoBio/goquery"

// main interface, every parser must implement it
type IParse interface {
	Parse(doc *goquery.Document, day int)
	BaseUrl() string
	Marshal() string
	GetLocalTime() string
	BootstrapLocalTime(local string) error
}
