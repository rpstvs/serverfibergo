package cache

import (
	"time"

	"github.com/rpstvs/serverfibergo/database"
)

type Quote struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
	Book   string `json:"book"`
}
type CacheQuote struct {
	Quote      Quote
	Expiration time.Time
}

var CachedQuote CacheQuote

func CreateCachedItem(quote database.Quote) {

	CachedQuote = CacheQuote{
		Quote: Quote{
			Author: quote.Author,
			Book:   quote.Book,
			Quote:  quote.Quote,
		},
		Expiration: time.Now().UTC().Add(24 * time.Hour),
	}

}

func GetCachedItem() *CacheQuote {
	return &CachedQuote
}
