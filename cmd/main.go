package main

import (
	"disburse-app/internal/config"
	dbSqlite "disburse-app/internal/repo/db/module"
	ucWallet "disburse-app/internal/usecase/wallet/module"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	appName = "disburse-app"
)

func main() {
	var (
		address string
	)

	// init config
	conf, err := config.New(appName)
	if err != nil {
		log.Printf("failed to load config: %v\n", err)
		return
	}

	// init repo
	rdb, dbconn, err := dbSqlite.New(conf.Database)
	if err != nil {
		log.Printf("failed to load db: %v\n", err)
		return
	}

	// init usecase logic
	_ = ucWallet.New(dbconn, rdb)

	address = fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)
	srv := http.Server{
		Addr:         address,
		ReadTimeout:  time.Duration(conf.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(conf.Server.WriteTimeout) * time.Second,
		Handler:      nil,
	}

	log.Println("!app starting on ", address)
	err = srv.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
