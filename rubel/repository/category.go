package repository

import (
	"database/sql"
	"time"

	"github.com/bmf-san/migrate-rubel-to-gobel/rubel/model"
)

type RubelCategoryRepository struct {
	RubelConn *sql.DB
}

func NewRubelCategoryRepository(conn *sql.DB) *RubelCategoryRepository {
	return &RubelCategoryRepository{
		RubelConn: conn,
	}
}

func (cr *RubelCategoryRepository) Read() (model.Categories, error) {
	var categories model.Categories

	const query = `
		SELECT
			*
		FROM
			categories
	`
	rows, err := cr.RubelConn.Query(query)

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
			id        int
			name      string
			createdat time.Time
			updatedat time.Time
			deletedat sql.NullTime
		)

		if err = rows.Scan(
			&id,
			&name,
			&createdat,
			&updatedat,
			&deletedat,
		); err != nil {
			return nil, err
		}

		category := model.Category{
			ID:        id,
			Name:      name,
			CreatedAt: createdat,
			UpdatedAt: updatedat,
			DeletedAt: deletedat,
		}
		categories = append(categories, category)
	}

	return categories, nil
}
