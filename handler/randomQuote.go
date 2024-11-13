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

func GetRandomQuote(c *fiber.Ctx, db *database.Queries, cacheQuote *cache.CacheQuote) error {
	var resp Response
	//quoteCache := cache.GetCachedItem()
	timeNow := time.Now().UTC()
	if cacheQuote == nil || timeNow.After(cacheQuote.Expiration) {
		totalQuotes, _ := db.GetTotalQuotes(c.Context())
		id := rand.IntN(int(totalQuotes))

		quote, _ := db.GetQuoteByID(c.Context(), int32(id))

		cacheQuote.Quote.Author = quote.Author
		cacheQuote.Quote.Book = quote.Book
		cacheQuote.Quote.Quote = quote.Book
		cacheQuote.Expiration = time.Now().Add(24 * time.Hour)
		resp = Response{
			Quote:  quote.Quote,
			Author: quote.Author,
			Book:   quote.Book,
		}

		return c.JSON(resp)
	}
	resp = Response{
		Quote:  cacheQuote.Quote.Quote,
		Author: cacheQuote.Quote.Author,
		Book:   cacheQuote.Quote.Book,
	}

	return c.JSON(resp)
}
