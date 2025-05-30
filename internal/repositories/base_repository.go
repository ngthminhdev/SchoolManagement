package repositories

import (
	"GolangBackend/helper"
	"GolangBackend/internal/dto"
	"GolangBackend/internal/entities"
	"GolangBackend/internal/global"
	"context"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
)

type IBaseRepository[T entities.IBaseEntity] interface {
	FindAll(ctx context.Context, options *dto.ListOptions) ([]T, error)
	FindById(ctx context.Context, options *dto.GetByIdOptions) (T, error)

	Create(ctx context.Context, entity T) (T, error)
	Update(ctx context.Context, id string, entity T) (T, error)
	Delete(ctx context.Context, id string) (bool, error)

	ExecuteOne(ctx context.Context, sql string, params ...any) (T, error)
	ExecuteMany(ctx context.Context, sql string, params ...any) ([]T, error)
}

type BaseRepository[T entities.IBaseEntity] struct {
	table  string
	entity func() T
}

func NewBaseRepository[T entities.IBaseEntity](table string, entity func() T) *BaseRepository[T] {
	return &BaseRepository[T]{
		table:  table,
		entity: entity,
	}
}

func (r *BaseRepository[T]) ScanRow(rows pgx.Rows) (T, error) {
	defer rows.Close()

	if !rows.Next() {
		return *new(T), nil
	}

	fieldDescriptions := rows.FieldDescriptions()

	scanArgs := make([]any, len(fieldDescriptions))
	for i := range scanArgs {
		var holder any
		scanArgs[i] = &holder
	}

	err := rows.Scan(scanArgs...)
	if err != nil {
		helper.LogError(err, "Error while scanning row")
		return *new(T), err
	}

	result := make(map[string]any, len(fieldDescriptions))
	for i, field := range fieldDescriptions {
		result[string(field.Name)] = *(scanArgs[i].(*any))
	}

	entity := r.entity()
	entity.FromMap(result)

	return entity, nil
}

func (r *BaseRepository[T]) ScanRows(rows pgx.Rows) ([]T, error) {
	defer rows.Close()

	fieldDescriptions := rows.FieldDescriptions()
	result := make([]T, 0)

	for rows.Next() {
		values := make([]any, len(fieldDescriptions))
		scanArgs := make([]any, len(fieldDescriptions))

		for i := range scanArgs {
			scanArgs[i] = &values[i]
		}

		if err := rows.Scan(scanArgs...); err != nil {
			return nil, err
		}

		rowMap := make(map[string]any, len(fieldDescriptions))
		for i, field := range fieldDescriptions {
			rowMap[string(field.Name)] = values[i]
		}

		entity := r.entity()
		entity.FromMap(rowMap)
		result = append(result, entity)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *BaseRepository[T]) Create(ctx context.Context, entity T) (T, error) {
	now := time.Now()

	data := entity.ToMap()
	data["created_at"] = now
	data["modified_at"] = now
	delete(data, "id")

	cols := make([]string, 0, len(data))
	args := make([]any, 0, len(data))
	placeholders := make([]string, 0, len(data))

	i := 1
	for k, v := range data {
		cols = append(cols, k)
		args = append(args, v)
		placeholders = append(placeholders, fmt.Sprintf("$%d", i))
		i++
	}

	query := fmt.Sprintf(
		`INSERT INTO %s (%s) VALUES (%s) RETURNING *`,
		r.table,
		strings.Join(cols, ", "),
		strings.Join(placeholders, ", "),
	)

	row, err := global.DB.Query(ctx, query, args...)
	helper.LogError(err, "err")
	if err != nil {
		return *new(T), err
	}

	return r.ScanRow(row)
}

func (r *BaseRepository[T]) FindById(ctx context.Context, options *dto.GetByIdOptions) (T, error) {
	var id string

	selectCols := "*"

	if options != nil {
		if len(options.Fields) > 0 {
			selectCols = strings.Join(options.Fields, ", ")
		}

		if options.ID != "" {
			id = options.ID
		}
	}

	query := fmt.Sprintf(`SELECT %s FROM %s WHERE id = $1 LIMIT 1`, selectCols, r.table)

	row, err := global.DB.Query(ctx, query, id)
	if err != nil {
		return *new(T), err
	}

	return r.ScanRow(row)
}

func (r *BaseRepository[T]) FindAll(ctx context.Context, options *dto.ListOptions) ([]T, error) {
	var result []T

	offset := 0
	limit := 10
	selectCols := "*"

	if options != nil {
		if options.Offset > 0 {
			offset = options.Offset
		}
		if options.Limit > 0 {
			limit = options.Limit
		}
		if len(options.Fields) > 0 {
			selectCols = strings.Join(options.Fields, ", ")
		}
	}

	query := fmt.Sprintf(`SELECT %s FROM %s LIMIT $1 OFFSET $2`, selectCols, r.table)

	rows, err := global.DB.Query(ctx, query, limit, offset)
	if err != nil {
		return result, err
	}

	return r.ScanRows(rows)
}

func (r *BaseRepository[T]) Update(ctx context.Context, id string, entity T) (T, error) {
	now := time.Now()
	data := entity.ToMap()

	excludeFields := []string{"id", "created_at", "deleted"}
	toUpdateData := make(map[string]any)

	for k, v := range data {
		if slices.Contains(excludeFields, k) {
			continue
		}

		if !helper.IsEmptyValue(v) {
			toUpdateData[k] = v
		}
	}

	toUpdateData["modified_at"] = now

	setClauses := make([]string, 0, len(toUpdateData))
	args := make([]any, 0, len(toUpdateData)+2)
	i := 1
	for k, v := range toUpdateData {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", k, i))
		args = append(args, v)
		i++
	}

	if len(toUpdateData) == 1 {
		return entity, nil
	}

	args = append(args, id)

	query := fmt.Sprintf(`
		UPDATE %s 
		SET %s
		WHERE id = $%d
		RETURNING *`,
		r.table,
		strings.Join(setClauses, ", "),
		i,
	)

	row, err := global.DB.Query(ctx, query, args...)
	if err != nil {
		return *new(T), err
	}

	return r.ScanRow(row)
}

func (r *BaseRepository[T]) Delete(ctx context.Context, id string) (bool, error) {
	query := fmt.Sprintf(`
		UPDATE %s 
		SET deleted = true
		WHERE id = $1
		`,
		r.table,
	)

	_, err := global.DB.Query(ctx, query, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *BaseRepository[T]) ExecuteOne(ctx context.Context, sql string, params ...any) (T, error) {
	rows, err := global.DB.Query(ctx, sql, params...)
	if err != nil {
		helper.LogError(err, "Error executing raw SQL query for single result")
		return *new(T), err
	}

	return r.ScanRow(rows)
}

func (r *BaseRepository[T]) ExecuteMany(ctx context.Context, sql string, params ...any) ([]T, error) {
	rows, err := global.DB.Query(ctx, sql, params...)
	if err != nil {
		helper.LogError(err, "Error executing raw SQL query for multiple results")
		return nil, err
	}

	return r.ScanRows(rows)
}
