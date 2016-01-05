package nu

//XML2 contains link of NU.nl Items (Hack for encoding/xml)
type xML2 struct {
	Channel channel2 `xml:"channel"`
}

//Channel2 contains channel of NUXML2 feed
type channel2 struct {
	Item []article2 `xml:"item"`
}

//Article2 constains link of single Nu.nl article
type article2 struct {
	Link []string `xml:"link" json:"link"`
}
