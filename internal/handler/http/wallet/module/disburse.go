package module

import (
	mWallet "disburse-app/internal/model/wallet"
	"encoding/json"
	"log"
	"net/http"
)

func (h *handler) Disburse(w http.ResponseWriter, r *http.Request) {
	var (
		req DisburseRequest
		res mWallet.Wallet
		err error
	)

	defer func() {
		// access log
		if err != nil {
			log.Printf("disburse: %+v; failed with err: %+v", req, err)
		} else {
			log.Printf("disburse: %+v; successfully with res: %+v", req, res)
		}
	}()

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err = h.ucWallet.Disburse(req.UserID, req.Amount)
	if err != nil {
		if err.Error() == mWallet.ErrorInsufficient {
			w.WriteHeader(http.StatusBadRequest)
			if err = json.NewEncoder(w).Encode(&DisburseResponse{
				UserID:        req.UserID,
				RequestAmount: req.Amount,
				Balance:       res.Balance,
				Status:        mWallet.ErrorInsufficient,
			}); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(&DisburseResponse{
		UserID:        req.UserID,
		RequestAmount: req.Amount,
		Balance:       res.Balance,
		Status:        "success",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
