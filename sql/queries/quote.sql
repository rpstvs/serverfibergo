-- name: GetTotalQuotes :exec
SELECT COUNT(*)
FROM Quotes;
-- name: GetQuoteByID :one
SELECT *
From Quotes
WHERE Id = $1;