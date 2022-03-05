package repository

import (
	"context"
	"github.com/armanchik1503/golangProject/library/internal/domain"
)

type BookRepo interface {
	Store(ctx context.Context, book domain.Book) (*domain.Book, error)
	Get(ctx context.Context, id string) (*domain.Book, error)
	GetMany(ctx context.Context, limit, offset int) ([]domain.Book, error)
	Remove(ctx context.Context, id string) error
}
