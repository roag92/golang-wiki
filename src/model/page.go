package model

import (
	"wiki/src/utils"
)

// Post structure
type Page struct {
	Title string
	Body  []byte
}

// Save function to store a new Post entry
func (p *Page) Save() error {
	return utils.WriteFile("/tmp/"+p.Title+".txt", p.Body)
}

// GetPage function to find a Post by title
func GetPage(title string) (*Page, error) {
	body, err := utils.ReadFile("/tmp/" + title + ".txt")

	if err != nil {
		return &Page{Title: title}, err
	}

	return &Page{Title: title, Body: body}, nil
}
