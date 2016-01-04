package nu

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
)

// func main() {
//
// 	JSON, err := Parse("http://www.nu.nl/rss/Algemeen")
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
//
// 	log.Println(string(JSON))
//
// }

//Parse returns JSON in []Byte from NU.nl RSS XML feed as string
func Parse(url string) (JSON []byte, err error) {

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

	JSON, err = preFeedToJSON(PreFeed, file.Body)
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
		return
	}

	return
}

func preFeedToJSON(PreFeed XML, f io.Reader) (JSON []byte, err error) {

	NuFeed2 := XML2{}
	d := xml.NewDecoder(f)
	err = d.Decode(&NuFeed2)
	if err != nil {
		return
	}

	for i := 0; i < len(PreFeed.Channel.Item); i++ {
		PreFeed.Channel.Item[i].Link = NuFeed2.Channel.Item[i].Link[0]
	}

	JSON, err = json.Marshal(PreFeed)
	if err != nil {
		return
	}

	return
}
