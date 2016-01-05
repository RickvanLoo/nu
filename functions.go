package nu

import (
	"encoding/xml"
	"io"
	"net/http"
)

//ParseToFeed returns JSON in []Byte from NU.nl RSS XML feed as string
func parseToFeed(url string) (Feed XML, err error) {

	file, err := http.Get(url)
	if err != nil {
		return
	}
	defer file.Body.Close()

	PreFeed, err := nuToPreFeed(file.Body)
	if err != nil {
		return
	}

	file, err = http.Get(url)
	if err != nil {
		return
	}
	defer file.Body.Close()

	Feed, err = preFeedToJSON(PreFeed, file.Body)
	if err != nil {
		return
	}

	return
}

func nuToPreFeed(f io.Reader) (NuFeed XML, err error) {

	NuFeed = XML{}
	d := xml.NewDecoder(f)
	err = d.Decode(&NuFeed)
	if err != nil {
		return XML{}, err
	}

	return NuFeed, err
}

func preFeedToJSON(PreFeed XML, f io.Reader) (Feed XML, err error) {

	NuFeed2 := xML2{}
	d := xml.NewDecoder(f)
	err = d.Decode(&NuFeed2)
	if err != nil {
		return XML{}, err
	}

	for i := 0; i < len(PreFeed.Channel.Item); i++ {
		PreFeed.Channel.Item[i].Link = NuFeed2.Channel.Item[i].Link[0]
	}

	return PreFeed, err
}
