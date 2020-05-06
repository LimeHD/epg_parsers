package base

import "github.com/PuerkitoBio/goquery"

// main interface, every parser must implement it
type IParse interface {
	BaseUrl() string
	Parse(doc *goquery.Document, day int)
	Marshal() string
	BootstrapLocalTime(local string) error
	GetLocalTime() string
}
