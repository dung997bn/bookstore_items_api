package items

import "fmt"

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
	PicID int64  `json:"pic_id"`
	URL   string `json:"url"`
}

type pictureUpdate map[string]interface{}

//MakeUpdateBody func
func MakeUpdateBody(itemOld *Item, itemNew *Item) map[string]interface{} {
	body := make(map[string]interface{})
	if itemOld.Seller != itemNew.Seller && itemNew.Seller != 0 {
		body["seller"] = &itemNew.Seller
	}
	if itemOld.Title != itemNew.Title && itemNew.Title != "" {
		body["title"] = itemNew.Title
	}

	//Description
	if itemOld.Description.PlainText != itemNew.Description.PlainText && itemNew.Description.PlainText != "" &&
		itemOld.Description.HTML != itemNew.Description.HTML && itemNew.Description.HTML != "" {
		descrip := make(map[string]interface{})
		descrip["plain_text"] = itemNew.Description.PlainText
		descrip["html"] = itemNew.Description.HTML
		body["description"] = descrip

	}

	if itemNew.Pictures != nil {
		pics := make([]pictureUpdate, 0)
		for i, p := range itemNew.Pictures {

			if p.PicID != 0 && p.URL != "" {
				pic := pictureUpdate{"id": p.PicID, "url": p.URL}
				pics = append(pics, pic)
				_ = i
				fmt.Println(pic)
			}
		}
		body["pictures"] = pics
	}
	if itemOld.Video != itemNew.Video && itemNew.Video != "" {
		body["video"] = itemNew.Video
	}
	if itemOld.Price != itemNew.Price && itemNew.Price != 0 {
		body["price"] = itemNew.Price
	}
	if itemOld.AvailableQuantity != itemNew.AvailableQuantity && itemNew.AvailableQuantity != 0 {
		body["available_quantity"] = itemNew.AvailableQuantity
	}
	if itemOld.SoldQuantity != itemNew.SoldQuantity && itemNew.SoldQuantity != 0 {
		body["sold_quantity"] = itemNew.SoldQuantity
	}
	if itemOld.Status != itemNew.Status && itemNew.Status != "" {
		body["status"] = itemNew.Status
	}
	return body
}
