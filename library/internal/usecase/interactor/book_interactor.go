package interactor

import (
	"context"
	"github.com/armanchik1503/golangProject/library/internal/domain"
	"github.com/armanchik1503/golangProject/library/internal/repository"
)

type bookInteractor struct {
	bookRepo repository.BookRepo
}

func (interactor bookInteractor) Create(
	ctx context.Context,
	book domain.Book,
	) (*domain.Book, error){
	createdBook, err := interactor.bookRepo.Store(ctx, book)

	if err != nil {
		return nil, err
	}

	return createdBook, nil
}

func (interactor bookInteractor) Get(
	ctx context.Context,
	id domain.Book.ID,
) (*domain.Book, error) {
	getBook, err := interactor.bookRepo.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	return getBook, nil
}
