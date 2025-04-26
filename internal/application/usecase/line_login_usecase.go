package usecase

import "backend-service/internal/domain"

func (uc *Usecase) FindLineUserByUID(lineUID string) (*domain.Line, error) {
	return uc.lineRepo.FindLineUserByUID(lineUID)
}

func (uc *Usecase) CreateLineUser(lineUser *domain.Line) error {
	return uc.lineRepo.CreateLineUser(lineUser)
}

func (uc *Usecase) VerifyIDToken(idToken string) (*domain.Line, error) {
	return uc.lineRepo.VerifyIDToken(idToken)
}

func (uc *Usecase) UpdateLineUser(lineUser domain.Line) error {
	return uc.lineRepo.UpdateLineUser(lineUser)
}

func (uc *Usecase) SaveProfile(idToken string) (*domain.Line, error) {
	profile, err := uc.VerifyIDToken(idToken)
	if err != nil {
		return nil, err
	}

	existingLineUser, err := uc.FindLineUserByUID(profile.UID)
	if err != nil {
		return nil, err
	}

	if existingLineUser != nil {
		existingLineUser.DisplayName = profile.DisplayName
		existingLineUser.PictureURL = profile.PictureURL

		err = uc.UpdateLineUser(*existingLineUser)
		if err != nil {
			return nil, err
		}

		return existingLineUser, nil
	}

	err = uc.CreateLineUser(profile)
	if err != nil {
		return nil, err
	}

	return profile, nil
}