package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://sportsfeeds.bovada.lv/basic/NHL.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf(Unmarshal(body))
}

func Unmarshal(data []byte) string {
	type Event struct {
		ID string `xml:"ID,attr"`
	}

	type Date struct {
		Event []Event
	}

	type EventType struct {
		ID   string `xml:"ID,attr"`
		Date Date
	}

	type Schedule struct {
		Date      string `xml:"PUBLISH_DATE,attr"`
		EventType EventType
	}

	v := Schedule{}

	err := xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return ""
	}

	fmt.Printf("Games: %v\n", v.EventType)

	return v.Date
}
