package config

import (
	"github.com/jmoiron/sqlx"
	"simpleapp/module/repository"
	"simpleapp/module/service"
	"simpleapp/module/usecase"
)

func newBookUsecase(db *sqlx.DB) service.BookUsecase {
	bookRepo := repository.NewBookRepository(db)

	return usecase.NewBookUsecase(bookRepo)
}
