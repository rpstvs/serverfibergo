package handler

import (
	"math/rand/v2"

	"github.com/gofiber/fiber/v2"
	"github.com/rpstvs/serverfibergo/database"
)

func GetQuote(c *fiber.Ctx, db *database.Queries) error {
	var resp Response
	totalQuotes, _ := db.GetTotalQuotes(c.Context())
	id := rand.IntN(int(totalQuotes))

	quote, _ := db.GetQuoteByID(c.Context(), int32(id))

	resp.Author = quote.Author
	resp.Book = quote.Book
	resp.Quote = quote.Quote

	return c.JSON(resp)
}
