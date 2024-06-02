package module

import (
	hWallet "disburse-app/internal/handler/http/wallet"
	ucWallet "disburse-app/internal/usecase/wallet"
)

type handler struct {
	ucWallet ucWallet.UseCase
}

func New(ucWallet ucWallet.UseCase) hWallet.Handler {
	return &handler{
		ucWallet: ucWallet,
	}
}
