package module

const (
	qInsertUser = `INSERT INTO users VALUES(NULL,?);`

	qSelectBalanceUser = `SELECT balance FROM users WHERE id=?;`

	qUpdateBalanceUser = `UPDATE users SET balance=? WHERE id=?;`
)
