package models

type Gallery struct {
	ID          int    `json:"id"`
	ThumbID     int    `json:"thumb_id"`
	OwnerID     int    `json:"owner_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Created     int    `json:"created"`
	Updated     int    `json:"updated"`
	Size        int    `json:"size"`
}
