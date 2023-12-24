package service

import (
	"be-park-ease/config"
	"be-park-ease/exception"
	"be-park-ease/internal/model"
	"be-park-ease/internal/repository"
	"be-park-ease/internal/request"
	"be-park-ease/internal/response"
	"be-park-ease/internal/sql"
	"context"
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

type UserService interface {
	AllUser(ctx context.Context, req request.AllUserRequest) response.BaseResponsePagination[response.User]
	UserById(ctx context.Context, userId int32) response.User
	CreateUser(ctx context.Context, req request.CreateUserRequest)
}

type userService struct {
	repo      repository.UserRepository
	exception exception.Exception
	conf      config.Config
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo:      repo,
		exception: exception.NewException("user-service"),
		conf:      *config.Get(),
	}
}

func (s *userService) AllUser(ctx context.Context, req request.AllUserRequest) response.BaseResponsePagination[response.User] {
	data, err := s.repo.GetAllUsers(ctx, req)
	s.exception.PanicIfError(err, true)
	s.exception.IsNotFound(data, true)

	content := model.ContentPagination[response.User]{
		Count:   data.Count,
		Content: []response.User{},
	}

	for _, val := range data.Content {
		content.Content = append(content.Content, response.User{
			ID:       val.ID,
			Username: val.Username,
			Name:     val.Name,
			Role:     val.Role,
			Status:   val.Status,
		})
	}

	return response.WithPagination[response.User](content, req.BasePagination)
}

func (s *userService) UserById(ctx context.Context, userId int32) response.User {
	data, err := s.repo.GetUserById(ctx, userId)
	s.exception.PanicIfErrorWithoutNoSqlResult(err, false)
	s.exception.IsNotFound(data, false)

	return response.User{
		ID:       data.ID,
		Username: data.Username,
		Name:     data.Name,
		Role:     data.Role,
		Status:   data.Status,
	}
}

func (s *userService) handleErrorUniqueUser(err error, isList bool) {
	var pgErr *pgconn.PgError
	ok := errors.As(err, &pgErr)
	if !ok {
		return
	}

	if pgErr.Code != pgerrcode.UniqueViolation {
		return
	}

	switch pgErr.ConstraintName {
	case "users_username_key":
		s.exception.IsBadRequestMessage("Username already exist.", isList)
	}
}

func (s *userService) CreateUser(ctx context.Context, req request.CreateUserRequest) {
	payload := sql.CreateUserParams{
		Username: req.Username,
		Password: req.Password,
		Name:     req.Name,
		Role:     req.Role,
		Status:   req.Status,
	}
	err := s.repo.CreateUser(ctx, payload)
	s.handleErrorUniqueUser(err, false)
	s.exception.PanicIfError(err, false)
}
