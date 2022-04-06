package repository

import (
	"database/sql"

	"github.com/bmf-san/migrate-rubel-to-gobel/gobel/model"
)

type GobelTagPostRepository struct {
	GobelConn *sql.DB
}

func NewGobelTagPostRepository(conn *sql.DB) *GobelTagPostRepository {
	return &GobelTagPostRepository{
		GobelConn: conn,
	}
}

func (ar *GobelTagPostRepository) Write(tagpost *model.TagPost) (int, error) {
	tx, err := ar.GobelConn.Begin()

	rslt, err := tx.Exec(`
		INSERT INTO
			tag_post(id, tag_id, post_id, created_at, updated_at)
		VALUES
			(?, ?, ?, ?, ?)
	`, tagpost.ID, tagpost.TagID, tagpost.PostID, tagpost.CreatedAt, tagpost.UpdatedAt)

	id, err := rslt.LastInsertId()
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return int(id), nil
}
