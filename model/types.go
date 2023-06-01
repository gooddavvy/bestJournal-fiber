package model

type JournalPage struct {
	Title     string `json:"title"`
	Date      string `json:"date"`
	ShortDesc string `json:"shortDesc"`
	Link      string `json:"link"`
	Body      string `json:"body"`
}
