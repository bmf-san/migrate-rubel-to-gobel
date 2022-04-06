package model

import (
	"database/sql"
	"time"

	gmodel "github.com/bmf-san/migrate-rubel-to-gobel/gobel/model"
)

// A TagPosts represents the plural of tag.
type TagPosts []TagPost

// A TagPost represetns the singular of tag.
type TagPost struct {
	ID        int          `json:"id"`
	TagID     int          `json:"tag_id"`
	PostID    int          `json:"post_id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

func (tp *TagPost) Convert() *gmodel.TagPost {
	return &gmodel.TagPost{
		ID:        tp.ID,
		TagID:     tp.TagID,
		PostID:    tp.PostID,
		CreatedAt: tp.CreatedAt,
		UpdatedAt: tp.UpdatedAt,
	}
}
