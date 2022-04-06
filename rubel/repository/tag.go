package repository

import (
	"database/sql"
	"time"

	"github.com/bmf-san/migrate-rubel-to-gobel/rubel/model"
)

type RubelTagRepository struct {
	RubelConn *sql.DB
}

func NewRubelTagRepository(conn *sql.DB) *RubelTagRepository {
	return &RubelTagRepository{
		RubelConn: conn,
	}
}

func (tr *RubelTagRepository) Read() (model.Tags, error) {
	var tags model.Tags

	const query = `
		SELECT
			*
		FROM
			tags
	`
	rows, err := tr.RubelConn.Query(query)

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

		tag := model.Tag{
			ID:        id,
			Name:      name,
			CreatedAt: createdat,
			UpdatedAt: updatedat,
			DeletedAt: deletedat,
		}
		tags = append(tags, tag)
	}

	return tags, nil
}
