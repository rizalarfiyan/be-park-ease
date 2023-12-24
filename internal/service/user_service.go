package service

import (
	"be-park-ease/config"
	"be-park-ease/exception"
	"be-park-ease/internal/model"
	"be-park-ease/internal/repository"
	"be-park-ease/internal/request"
	"be-park-ease/internal/response"
	"be-park-ease/internal/sql"
	"be-park-ease/utils"
	"context"
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

type UserService interface {
	AllUser(ctx context.Context, req request.AllUserRequest) response.BaseResponsePagination[response.User]
	UserById(ctx context.Context, userId int32) response.User
	CreateUser(ctx context.Context, req request.CreateUserRequest)
	UpdateUser(ctx context.Context, req request.UpdateUserRequest)
	ChangePassword(ctx context.Context, req request.ChangePasswordRequest)
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

func (s *userService) UpdateUser(ctx context.Context, req request.UpdateUserRequest) {
	data, err := s.repo.GetUserById(ctx, req.UserId)
	s.exception.PanicIfErrorWithoutNoSqlResult(err, false)
	s.exception.IsNotFound(data, false)

	payload := sql.UpdateUserParams{
		Name:     data.Name,
		Username: data.Username,
		Password: data.Password,
		Role:     data.Role,
		Status:   data.Status,
		ID:       req.UserId,
	}

	if req.Username != "" {
		payload.Username = req.Username
	}

	if req.Password != "" {
		payload.Password = req.Password
	}

	if req.Name != "" {
		payload.Name = req.Name
	}

	if req.Role != "" {
		payload.Role = req.Role
	}

	if req.Status != "" {
		payload.Status = req.Status
	}

	err = s.repo.UpdateUser(ctx, payload)
	s.handleErrorUniqueUser(err, false)
	s.exception.PanicIfError(err, false)
}

func (s *userService) ChangePassword(ctx context.Context, req request.ChangePasswordRequest) {
	user, err := s.repo.GetUserById(ctx, req.UserId)
	s.exception.PanicIfErrorWithoutNoSqlResult(err, false)
	s.exception.IsUnprocessableEntity(user, "Ops... Something wrong for your request", false)

	isValidPassword := utils.ComparePassword(user.Password, req.OldPassword)
	s.exception.IsBadRequest(isValidPassword, "Old password doesn't match", false)
	s.exception.IsBadRequest(req.OldPassword != req.Password, "Password must be update from old password", false)

	password, err := utils.HashPassword(req.Password)
	s.exception.PanicIfError(err, false)

	payload := sql.UpdatePasswordParams{
		Password: password,
		ID:       req.UserId,
	}

	err = s.repo.ChangePasswordUser(ctx, payload)
	s.exception.PanicIfError(err, false)
}
