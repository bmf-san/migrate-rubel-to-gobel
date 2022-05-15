package model

import (
	"database/sql"
	"time"

	gmodel "github.com/bmf-san/migrate-rubel-to-gobel/gobel/model"
)

// A Posts represents the plural of post.
type Posts []Post

// A Post represents the singular of post.
type Post struct {
	ID                int          `json:"id"`
	AdminID           int          `json:"admin_id"`
	CategoryID        int          `json:"category_id"`
	Title             string       `json:"title"`
	MDContent         string       `json:"md_content"`
	HTMLContent       string       `json:"html_content"`
	Views             int          `json:"views"`
	PublicationStatus string       `json:"publication_status"`
	PublishedAt       sql.NullTime `json:"published_at"`
	CreatedAt         time.Time    `json:"created_at"`
	UpdatedAt         time.Time    `json:"updated_at"`
	DeletedAt         sql.NullTime `json:"deleted_at"`
}

func (p *Post) Convert() *gmodel.Post {
	var status string
	switch p.PublicationStatus {
	case "public":
		status = "publish"
	case "private":
		status = "draft"
	case "draft":
		status = "draft"
	}

	return &gmodel.Post{
		ID:         p.ID,
		AdminID:    p.AdminID,
		CategoryID: p.CategoryID,
		Title:      p.Title,
		MDBody:     p.MDContent,
		HTMLBody:   p.HTMLContent,
		Status:     status,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}
}
