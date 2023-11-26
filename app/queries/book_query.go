package queries

import (
	"library-rest-api/app/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type BookQueries struct {
	*sqlx.DB
}

func (q *BookQueries) GetBooks() ([]models.Book, error) {
	books := []models.Book{}

	query := `SELECT * FROM books`

	err := q.Get(&books, query)

	if err != nil {
		return books, err
	}

	return books, nil
}

func (q *BookQueries) GetBook(id uuid.UUID) (models.Book, error) {
	book := models.Book{}

	query := `SELECT * FROM books WHERE id = $1`

	err := q.Get(&book, query, id)

	if err != nil {
		return book, err
	}

return book, nil
}

