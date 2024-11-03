package database

import (
	"encoding/json"
	"errors"
	"os"
)

type DBStructure struct {
	Authors []struct {
		Authorname  string `json:"authorname"`
		Description string `json:"description"`
		ImgLink     string `json:"imgLink"`
		Books       []struct {
			Name string `json:"name"`
			Link string `json:"link"`
		} `json:"books"`
	} `json:"authors"`
}

func LoadAuthors(path string) (DBStructure, error) {

	dbStructure := DBStructure{}
	dat, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return dbStructure, err
	}
	err = json.Unmarshal(dat, &dbStructure)
	if err != nil {
		return dbStructure, nil
	}
	return dbStructure, nil
}
