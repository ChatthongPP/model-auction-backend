package domain

type WalletLogInterface interface {
	GetWalletLogs(filter *FilterRequest, offset int) ([]*WalletLog, int, error)
}
