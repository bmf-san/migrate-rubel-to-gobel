package repository

import (
	"database/sql"

	"github.com/bmf-san/migrate-rubel-to-gobel/gobel/model"
)

type GobelCategoryRepository struct {
	GobelConn *sql.DB
}

func NewGobelCategoryRepository(conn *sql.DB) *GobelCategoryRepository {
	return &GobelCategoryRepository{
		GobelConn: conn,
	}
}

func (ar *GobelCategoryRepository) Write(category *model.Category) (int, error) {
	tx, err := ar.GobelConn.Begin()

	rslt, err := tx.Exec(`
		INSERT INTO
			categories(id, name, created_at, updated_at)
		VALUES
			(?, ?, ?, ?)
	`, category.ID, category.Name, category.CreatedAt, category.UpdatedAt)

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
