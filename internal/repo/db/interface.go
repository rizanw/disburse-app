package db

import "database/sql"

//go:generate mockgen -package=mock -source=./interface.go -destination=./_mock/db_mock.go
type DB interface {
	UpdateBalanceUser(tx *sql.Tx, id int64, newBalance float64) error
	GetBalanceUser(tx *sql.Tx, id int64) (float64, error)
}
