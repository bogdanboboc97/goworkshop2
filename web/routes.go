package web

import "net/http"

const (
	authorBaseUrl   = "/authors"
	authorByUuidUrl = authorBaseUrl + "/{uuid}"
	authorsByName = authorBaseUrl + "ByName/{name}"
	booksBaseUrl    = "/books"
	booksByTitle = booksBaseUrl + "ByTitle/{title}"
	bookByUuidUrl   = booksBaseUrl + "/{uuid}"
	authorByFirstName = authorBaseUrl + "ByFirstName/{name}"
	authorByLastName = authorBaseUrl + "ByLastName/{name}"
)

type Route struct {
	//method type (GET, POST, PUT, DELETE, etc)
	Method string

	//the method path
	Pattern string

	//the method that the endpoint should call
	HandlerFunc RouteFunc
}

type RouteFunc func(http.ResponseWriter, *http.Request) error

type Routes []Route

var routes = Routes{
	//book_handlers
	Route{
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: Index,
	},
	Route{
		Method:      "GET",
		Pattern:     booksBaseUrl,
		HandlerFunc: GetAllBooks,
	},
	Route{
		Method:      "POST",
		Pattern:     booksBaseUrl,
		HandlerFunc: AddBook,
	},
	Route{
		Method:      "GET",
		Pattern:     bookByUuidUrl,
		HandlerFunc: GetBookByUUID,
	},
	Route{
		Method:      "GET",
		Pattern:	 booksByTitle,
		HandlerFunc: GetBooksByTitle,
	},
	Route{
		Method:      "DELETE",
		Pattern:     bookByUuidUrl,
		HandlerFunc: DeleteBookByUUID,
	},
	Route{
		Method:      "PUT",
		Pattern:     bookByUuidUrl,
		HandlerFunc: UpdateBook,
	},

	//author_handlers
	Route{
		Method:      "GET",
		Pattern:     authorBaseUrl,
		HandlerFunc: GetAllAuthors,
	},
	Route{
		Method:      "GET",
		Pattern:     authorsByName,
		HandlerFunc: GetAuthorsByName,
	},
	Route{
		Method:      "GET",
		Pattern:     authorByFirstName,
		HandlerFunc: GetAuthorsByFirstName,
	},
	Route{
		Method:      "GET",
		Pattern:     authorByLastName,
		HandlerFunc: GetAuthorsByLastName,
	},
	Route{
		Method:      "POST",
		Pattern:     authorBaseUrl,
		HandlerFunc: AddAuthor,
	},
	Route{
		Method:      "GET",
		Pattern:     authorByUuidUrl,
		HandlerFunc: GetAuthorByUUID,
	},
	Route{
		Method:      "DELETE",
		Pattern:     authorByUuidUrl,
		HandlerFunc: DeleteAuthorByUUID,
	},
	Route{
		Method:      "PUT",
		Pattern:     authorByUuidUrl,
		HandlerFunc: UpdateAuthor,
	},
}
