package inventory

import "time"

type ShirtInventory struct {
	ShirtID     string    `json:"id"`
	UserId      string    `json:"user_id"`
	ItemName    string    `json:"item_name"`
	Quantity    int       `json:"quantity"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}

// NewShirtInventory contains information needed to create a ShirtInventory.
type NewShirtInventory struct {
	ItemName string `json:"item_name"`
	Quantity int    `json:"quantity"`
}
