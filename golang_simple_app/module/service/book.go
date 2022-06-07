package service

import (
	"context"
	"fmt"
	"net/http"
	"encoding/json"

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

func (bs Book) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var bookPayload entity.Book
	err := decoder.Decode(&bookPayload)
    if err != nil {
        panic(err)
    }

	bookCreated, err := bs.BookUsecase.Create(r.Context(), bookPayload)
	if err != nil {
		codes := http.StatusInternalServerError
		http.Error(w, http.StatusText(codes), codes)
		return
	}

	writeResponse(w, http.StatusOK, bookCreated)
}

func (bs Book) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "Update Success!\n")
}

func (bs Book) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "Delete Success!\n")
}

func (bs Book) Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	strID := ps.ByName("id")
	book, err := bs.BookUsecase.Get(r.Context(), strID)
	if err != nil {
		codes := http.StatusInternalServerError
		if err.Error() == "Not Found" {
			codes = http.StatusNotFound
		}

		http.Error(w, http.StatusText(codes), codes)
		return
	}

	writeResponse(w, http.StatusOK, book)
}

func writeResponse(w http.ResponseWriter, code int, book *entity.Book) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(book)
}
