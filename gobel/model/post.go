package model

import "time"

// A Posts represents the plural of post.
type Posts []Post

// A Post represents the singular of post.
type Post struct {
	ID         int       `json:"id"`
	AdminID    int       `json:"admin_id"`
	CategoryID int       `json:"category_id"`
	Title      string    `json:"title"`
	MDBody     string    `json:"md_content"`
	HTMLBody   string    `json:"html_content"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}
