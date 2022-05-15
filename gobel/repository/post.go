package repository

import (
	"database/sql"

	"github.com/bmf-san/migrate-rubel-to-gobel/gobel/model"
)

type GobelPostRepository struct {
	GobelConn *sql.DB
}

func NewGobelPostRepository(conn *sql.DB) *GobelPostRepository {
	return &GobelPostRepository{
		GobelConn: conn,
	}
}

func (gr *GobelPostRepository) Write(post *model.Post) (int, error) {
	tx, err := gr.GobelConn.Begin()

	rslt, err := tx.Exec(`
		INSERT INTO
			posts(id, admin_id, category_id, title, md_body, html_body, status, created_at, updated_at)
		VALUES
			(?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, post.ID, post.AdminID, post.CategoryID, post.Title, post.MDBody, post.HTMLBody, post.Status, post.CreatedAt, post.UpdatedAt)

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
