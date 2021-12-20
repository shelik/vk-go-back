package models

type Gallery struct {
	ID          int    `json:"id"`
	ThumbID     int    `json:"thumb_id"`
	OwnerID     int    `json:"owner_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Created     string `json:"created"`
	Updated     string `json:"updated"`
	Size        string `json:"size"`
}
