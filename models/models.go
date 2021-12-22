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

type Photo struct {
	Id       int    `json:"id"`
	Album_id int    `json:"album_id"`
	Owner_id int    `json:"owner_id"`
	User_id  int    `json:"user_id"`
	Text     string `json:"text"`
	Date     int    `json:"date"`
}
