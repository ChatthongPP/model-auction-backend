package usecase

import "backend-service/internal/domain"

func (uc *Usecase) GetWalletLogs(filter *domain.FilterRequest) ([]*domain.WalletLog, int, int, error) {
	offset := (filter.CurrrentPage - 1) * filter.Limit

	WalletLogs, totalCount, err := uc.walletLogRepo.GetWalletLogs(filter, offset)
	if err != nil {
		return nil, 0, 0, err
	}

	totalPages := (totalCount + filter.Limit - 1) / filter.Limit

	return WalletLogs, totalCount, totalPages, nil
}

func (uc *Usecase) CreateWalletLog(walletLog *domain.WalletLog) error {
	err := uc.walletLogRepo.CreateWalletLog(walletLog)
	if err != nil {
		return err
	}
	return nil
}
