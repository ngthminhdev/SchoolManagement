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
	ID         string    `json:"id" db:"id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	ModifiedAt time.Time `json:"modified_at" db:"modified_at"`
	Deleted    bool      `json:"-" db:"deleted"`
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
