package model

import (
	"database/sql"
	"time"

	gmodel "github.com/bmf-san/migrate-rubel-to-gobel/gobel/model"
)

// A Categories represents the plural of category.
type Categories []Category

// A Category represents the singular of category.
type Category struct {
	ID        int          `json:"id"`
	Name      string       `json:"name"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

func (c *Category) Convert() *gmodel.Category {
	return &gmodel.Category{
		ID:        c.ID,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}
