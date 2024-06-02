package module

import (
	"database/sql"
	rDB "disburse-app/internal/repo/db"
	ucDisbursment "disburse-app/internal/usecase/wallet"
)

type usecase struct {
	db  *sql.DB
	rDB rDB.DB
}

func New(db *sql.DB, rDB rDB.DB) ucDisbursment.UseCase {
	return &usecase{
		db:  db,
		rDB: rDB,
	}
}
