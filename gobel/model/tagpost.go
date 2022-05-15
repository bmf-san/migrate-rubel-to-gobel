package model

import (
	"time"
)

// A TagPosts represents the plural of tag.
type TagPosts []TagPost

// A TagPost represetns the singular of tag.
type TagPost struct {
	ID        int       `json:"id"`
	TagID     int       `json:"tag_id"`
	PostID    int       `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
