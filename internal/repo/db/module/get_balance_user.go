package module

import "database/sql"

func (s *sqlite) GetBalanceUser(tx *sql.Tx, id int64) (float64, error) {
	var balance float64

	if err := tx.QueryRow(qSelectBalanceUser, id).Scan(&balance); err != nil {
		return 0, err
	}

	return balance, nil
}
