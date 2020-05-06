package base

import (
	"encoding/json"
	"fmt"
	"time"
)

type (
	Common struct {
		baseUrl       string
		BaseName      string
		Version       string
		localTimeBase string
		LocalTime     Time
		Channels      map[string]*Channel `json:"channels"`
	}

	Time struct {
		Location *time.Location
		Year     int
		Month    time.Month
		Day      int
	}
)

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

// default implementation marshaling structure array of channels
func (common *Common) Marshal() string {
	jsonString, err := json.Marshal(common.Channels)

	if err != nil {
		panic(err)
	}

	return string(jsonString)
}

// required implementation in every each parser
// set local time by custom location
func (common *Common) BootstrapLocalTime(location string) error {
	loc, err := time.LoadLocation(location)

	if err != nil {
		return err
	}

	currentTime := time.Now().In(loc)

	common.LocalTime = Time{
		Year:     currentTime.Year(),
		Month:    currentTime.Month(),
		Day:      currentTime.Day(),
		Location: loc,
	}

	return nil
}

// default implement handler wrong date and time
func (c *Common) HandleWrongTime(inConsole bool, channel, tv, start, stop string) {
	if inConsole {
		fmt.Println(fmt.Sprintf("Wrong datetime in channel %s: \t %s - %s", channel, start, stop))
	}

	// do stuff
}

// getter & setters

func (common *Common) SetBaseUrl(url string) *Common {
	common.baseUrl = url

	return common
}

func (common *Common) BaseUrl() string {
	return common.baseUrl
}

func (common *Common) SetLocalTime(local string) *Common {
	common.localTimeBase = local

	return common
}

func (common *Common) GetLocalTime() string {
	if common.localTimeBase == "" {
		panic("Not initialize local time")
	}

	return common.localTimeBase
}
