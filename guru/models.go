package guru

type Fact struct {
	Id      string `json:"id"`
	Content string `json:"content"`
	Title   string `json:"preferredPhrase"`
	Type    string `json:"cardType"`
}

//{"preferredPhrase":"gitlander","content":"table","verificationInterval":30,"shareStatus":"TEAM","cardType":"CARD"}

type Card struct {
	Id                   string `json:"id, omitempty"`
	PreferredPhrase      string `json:"preferredPhrase"`
	Content              string `json:"content"`
	VerificationInterval string `json:"verificationInterval"`
	ShareStatus          string `json:"shareStatus"`
	CardType             string `json:"cardType"`
}

type Board struct {
	Id          string  `json:"id, omitempty"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Items       []*Card `json:"items"`
}
