package repository

import (
	"database/sql"

	"github.com/bmf-san/migrate-rubel-to-gobel/gobel/model"
)

type GobelAdminRepository struct {
	GobelConn *sql.DB
}

func NewGobelAdminRepository(conn *sql.DB) *GobelAdminRepository {
	return &GobelAdminRepository{
		GobelConn: conn,
	}
}

func (ar *GobelAdminRepository) Write(admin *model.Admin) (int, error) {
	tx, err := ar.GobelConn.Begin()

	rslt, err := tx.Exec(`
		INSERT INTO
			admins(id, name, email, password, created_at, updated_at)
		VALUES
			(?, ?, ?, ?, ?, ?)
	`, admin.ID, admin.Name, admin.Email, admin.Password, admin.CreatedAt, admin.UpdatedAt)

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
