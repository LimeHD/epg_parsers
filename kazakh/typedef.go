package main

import "time"

type (
	Epg struct {
		Channel string             `json:"name"`
		Icon    string             `json:"icon"`
		Days    map[time.Time]*Day `json:"days"`
	}
	Day struct {
		Programs []*Program `json:"programs"`
	}
	Program struct {
		Timestart   string `json:"timestart"`
		Timestop    string `json:"timestop"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	Row struct {
		SheetOne Values
		SheetTwo Values
	}
	Values struct {
		Time     string
		Title    string
		Duration string
	}
)
