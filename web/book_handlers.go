package web

import (
	"net/http"
	"github.com/gorilla/mux"
	"../model"
	"io/ioutil"
	"encoding/json"
	"../persistence"
)

//Demonstrates the basic functionality of private and public modifiers in GO
func Index(w http.ResponseWriter, r *http.Request) error {
	helloWorkshop := struct {
		Message        string `json:"message"`
		privateMessage string `json:"privateMessage"`
		NoTagField     string `json:"-"`
	}{
		Message:        "Hello workshop!",
		privateMessage: "Message that does not appear in response :).",
		NoTagField:     "This message won't appear either",
	}
	WriteJson(w, helloWorkshop)
	return nil
}

func GetAllBooks(w http.ResponseWriter, _ *http.Request) error {
	books, err := persistence.Store.GetBooks()
	if err != nil {
		return err
	}
	WriteJson(w, books)
	return nil
}

func GetBooksByTitle(w http.ResponseWriter, r *http.Request) error {
	title := mux.Vars(r)["title"]

	if books, err := persistence.Store.GetBooksByTitle(title); err != nil {
		return err
	} else {
		WriteJson(w, books)
		return nil
	}
}

func GetBookByUUID(w http.ResponseWriter, r *http.Request) error {
	bookUUID := mux.Vars(r)["uuid"]
	if book, err := persistence.Store.GetBook(bookUUID); err != nil {
		return err
	} else {
		WriteJson(w, book)
		return nil
	}
}

func DeleteBookByUUID(w http.ResponseWriter, r *http.Request) error {
	var bookUUID = mux.Vars(r)["uuid"]
	if err := persistence.Store.DeleteBook(bookUUID); err != nil {
		return err
	} else {
		WriteJson(w, struct{ Message string }{Message: "Deleted"})
		return nil
	}
}

func AddBook(w http.ResponseWriter, r *http.Request) error {
	var book model.Book
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bytes, &book); err != nil {
		return err
	} else {
		persistence.Store.CreateBook(&book)
		WriteJson(w, book)
		return nil
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) error {
	var book model.Book
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bytes, &book); err != nil {
		return err
	}

	if err := persistence.Store.UpdateBook(&book); err != nil {
		return err
	} else {
		WriteJson(w, book)
	}
	return nil
}
