package usecase

import (
	"context"
	. "go-server/domain"
	"go-server/internal/tokenutil"
	"time"
)

type refreshTokenUsecase struct {
	userRepository UserRepository
	contextTimeout time.Duration
}

func NewRefreshTokenUsecase(userRepository UserRepository, timeout time.Duration) RefreshTokenUsecase {
	return &refreshTokenUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}
func (u *refreshTokenUsecase) GetUserByID(c context.Context, email string) (User, error) {
	_, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.userRepository.GetByID(email)
}

func (u *refreshTokenUsecase) CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (u *refreshTokenUsecase) CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}

func (u *refreshTokenUsecase) ExtractIDFromToken(requestToken string, secret string) (string, error) {
	return tokenutil.ExtractIDFromToken(requestToken, secret)
}
