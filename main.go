package main

import (
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"fmt"
)

func main() {
	resp, err := http.Get("http://sportsfeeds.bovada.lv/basic/MLB.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf(Unmarshal(string(body)))
}

func Unmarshal(data string) string {
	type Details struct {
		Date string `xml:"PUBLISH_DATE,attr"`
	}

	v := Details{}

	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return ""
	}
	return v.Date
}
