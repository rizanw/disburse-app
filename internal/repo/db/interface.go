package db

import "database/sql"

type DB interface {
	UpdateBalanceUser(tx *sql.Tx, id int64, newBalance float64) error
	GetBalanceUser(tx *sql.Tx, id int64) (float64, error)
}
