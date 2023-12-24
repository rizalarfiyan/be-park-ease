package repository

import (
	"be-park-ease/internal/model"
	"be-park-ease/internal/request"
	"be-park-ease/internal/sql"
	"be-park-ease/utils"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	GetAllUsers(ctx context.Context, req request.AllUserRequest) (*model.ContentPagination[sql.User], error)
	GetUserById(ctx context.Context, id int32) (sql.User, error)
	CreateUser(ctx context.Context, user sql.CreateUserParams) error
}

type userRepository struct {
	db           *pgxpool.Pool
	query        *sql.Queries
	queryBuilder *sql.Queries
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{
		db:           db,
		query:        sql.New(db),
		queryBuilder: sql.New(utils.QueryWrap(db)),
	}
}

func (r *userRepository) GetAllUsers(ctx context.Context, req request.AllUserRequest) (*model.ContentPagination[sql.User], error) {
	var res model.ContentPagination[sql.User]

	baseBuilder := func(b *utils.QueryBuilder) {
		if req.Role != "" {
			b.Where("role = $1", req.Role)
		}

		if req.Status != "" {
			b.Where("status = $1", req.Status)
		}

		if req.Search != "" {
			b.Where("LOWER(name) LIKE $1 OR LOWER(username) LIKE $1", fmt.Sprintf("%%%s%%", req.Search))
		}
	}

	entryHistory, err := r.queryBuilder.GetAllUser(utils.QueryBuild(ctx, func(b *utils.QueryBuilder) {
		baseBuilder(b)
		if req.OrderBy != "" && req.Order != "" {
			b.Ordering(req.OrderBy, req.Order)
		} else {
			b.Order("created_at DESC")
		}
		b.Pagination(req.Page, req.Limit)
	}))
	if err != nil {
		return nil, err
	}

	count, err := r.queryBuilder.CountAllUser(utils.QueryBuild(ctx, baseBuilder))
	if err != nil {
		return nil, err
	}

	res.Content = entryHistory
	res.Count = count
	return &res, nil
}

func (r *userRepository) GetUserById(ctx context.Context, id int32) (sql.User, error) {
	return r.query.GetUserById(ctx, id)
}

func (r *userRepository) CreateUser(ctx context.Context, user sql.CreateUserParams) error {
	return r.query.CreateUser(ctx, user)
}
