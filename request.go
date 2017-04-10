package bovadaparser

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Request(sport string) ([]byte, error) {
	resp, err := http.Get(
		fmt.Sprintf("http://sportsfeeds.bovada.lv/basic/%s.xml", sport))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, nil
}
