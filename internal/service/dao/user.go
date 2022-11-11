package dao

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/ivanov-gv/color-picker-test-app/internal/model"
	"github.com/ivanov-gv/color-picker-test-app/internal/service"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	usersTable = "user_account"
)

type UserDaoPostgres struct {
	pool *pgxpool.Pool
}

var _ service.UserDao = &UserDaoPostgres{}

func NewUserPostgres(pool *pgxpool.Pool) service.UserDao {
	return &UserDaoPostgres{
		pool: pool,
	}
}

func (p *UserDaoPostgres) Create(ctx context.Context) (model.User, error) {
	query, args, err := sq.
		Insert(usersTable).
		Values(sq.Expr("DEFAULT")).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return model.User{}, fmt.Errorf("can't build query: %w", err)
	}

	rows, err := p.pool.Query(ctx, query, args...)
	if err != nil {
		return model.User{}, fmt.Errorf("UserDaoPostgres - Create - p.Pool.Query: %w", err)
	}
	defer rows.Close()

	var userId int
	for rows.Next() {
		if err := rows.Scan(&userId); err != nil {
			return model.User{}, fmt.Errorf("can't scan order_id: %w", err)
		}
	}

	return model.User{Id: userId}, nil
}

func (p *UserDaoPostgres) Exist(ctx context.Context, userId int) (bool, error) {
	query, args, err := sq.
		Select("1").
		Prefix("SELECT EXISTS (").
		From(usersTable).
		Where(sq.Eq{"id": userId}).
		Suffix(")").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return false, fmt.Errorf("can't build query: %w", err)
	}

	rows, err := p.pool.Query(ctx, query, args...)
	if err != nil {
		return false, fmt.Errorf("UserDaoPostgres - Exist - p.Pool.Query: %w", err)
	}
	defer rows.Close()

	var ok bool
	for rows.Next() {
		if err := rows.Scan(&ok); err != nil {
			return false, fmt.Errorf("UserDaoPostgres - Exist - rows.Scan: %w", err)
		}
	}

	return ok, nil
}
