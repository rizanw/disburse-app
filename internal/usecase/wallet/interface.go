package wallet

import mWallet "disburse-app/internal/model/wallet"

type UseCase interface {
	Disburse(userID int64, amount float64) (mWallet.Wallet, error)
}
