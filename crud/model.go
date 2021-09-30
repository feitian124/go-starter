package crud

import (
	"context"
	"time"
)

type ID uint64

type BaseModel struct {
	Id        ID
	CreatedAt time.Time
	UpdatedAt time.Time
}

// interfaces in golang are often short, usually contains 1 or at most 3 methods.

type Creatable interface {
	Create(ctx context.Context, m *BaseModel) error
}

type Readable interface {
	Get(ctx context.Context, id ID) (BaseModel, error)
}

type Updatable interface {
	Update(ctx context.Context, m *BaseModel) error
}

type Deletable interface {
	Delete(ctx context.Context, id ID) error
}

type Crud interface {
	Creatable
	Readable
	Updatable
	Deletable
}