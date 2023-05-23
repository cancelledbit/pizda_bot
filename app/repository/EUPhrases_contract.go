package repository

type EUPhrase struct {
	ID   int    `json:"id,omitempty"`
	Text string `json:"text,omitempty"`
}

type EUPhrases map[int]*EUPhrase
