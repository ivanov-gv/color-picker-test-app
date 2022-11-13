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
	colorsTable = "color"
)

type ColorDaoPostgres struct {
	pool *pgxpool.Pool
}

var _ service.ColorDao = &ColorDaoPostgres{}

func NewColorPostgres(pool *pgxpool.Pool) service.ColorDao {
	return &ColorDaoPostgres{
		pool: pool,
	}
}

func (p *ColorDaoPostgres) GetAll(ctx context.Context, userId int) ([]model.Color, error) {
	query, args, err := sq.
		Select("id", "name", "hex").
		From(colorsTable).
		Where(sq.Eq{
			"user_id": userId,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("can't build query: %w", err)
	}

	rows, err := p.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("ColorDaoPostgres - GetAll - p.Pool.Query: %w", err)
	}
	defer rows.Close()

	models := make([]model.Color, 0)

	for rows.Next() {
		m := model.Color{}

		err = rows.Scan(&m.Id, &m.Name, &m.HEX)
		if err != nil {
			return nil, fmt.Errorf("ColorDaoPostgres - GetAll - rows.Scan: %w", err)
		}

		models = append(models, m)
	}

	return models, nil
}

func (p *ColorDaoPostgres) Get(ctx context.Context, userId int, colorId int) (model.Color, error) {
	query, args, err := sq.
		Select("id", "name", "hex").
		From(colorsTable).
		Where(sq.Eq{
			"user_id": userId,
			"id":      colorId,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return model.Color{}, fmt.Errorf("can't build query: %w", err)
	}

	rows, err := p.pool.Query(ctx, query, args...)
	if err != nil {
		return model.Color{}, fmt.Errorf("ColorDaoPostgres - Get - p.Pool.Query: %w", err)
	}
	defer rows.Close()

	m := model.Color{}

	for rows.Next() {
		if err = rows.Scan(&m.Id, &m.Name, &m.HEX); err != nil {
			return model.Color{}, fmt.Errorf("ColorDaoPostgres - Get - rows.Scan: %w", err)
		}
	}

	return m, nil
}

func (p *ColorDaoPostgres) Add(ctx context.Context, userId int, color model.Color) (model.Color, error) {
	query, args, err := sq.
		Insert(colorsTable).
		Columns("user_id", "name", "hex").
		Values(
			userId,
			color.Name,
			color.HEX).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return model.Color{}, fmt.Errorf("can't build query: %w", err)
	}

	rows, err := p.pool.Query(ctx, query, args...)
	if err != nil {
		return model.Color{}, fmt.Errorf("ColorDaoPostgres - Add - p.Pool.Query: %w", err)
	}
	defer rows.Close()

	var colorId int
	for rows.Next() {
		if err := rows.Scan(&colorId); err != nil {
			return model.Color{}, fmt.Errorf("can't scan order_id: %w", err)
		}
	}

	if err = rows.Err(); err != nil {
		return model.Color{}, fmt.Errorf("can't add color: %w", err)
	}

	color.Id = colorId
	return color, nil
}

func (p *ColorDaoPostgres) Delete(ctx context.Context, userId int, colorId int) error {
	query, args, err := sq.
		Delete(colorsTable).
		Where(sq.Eq{
			"user_id": userId,
			"id":      colorId,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return fmt.Errorf("can't build query: %w", err)
	}

	rows, err := p.pool.Query(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("ColorDaoPostgres - Delete - p.Pool.Query: %w", err)
	}
	defer rows.Close()

	return nil
}
