package base

type (
	Programm struct {
		Timestart   string `json:"timestart"`
		Timestop    string `json:"timestop"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	Channel struct {
		Name      string     `json:"name"`
		Icon      string     `json:"icon"`
		Programms []Programm `json:"programms"`
	}

	Day struct {
		Name   string
		Common *Common
	}

	Epg struct {
		Days map[string]*Day
	}
)

func (epg *Epg) AppendDay(name string, day *Day) {
	epg.Days[name] = day
}

func (epg *Epg) DayExist(day string) bool {
	_, exist := epg.Days[day]

	return exist
}
