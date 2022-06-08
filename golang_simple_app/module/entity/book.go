package entity

type Book struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Category string `db:"category"`
}
