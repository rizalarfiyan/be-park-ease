package service

import (
	"be-park-ease/config"
	"be-park-ease/exception"
	"be-park-ease/internal/repository"
	"be-park-ease/internal/request"
	"be-park-ease/internal/response"
	"be-park-ease/internal/sql"
	"be-park-ease/utils"
	"context"
	"time"
)

type AuthService interface {
	Login(ctx context.Context, req request.AuthLoginRequest) response.AuthLoginResponse
}

type authService struct {
	repo      repository.AuthRepository
	exception exception.Exception
	conf      config.Config
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{
		repo:      repo,
		exception: exception.NewException("auth-service"),
		conf:      *config.Get(),
	}
}

func (s authService) Login(ctx context.Context, req request.AuthLoginRequest) response.AuthLoginResponse {
	user, err := s.repo.GetUserByUsername(ctx, req.Username)
	s.exception.PanicIfErrorWithoutNoSqlResult(err, false)
	s.exception.IsBadRequest(user, "Invalid username or password", false)

	isValidPassword := utils.ComparePassword(user.Password, req.Password)

	s.exception.IsBadRequest(isValidPassword, "Invalid username or password", false)

	isBanned := user.Status != sql.UserStatusBanned
	s.exception.IsUnprocessableEntity(isBanned, "Your account has been disabled", false)

	token := utils.GenerateToken(user.ID)
	expiredAt := time.Now().Add(s.conf.Auth.ExpiredDuration)

	payload := sql.UpdateUserTokenParams{
		ID:        user.ID,
		Token:     utils.PGText(token),
		ExpiredAt: utils.PGTimeStamp(expiredAt),
	}
	err = s.repo.UpdateUserToken(ctx, payload)
	s.exception.PanicIfError(err, false)

	return response.AuthLoginResponse{
		ID:        user.ID,
		Name:      user.Name,
		Username:  user.Username,
		Role:      user.Role,
		Status:    user.Status,
		Token:     token,
		ExpiredAt: expiredAt,
	}
}
