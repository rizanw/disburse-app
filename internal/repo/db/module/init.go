package module

import (
	"database/sql"
	"disburse-app/internal/config"
	rDB "disburse-app/internal/repo/db"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type sqlite struct {
	db *sql.DB
}

func New(conf config.SqliteConfig) (rDB.DB, error) {
	db, err := sql.Open("sqlite3", conf.Path)
	if err != nil {
		return nil, err
	}

	if err = initDB(db); err != nil {
		return nil, err
	}

	return &sqlite{
		db: db,
	}, nil
}

func initDB(db *sql.DB) error {
	var (
		count int
		err   error
	)

	if err = db.QueryRow(`SELECT count(*) FROM sqlite_master WHERE type='table' AND name=?`, "users").Scan(&count); err != nil {
		return err
	}

	if count <= 0 {
		// create table
		if _, err = db.Exec(createUsersTable); err != nil {
			return err
		}
		// populate initiated data
		db.Exec(qInsertUser, 10000000)
		db.Exec(qInsertUser, 20000000)
		db.Exec(qInsertUser, 25000000)

		log.Println("initiate table created successfully")
	}

	return nil
}
