package model

import (
	"database/sql"
	"time"

	gmodel "github.com/bmf-san/migrate-rubel-to-gobel/gobel/model"
)

// A Tags represents the plural of tag.
type Tags []Tag

// A Tag represetns the singular of tag.
type Tag struct {
	ID        int          `json:"id"`
	Name      string       `json:"name"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

func (t *Tag) Convert() *gmodel.Tag {
	return &gmodel.Tag{
		ID:        t.ID,
		Name:      t.Name,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}
