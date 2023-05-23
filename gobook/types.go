package main

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID              uuid.UUID `json:"id"`
	Title           string    `json:"title" binding:"required"`
	Author          string    `json:"author" binding:"required"`
	PublicationDate time.Time `json:"publicationDate" binding:"required"`
	Publisher       string    `json:"publisher" binding:"required"`
}

func NewBook(title string, author string, publicationDate time.Time, publisher string) *Book {
	return &Book{
		ID:              uuid.New(),
		Title:           title,
		Author:          author,
		PublicationDate: publicationDate,
		Publisher:       publisher,
	}
}
