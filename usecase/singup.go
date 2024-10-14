package usecase

import (
	"context"
	. "go-server/domain"
	"go-server/internal/tokenutil"
	"time"
)

type signupUsecase struct {
	userRepository UserRepository
	contextTimeout time.Duration
}

func NewSignupUsecase(userRepository UserRepository, timeout time.Duration) SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (u *signupUsecase) Create(c context.Context, user *User) error {
	_, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.userRepository.Create(user)
}

func (u *signupUsecase) GetUserByEmail(c context.Context, email string) (User, error) {
	_, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.userRepository.GetByEmail(email)
}

func (u *signupUsecase) CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (u *signupUsecase) CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
