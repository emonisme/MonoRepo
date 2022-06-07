package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"simpleapp/module/entity"
)

type Book struct {
	db *sqlx.DB
}

func NewBookRepository(db *sqlx.DB) *Book {
	return &Book{
		db: db,
	}
}

func (br Book) Create(ctx context.Context, book entity.Book) (*entity.Book, error) {
	return nil, nil
}

func (br Book) Update(ctx context.Context, book entity.Book) (*entity.Book, error) {
	return nil, nil
}

func (br Book) Delete(ctx context.Context, id string) error {
	return nil
}

func (br Book) Get(ctx context.Context, id string) (*entity.Book, error) {
	book := entity.Book{}
	row := br.db.QueryRowxContext(ctx, `SELECT id, name, category FROM books WHERE id = $1`, id)
	if err := row.StructScan(&book); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Not Found")
		}

		return nil, err
	}

	return &book, nil
}
