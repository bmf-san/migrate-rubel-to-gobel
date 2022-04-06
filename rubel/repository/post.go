package repository

import (
	"database/sql"
	"time"

	"github.com/bmf-san/migrate-rubel-to-gobel/rubel/model"
)

type RubelPostRepository struct {
	RubelConn *sql.DB
}

func NewRubelPostRepository(conn *sql.DB) *RubelPostRepository {
	return &RubelPostRepository{
		RubelConn: conn,
	}
}

func (pr *RubelPostRepository) Read() (model.Posts, error) {
	var posts model.Posts

	const query = `
		SELECT
			*
		FROM
			posts
	`
	rows, err := pr.RubelConn.Query(query)

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
			id                int
			adminId           int
			categoryId        int
			title             string
			mdcontent         string
			htmlcontent       string
			views             int
			publicationStatus string
			publishedat       sql.NullTime
			createdat         time.Time
			updatedat         time.Time
			deletedat         sql.NullTime
		)

		if err = rows.Scan(
			&id,
			&adminId,
			&categoryId,
			&title,
			&mdcontent,
			&htmlcontent,
			&views,
			&publicationStatus,
			&publishedat,
			&createdat,
			&updatedat,
			&deletedat,
		); err != nil {
			return nil, err
		}

		post := model.Post{
			ID:                id,
			AdminID:           adminId,
			CategoryID:        categoryId,
			Title:             title,
			MDContent:         mdcontent,
			HTMLContent:       htmlcontent,
			Views:             views,
			PublicationStatus: publicationStatus,
			PublishedAt:       publishedat,
			CreatedAt:         createdat,
			UpdatedAt:         updatedat,
			DeletedAt:         deletedat,
		}
		posts = append(posts, post)
	}

	return posts, nil
}
