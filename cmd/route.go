package main

import (
	hWallet "disburse-app/internal/handler/http/wallet/module"
	ucWallet "disburse-app/internal/usecase/wallet"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func newRoute(ucWallet ucWallet.UseCase) *mux.Router {
	router := mux.NewRouter()

	// auth routes
	handlerWallet := hWallet.New(ucWallet)
	router.HandleFunc("/disburse", handlerWallet.Disburse).Methods(http.MethodPost)

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Println("server OK!")
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)

	return router
}
