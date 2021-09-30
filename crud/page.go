package crud

import (
	"context"
)

type Page struct {
	Index uint64
	Size uint16
	Total uint64
	Items []BaseModel
}

type Pageable interface {
	GetPage(ctx context.Context, page *Page) error
}
