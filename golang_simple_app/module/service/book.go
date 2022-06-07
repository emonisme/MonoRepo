package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"simpleapp/module/entity"
)

type BookUsecase interface {
	Create(context.Context, entity.Book) (*entity.Book, error)
	Update(context.Context, entity.Book) (*entity.Book, error)
	Delete(context.Context, string) error
	Get(context.Context, string) (*entity.Book, error)
}

type Book struct {
	BookUsecase BookUsecase
}

func NewBookService(usecase BookUsecase) *Book {
	return &Book{
		BookUsecase: usecase,
	}
}

func (bs Book) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Create!\n")
}

func (bs Book) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Update!\n")
}

func (bs Book) Delete(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Delete!\n")
}

func (bs Book) Get(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Get!\n")
}
