package book

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type BookObject struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Category string `db:"category"`
}

type Book struct {
	db *sqlx.DB
}

func NewBookRepository(db *sqlx.DB) *Book {
	return &Book{
		db: db,
	}
}

func (br *Book) Create(ctx Context, obj BookObject) (*BookObject, error) {
	return nil, err
}

func (br *Book) Update(ctx Context, obj BookObject) (*BookObject, error) {
	return nil, err
}

func (br *Book) Delete(ctx Context, id int) error {
	return err
}

func (br *Book) Get(ctx Context, id int) (*BookObject, error) {
	return nil, err
}
