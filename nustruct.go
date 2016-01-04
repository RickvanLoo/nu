package nu

//XML is a struct containing Nu.nl RSS file
type XML struct {
	Channel Channel `xml:"channel"`
}

//Channel contains channel of NUXML feed
type Channel struct {
	Title         string    `xml:"title" json:"title"`
	Description   string    `xml:"description" json:"description"`
	Language      string    `xml:"language" json:"language"`
	Copyright     string    `xml:"copyright" json:"copyright"`
	LastBuildDate string    `xml:"lastBuildDate" json:"lastbuilddate"`
	Item          []Article `xml:"item" json:"item"`
}

//Article constains information of single Nu.nl article
type Article struct {
	Title       string    `xml:"title" json:"title"`
	Link        string    `json:"link"`
	Description string    `xml:"description" json:"description"`
	PubData     string    `xml:"pubDate" json:"pubdate"`
	GUID        int       `xml:"guid" json:"guid"`
	Categories  []string  `xml:"category" json:"categories"`
	Creator     string    `xml:"creator" json:"creator"`
	Rights      string    `xml:"rights" json:"rights"`
	Image       Enclosure `xml:"enclosure" json:"image"`
	Related     []Atom    `xml:"http://www.w3.org/2005/Atom link" json:"related"`
}

//Enclosure contains image of article
type Enclosure struct {
	Link string `xml:"url,attr" json:"url"`
}

//Atom contains reference articles
type Atom struct {
	Link  string `xml:"href,attr"`
	Title string `xml:"title,attr"`
}
