package items

//Item type
type Item struct {
	ID                string      `json:"id"`
	Seller            int64       `json:"seller"`
	Title             string      `json:"title"`
	Description       Description `json:"description"`
	Pictures          []Picture   `json:"pictures"`
	Video             string      `json:"video"`
	Price             float32     `json:"price"`
	AvailableQuantity int         `json:"available_quantity"`
	SoldQuantity      int         `json:"sold_quantity"`
	Status            string      `json:"status"`
}

//Description type
type Description struct {
	PlainText string `json:"plain_text"`
	HTML      string `json:"html"`
}

//Picture type
type Picture struct {
	ID  int64  `json:"id"`
	URL string `json:"url"`
}
