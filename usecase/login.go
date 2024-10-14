package usecase

import (
	"context"
	. "go-server/domain"
	"go-server/internal/tokenutil"
	"time"
)

type loginUsecase struct {
	userRepository UserRepository
	contextTimeout time.Duration
}

func NewLoginUsecase(userRepository UserRepository, timeout time.Duration) LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (u *loginUsecase) GetUserByEmail(c context.Context, email string) (User, error) {
	_, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.userRepository.GetByEmail(email)
}

func (u *loginUsecase) CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (u *loginUsecase) CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
