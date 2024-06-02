package module

import (
	"database/sql"
)

func (s *sqlite) UpdateBalanceUser(tx *sql.Tx, id int64, newBalance float64) error {
	if _, err := tx.Exec(qUpdateBalanceUser, newBalance, id); err != nil {
		return err
	}

	return nil
}
