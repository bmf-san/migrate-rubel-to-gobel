package repository

import (
	"database/sql"
	"time"

	"github.com/bmf-san/migrate-rubel-to-gobel/rubel/model"
)

type RubelAdminRepository struct {
	RubelConn *sql.DB
}

func NewRubelAdminRepository(conn *sql.DB) *RubelAdminRepository {
	return &RubelAdminRepository{
		RubelConn: conn,
	}
}

func (ar *RubelAdminRepository) Read() (model.Admins, error) {
	var admins model.Admins

	const query = `
		SELECT
			*
		FROM
			admins
	`
	rows, err := ar.RubelConn.Query(query)

	defer func() {
		if rerr := rows.Close(); rerr != nil {
			err = rerr
		}
	}()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id          int
			name        string
			email       string
			password    string
			rembertoken sql.NullString
			createdat   time.Time
			updatedat   time.Time
			deletedat   sql.NullTime
		)

		if err = rows.Scan(
			&id,
			&name,
			&email,
			&password,
			&rembertoken,
			&createdat,
			&updatedat,
			&deletedat,
		); err != nil {
			return nil, err
		}

		admin := model.Admin{
			ID:           id,
			Name:         name,
			Email:        email,
			Password:     password,
			RemeberToken: rembertoken,
			CreatedAt:    createdat,
			UpdatedAt:    updatedat,
			DeletedAt:    deletedat,
		}
		admins = append(admins, admin)
	}

	return admins, nil
}
