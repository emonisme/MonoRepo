package usecase

import (
	"context"

	"simpleapp/module/entity"
)

type BookRepository interface {
	Create(context.Context, entity.Book) (*entity.Book, error)
	Update(context.Context, entity.Book) (*entity.Book, error)
	Delete(context.Context, string) error
	Get(context.Context, string) (*entity.Book, error)
}

type Book struct {
	BookRepo BookRepository
}

func NewBookUsecase(repo BookRepository) *Book {
	return &Book{
		BookRepo: repo,
	}
}

func (bu Book) Create(ctx context.Context, book entity.Book) (*entity.Book, error) {
	return nil, nil
}

func (bu Book) Update(ctx context.Context, book entity.Book) (*entity.Book, error) {
	return nil, nil
}

func (bu Book) Delete(ctx context.Context, id string) error {
	return nil
}

func (bu Book) Get(ctx context.Context, id string) (*entity.Book, error) {
	return nil, nil
}
