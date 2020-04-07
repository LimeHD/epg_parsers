package base

type (
	Epg struct {
		Timestart int64
		Timestop  int64
		Title     string
		Desc      string
		Rating    int
		Genre     []string
		Category  []string
	}

	Response struct {
		Days map[int64]Epg
	}
)

type IParse interface {
	Parse() Response
}

// todo methods which may come in handy

func (response *Response) IssetDay() bool {
	return true
}

func (response *Response) CreateDay() {
	// todo
}

func (response *Response) AppendProgramm() {
	// todo
}
