package repository

import (
	"database/sql"

	"github.com/bmf-san/migrate-rubel-to-gobel/gobel/model"
)

type GobelTagRepository struct {
	GobelConn *sql.DB
}

func NewGobelTagRepository(conn *sql.DB) *GobelTagRepository {
	return &GobelTagRepository{
		GobelConn: conn,
	}
}

func (ar *GobelTagRepository) Write(tag *model.Tag) (int, error) {
	tx, err := ar.GobelConn.Begin()

	rslt, err := tx.Exec(`
		INSERT INTO
			tags(id, name, created_at, updated_at)
		VALUES
			(?, ?, ?, ?)
	`, tag.ID, tag.Name, tag.CreatedAt, tag.UpdatedAt)

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
