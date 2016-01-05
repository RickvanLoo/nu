package nu

import (
	"encoding/json"
	"fmt"
)

//Struct returns the NU RSS feed as struct XML
func Struct(url string) (feed XML, err error) {
	feed, err = parseToFeed(url)
	if err != nil {
		return XML{}, err
	}
	return
}

//JSON returns the NU RSS feed as JSON in []byte
func JSON(url string) (JSON []byte, err error) {
	feed, err := parseToFeed(url)
	if err != nil {
		return nil, err
	}

	fmt.Println(feed)

	JSON, err = json.Marshal(feed)
	if err != nil {
		return nil, err
	}

	return
}
