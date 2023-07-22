package repository

type HagzorPhrase struct {
	ID   int    `json:"id,omitempty"`
	Text string `json:"text,omitempty"`
}

type HagzorPhrases map[int]*HagzorPhrase
