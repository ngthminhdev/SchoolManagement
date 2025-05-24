package entities

import (
	"time"

	"github.com/google/uuid"
)

type IBaseEntity interface {
	FromMap(map[string]any)
	ToMap() map[string]any
}

type BaseEntity struct {
	ID         string
	CreatedAt  time.Time
	ModifiedAt time.Time
	Deleted    bool
}

func (e *BaseEntity) FromMap(data map[string]any) {
	if v, ok := data["id"].([16]uint8); ok {
		if uid, err := uuid.FromBytes(v[:]); err == nil {
			e.ID = uid.String()
		}
	}
	if v, ok := data["created_at"].(time.Time); ok {
		e.CreatedAt = v
	}
	if v, ok := data["modified_at"].(time.Time); ok {
		e.ModifiedAt = v
	}
	if v, ok := data["deleted"].(bool); ok {
		e.Deleted = v
	}
}

func (b *BaseEntity) ToMap() map[string]any {
	return map[string]any{
		"id":          b.ID,
		"created_at":  b.CreatedAt,
		"modified_at": b.ModifiedAt,
		"deleted":     b.Deleted,
	}
}

// ToSQLParams returns the values in order for SQL query args
func (b *BaseEntity) ToSQLParams() []any {
	return []any{
		b.ID,
		b.CreatedAt,
		b.ModifiedAt,
		b.Deleted,
	}
}
