package handler

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/rpstvs/serverfibergo/cache"
	"github.com/rpstvs/serverfibergo/database"
)

type Book struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

type AuthorInfo struct {
	AuthorName string `json:"authorname"`
	Bio        string `json:"bio"`
	ImgLink    string `json:"imglink"`
	Books      []Book `json:"books"`
}

func GetAuthor(c *fiber.Ctx) error {
	quoteCache := cache.GetCachedItem()
	AuthorResp := AuthorInfo{}
	authors, err := database.LoadAuthors("database.json")
	if err != nil {
		log.Println(err)
	}

	for _, author := range authors.Authors {
		if strings.Contains(quoteCache.Quote.Author, author.Authorname) {
			AuthorResp.AuthorName = author.Authorname
			AuthorResp.Bio = author.Description
			AuthorResp.ImgLink = author.ImgLink
			for _, book := range author.Books {
				AuthorResp.Books = append(AuthorResp.Books, book)
			}
		}
	}
	return c.JSON(AuthorResp)
}
