package queries

import (
	"library-rest-api/app/models"

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


