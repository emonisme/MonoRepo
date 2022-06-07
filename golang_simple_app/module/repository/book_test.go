package repository_test

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"simpleapp/module/entity"
	"simpleapp/module/repository"
)

func TestBook_Create(t *testing.T) {
	type args struct {
		ctx  context.Context
		book entity.Book
	}

	testCases := map[string]struct {
		args    args
		mockFn  func(sql sqlmock.Sqlmock, args args)
		want    *entity.Book
		wantErr error
	}{
		"success": {
			args: args{
				ctx:  context.Background(),
				book: entity.Book{Name: "book1", Category: "category1"},
			},
			mockFn: func(sql sqlmock.Sqlmock, args args) {
				column := []string{"id"}
				row := sql.NewRows(column).AddRow(1)
				sql.ExpectQuery(regexp.QuoteMeta("INSERT INTO books (name, category) VALUES (?, ?) RETURNING id")).
					WithArgs(args.book.Name, args.book.Category).
					WillReturnRows(row)
			},
			want:    &entity.Book{ID: "1", Name: "book1", Category: "category1"},
			wantErr: nil,
		},
		"failed": {
			args: args{
				ctx:  context.Background(),
				book: entity.Book{Name: "", Category: "category1"},
			},
			mockFn: func(sql sqlmock.Sqlmock, args args) {
				sql.ExpectQuery(regexp.QuoteMeta("INSERT INTO books (name, category) VALUES (?, ?) RETURNING id")).
					WithArgs(args.book.Name, args.book.Category).
					WillReturnError(errors.New("database error"))
			},
			want:    nil,
			wantErr: errors.New("database error"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockDB, sMock, _ := sqlmock.New()
			sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

			bookRepo := repository.NewBookRepository(sqlxDB)

			tc.mockFn(sMock, tc.args)
			actual, err := bookRepo.Create(tc.args.ctx, tc.args.book)
			if tc.wantErr != nil {
				assert.Nil(t, actual)
				assert.Equal(t, tc.wantErr.Error(), err.Error())
				return
			}

			assert.Nil(t, err)
			assert.Equal(t, tc.want, actual)
		})
	}
}
