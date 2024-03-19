package structs

type Chapter struct {
	Title   string    `json:"title"`
	Paragraphs   []string  `json:"story"`
	Options []Options `json:"options"`
}

type Options struct {
	Text string `json:"text"`
	Chapter  string `json:"arc"`
}

type Story map[string]Chapter