package model

import "fmt"

type AuthorDto struct {
	UUID      string `json:"uuid"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Birthday  string `json:"birthday"`
	Death 	string `json:"death"`
}

type AuthorsList []AuthorDto

var Authors AuthorsList

func (author AuthorDto) String() string {
	return fmt.Sprintf("AuthorDto{\n\tUUID='%s',\n\tFirstName='%s',\n\tLastName='%s',\n\tBirthday='%s',\n\tDeath='%s'\n}\n",
		author.UUID, author.FirstName, author.LastName, author.Birthday, author.Death)
}

func (a *AuthorsList) Get(authorUUID string) (AuthorDto, error) {
	err := fmt.Errorf("Couldn't find author with UUID = %s", authorUUID)

	for _, author := range *a {
		if author.UUID == authorUUID {
			return author, nil
		}
	}

	return AuthorDto{}, err
}

func (a *AuthorsList) Add(author AuthorDto) error {
	//err :=  fmt.Errorf("Error while adding book %s", book)

	*a = append(*a, author)

	return nil
}