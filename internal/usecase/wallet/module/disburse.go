package module

import (
	mWallet "disburse-app/internal/model/wallet"
	"errors"
)

func (u *usecase) Disburse(userID int64, amount float64) (mWallet.Wallet, error) {
	var (
		wallet mWallet.Wallet
		err    error
	)
	wallet.UserID = userID

	// create db transaction
	tx, err := u.db.Begin()
	defer tx.Rollback()
	if err != nil {
		return wallet, err
	}

	// get user balance
	balance, err := u.rDB.GetBalanceUser(tx, userID)
	if err != nil {
		return wallet, err
	}
	wallet.Balance = balance

	// validate balance funds
	if balance < amount {
		return wallet, errors.New("insufficient funds")
	}

	// update new balance
	newBalance := balance - amount
	if err = u.rDB.UpdateBalanceUser(tx, userID, newBalance); err != nil {
		return wallet, err
	}

	// commit db tx
	if err = tx.Commit(); err != nil {
		return wallet, err
	}

	return wallet, nil
}
