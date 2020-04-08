package base

import "encoding/json"

type (
	Programm struct {
		Timestart string `json:"timestart"`
		Timestop  string `json:"timestop"`
		Title     string `json:"title"`
	}

	Channel struct {
		Name      string     `json:"name"`
		Icon      string     `json:"icon"`
		Programms []Programm `json:"programms"`
	}

	Common struct {
		Channels map[string]*Channel `json:"channels"`
	}
)

type IParse interface {
	BaseUrl()
	Parse()
	Marshal() string
}

func (common *Common) AppendProgramm(name string, programm Programm) {
	common.Channels[name].Programms = append(common.Channels[name].Programms, programm)
}

func (common *Common) AppendChannel(name string, channel *Channel) {
	common.Channels[name] = channel
}

func (common *Common) ChannelExist(name string) bool {
	_, exist := common.Channels[name]

	return exist
}

func (common *Common) Marshal() string {
	jsonString, err := json.Marshal(common.Channels)

	if err != nil {
		panic(err)
	}

	return string(jsonString)
}
