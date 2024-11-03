package handler

import (
	"math/rand/v2"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rpstvs/serverfibergo/cache"
	"github.com/rpstvs/serverfibergo/database"
)

type Response struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
	Book   string `json:"book"`
}

func GetRandomQuote(c *fiber.Ctx, db *database.Queries) error {
	var resp Response
	quoteCache := cache.GetCachedItem()
	timeNow := time.Now().UTC()
	if quoteCache == nil || timeNow.After(quoteCache.Expiration) {
		totalQuotes, _ := db.GetTotalQuotes(c.Context())
		id := rand.IntN(int(totalQuotes))

		quote, _ := db.GetQuoteByID(c.Context(), int32(id))

		cache.CreateCachedItem(quote)
		resp = Response{
			Quote:  quote.Quote,
			Author: quote.Author,
			Book:   quote.Book,
		}

		return c.JSON(resp)
	}
	resp = Response{
		Quote:  quoteCache.Quote.Quote,
		Author: quoteCache.Quote.Author,
		Book:   quoteCache.Quote.Book,
	}

	return c.JSON(resp)
}
