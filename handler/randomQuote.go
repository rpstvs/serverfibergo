package handler

import (
	"log"
	"math/rand/v2"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rpstvs/serverfibergo/cache"
	"github.com/rpstvs/serverfibergo/database"
)

type Response struct {
	Quote   string `json:"quote"`
	Author  string `json:"author"`
	Book    string `json:"book"`
	Bio     string `json:"bio"`
	ImgLink string `json:"imglink"`
	Books   []Book `json:"books"`
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
		cacheQuote.Quote.Quote = quote.Quote
		cacheQuote.Expiration = time.Now().Add(6 * time.Hour)
		authorInf := getAuthorinfo(quote.Author)
		resp = Response{
			Quote:   quote.Quote,
			Author:  quote.Author,
			Book:    quote.Book,
			Bio:     authorInf.Bio,
			ImgLink: authorInf.ImgLink,
			Books:   authorInf.Books,
		}

		return c.JSON(resp)
	}
	authorInf := getAuthorinfo(cacheQuote.Quote.Author)

	resp = Response{
		Quote:   cacheQuote.Quote.Quote,
		Author:  cacheQuote.Quote.Author,
		Book:    cacheQuote.Quote.Book,
		Bio:     authorInf.Bio,
		ImgLink: authorInf.ImgLink,
		Books:   authorInf.Books,
	}

	return c.JSON(resp)
}

func getAuthorinfo(authorName string) AuthorInfo {
	AuthorResp := AuthorInfo{}
	authors, err := database.LoadAuthors("database.json")
	if err != nil {
		log.Println(err)
	}

	for _, author := range authors.Authors {
		if strings.Contains(authorName, author.Authorname) {
			AuthorResp.AuthorName = author.Authorname
			AuthorResp.Bio = author.Description
			AuthorResp.ImgLink = author.ImgLink
			for _, book := range author.Books {
				AuthorResp.Books = append(AuthorResp.Books, book)
			}
		}
	}
	return AuthorResp
}
