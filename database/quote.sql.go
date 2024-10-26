// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: quote.sql

package database

import (
	"context"
)

const getQuoteByID = `-- name: GetQuoteByID :one
SELECT id, quote, author, book, post_date
From Quotes
WHERE Id = $1
`

func (q *Queries) GetQuoteByID(ctx context.Context, id int32) (Quote, error) {
	row := q.db.QueryRowContext(ctx, getQuoteByID, id)
	var i Quote
	err := row.Scan(
		&i.ID,
		&i.Quote,
		&i.Author,
		&i.Book,
		&i.PostDate,
	)
	return i, err
}

const getTotalQuotes = `-- name: GetTotalQuotes :exec
SELECT COUNT(*)
FROM Quotes
`

func (q *Queries) GetTotalQuotes(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, getTotalQuotes)
	return err
}