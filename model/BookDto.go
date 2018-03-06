package model

import "fmt"

type BookDto struct {
	UUID        string    `json:"uuid"`
	Title       string    `json:"title"`
	NoPages     int       `json:"noPages"`
	ReleaseDate string    `json:"releaseDate"`
	Author 	AuthorDto `json:"author"`
}

type BooksList []BookDto

var Books BooksList

func (book BookDto) String() string {
	return fmt.Sprintf("BookDto{\n\tUUID='%s',\n\tTitle='%s',\n\tNoPages=%d,\n\tReleaseDate='%s',\n\tAuthor=%s\n}\n",
		book.UUID, book.Title, book.NoPages, book.ReleaseDate, book.Author)
}

func (b *BooksList) Get(bookUUID string) (BookDto, error) {
	err := fmt.Errorf("Couldn't find book with UUID = %s", bookUUID)

	for _, book := range *b {
		if book.UUID == bookUUID {
			return book, nil
		}
	}

	return BookDto{}, err
}

func (b *BooksList) Add(book BookDto) error {
	//err :=  fmt.Errorf("Error while adding book %s", book)

	*b = append(*b, book)

	return nil
}
