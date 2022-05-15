package model

import (
	"database/sql"
	"time"

	gmodel "github.com/bmf-san/migrate-rubel-to-gobel/gobel/model"
)

// A Admins represents the plural of admin.
type Admins []Admin

// A Admin represents the singular of admin.
type Admin struct {
	ID           int            `json:"id"`
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	Password     string         `json:"password"`
	RemeberToken sql.NullString `json:"rember_token"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    sql.NullTime   `json:"delted_at"`
}

func (a *Admin) Convert() *gmodel.Admin {
	return &gmodel.Admin{
		ID:        a.ID,
		Name:      a.Name,
		Email:     a.Email,
		Password:  a.Password,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}
