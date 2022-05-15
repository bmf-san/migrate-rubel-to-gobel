package repository

import (
	"database/sql"
	"time"

	"github.com/bmf-san/migrate-rubel-to-gobel/rubel/model"
)

type RubelTagPostRepository struct {
	RubelConn *sql.DB
}

func NewRubelTagPostRepository(conn *sql.DB) *RubelTagPostRepository {
	return &RubelTagPostRepository{
		RubelConn: conn,
	}
}

func (tr *RubelTagPostRepository) Read() (model.TagPosts, error) {
	var tagposts model.TagPosts

	const query = `
		SELECT
			*
		FROM
			tag_post
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
			tagId     int
			postId    int
			createdat time.Time
			updatedat time.Time
			deletedat sql.NullTime
		)

		if err = rows.Scan(
			&id,
			&tagId,
			&postId,
			&createdat,
			&updatedat,
			&deletedat,
		); err != nil {
			return nil, err
		}

		tagpost := model.TagPost{
			ID:        id,
			TagID:     tagId,
			PostID:    postId,
			CreatedAt: createdat,
			UpdatedAt: updatedat,
			DeletedAt: deletedat,
		}
		tagposts = append(tagposts, tagpost)
	}

	return tagposts, nil
}
