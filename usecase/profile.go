package usecase

import (
	"context"
	. "go-server/domain"
	"time"
)

type profileUsecase struct {
	userRepository UserRepository
	contextTimeout time.Duration
}

func NewProfileUsecase(userRepository UserRepository, timeout time.Duration) ProfileUsecase {
	return &profileUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (pu *profileUsecase) GetProfileByID(c context.Context, userID string) (*Profile, error) {
	_, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	user, err := pu.userRepository.GetByID(userID)
	if err != nil {
		return nil, err
	}

	return &Profile{Name: user.Name, Email: user.Email}, nil
}
