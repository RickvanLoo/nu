package nu

//XML2 contains link of NU.nl Items (Hack for encoding/xml)
type XML2 struct {
	Channel Channel2 `xml:"channel"`
}

//Channel2 contains channel of NUXML2 feed
type Channel2 struct {
	Item []Article2 `xml:"item"`
}

//Article2 constains link of single Nu.nl article
type Article2 struct {
	Link []string `xml:"link" json:"link"`
}
