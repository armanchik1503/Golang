package usecase

import (
	"context"
	"github.com/armanchik1503/golangProject/library/internal/domain"
)

type BookUsecase interface {
	Create(ctx context.Context, book domain.Book) (*domain.Book, error)
	Get(ctx context.Context, id string) (*domain.Book, error)
	GetMany(ctx context.Context, limit, offset int) ([]domain.Book, error)
	Delete(ctx context.Context, id string) error
}
