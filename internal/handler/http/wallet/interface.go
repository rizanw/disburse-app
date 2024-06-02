package wallet

import "net/http"

type Handler interface {
	Disburse(w http.ResponseWriter, r *http.Request)
}
