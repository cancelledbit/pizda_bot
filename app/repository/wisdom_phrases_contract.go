package repository

type WisdomPhrase struct {
	ID       int    `json:"id,omitempty"`
	Text     string `json:"text,omitempty"`
	AuthorId string `json:"author,omitempty"`
}

type WisdomPhrases map[int]*WisdomPhrase
