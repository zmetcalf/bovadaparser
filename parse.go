package bovadaparser

import (
	"encoding/xml"
)

type Odds struct {
	Fraction   string `xml:"FRACTION,attr"`
	Line       string `xml:"Line,attr"`
	Multiplier string `xml:"MULTIPLIER,attr"`
	Risk       string `xml:"RISK,attr"`
	Win        string `xml:"WIN,attr"`
}

type Choice struct {
	Value string `xml:"VALUE,attr"`
	Odds  []Odds
}

type Line struct {
	Type   string `xml:"TYPE,attr"`
	Choice []Choice
}

type Competitor struct {
	ID   string `xml:"ID,attr"`
	Name string `xml:"NAME,attr"`
	Line []Line
}

type Time struct {
	Start string `xml:"TTEXT,attr"`
}

type Event struct {
	ID         string `xml:"ID,attr"`
	Competitor []Competitor
	Time       Time
}

type Date struct {
	GameDate string `xml:"DTEXT,attr"`
	Event    []Event
}

type EventType struct {
	ID   string `xml:"ID,attr"`
	Date Date
}

type Schedule struct {
	PublishDate string `xml:"PUBLISH_DATE,attr"`
	PublishTime string `xml:"PUBLISH_TIME,attr"`
	EventType   EventType
}

func Unmarshal(data []byte) (Schedule, error) {
	v := Schedule{}

	err := xml.Unmarshal(data, &v)
	if err != nil {
		return Schedule{}, err
	}

	return v, nil
}
